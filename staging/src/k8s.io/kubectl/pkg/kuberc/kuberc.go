/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kuberc

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"k8s.io/kubectl/pkg/config"
	kuberc "k8s.io/kubectl/pkg/config/install"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
)

const RecommendedKubeRCFileName = "kuberc"

var (
	RecommendedConfigDir  = filepath.Join(homedir.HomeDir(), clientcmd.RecommendedHomeDir)
	RecommendedKubeRCFile = filepath.Join(RecommendedConfigDir, RecommendedKubeRCFileName)

	aliasNameRegex = regexp.MustCompile("^[a-zA-Z]+$")
	shortHandRegex = regexp.MustCompile("^-[a-zA-Z]+$")

	scheme        = runtime.NewScheme()
	strictCodecs  = serializer.NewCodecFactory(scheme, serializer.EnableStrict)
	lenientCodecs = serializer.NewCodecFactory(scheme)
)

func init() {
	kuberc.Install(scheme)
}

// PreferencesHandler is responsible for setting default flags
// arguments based on user's kuberc configuration.
type PreferencesHandler interface {
	AddFlags(flags *pflag.FlagSet)
	Apply(rootCmd *cobra.Command, args []string, errOut io.Writer) ([]string, error)
}

// Preferences stores the kuberc file coming either from environment variable
// or file from set in flag or the default kuberc path.
type Preferences struct {
	getPreferencesFunc func(kuberc string) (*config.Preference, error)

	aliases map[string]struct{}
}

// NewPreferences returns initialized Prefrences object.
func NewPreferences() PreferencesHandler {
	return &Preferences{
		getPreferencesFunc: DefaultGetPreferences,
		aliases:            make(map[string]struct{}),
	}
}

type aliasing struct {
	appendArgs  []string
	prependArgs []string
	flags       []config.CommandOverrideFlag
	command     *cobra.Command
}

// AddFlags adds kuberc related flags into the command.
func (p *Preferences) AddFlags(flags *pflag.FlagSet) {
	flags.String("kuberc", "", "Path to the kuberc file to use for preferences. This can be disabled by exporting KUBECTL_KUBERC=false.")
}

// Apply firstly applies the aliases in the preferences file and secondly overrides
// the default values of flags.
func (p *Preferences) Apply(rootCmd *cobra.Command, args []string, errOut io.Writer) ([]string, error) {
	if len(args) <= 1 {
		return args, nil
	}

	kubercPath, err := getExplicitKuberc(args)
	if err != nil {
		return args, err
	}
	kuberc, err := p.getPreferencesFunc(kubercPath)
	if err != nil {
		return args, fmt.Errorf("kuberc error %w", err)
	}

	if kuberc == nil {
		return args, nil
	}

	err = validate(kuberc)
	if err != nil {
		return args, err
	}

	args, err = p.applyAliases(rootCmd, kuberc, args, errOut)
	if err != nil {
		return args, err
	}
	err = p.applyOverrides(rootCmd, kuberc, args, errOut)
	if err != nil {
		return args, err
	}
	return args, nil
}

// applyOverrides finds the command and sets the defaulted flag values in kuberc.
func (p *Preferences) applyOverrides(rootCmd *cobra.Command, kuberc *config.Preference, args []string, errOut io.Writer) error {
	args = args[1:]
	cmd, _, err := rootCmd.Find(args)
	if err != nil {
		return nil
	}

	for _, c := range kuberc.Overrides {
		parsedCmds := strings.Fields(c.Command)
		overrideCmd, _, err := rootCmd.Find(parsedCmds)
		if err != nil {
			fmt.Fprintf(errOut, "Warning: command %q not found to set kuberc override\n", c.Command)
			continue
		}
		if overrideCmd.Name() != cmd.Name() {
			continue
		}

		if _, ok := p.aliases[cmd.Name()]; ok {
			return fmt.Errorf("alias %s can not be overridden", cmd.Name())
		}

		// This function triggers merging the persistent flags in the parent commands.
		_ = cmd.InheritedFlags()

		allShorthands := make(map[string]struct{})
		cmd.Flags().VisitAll(func(flag *pflag.Flag) {
			if flag.Shorthand != "" {
				allShorthands[flag.Shorthand] = struct{}{}
			}
		})

		for _, fl := range c.Flags {
			existingFlag := cmd.Flag(fl.Name)
			if existingFlag == nil {
				return fmt.Errorf("invalid flag %s for command %s", fl.Name, c.Command)
			}
			if searchInArgs(existingFlag.Name, existingFlag.Shorthand, allShorthands, args) {
				// Don't modify the value implicitly, if it is passed in args explicitly
				continue
			}
			err = cmd.Flags().Set(fl.Name, fl.Default)
			if err != nil {
				return fmt.Errorf("could not apply override value %s to flag %s in command %s err: %w", fl.Default, fl.Name, c.Command, err)
			}
		}
	}

	return nil
}

// applyAliases firstly appends all defined aliases in kuberc file to the root command.
// Since there may be several alias definitions belonging to the same command, it extracts the
// alias that is currently executed from args. After that it sets the flag definitions in alias as default values
// of the command. Lastly, others parameters (e.g. resources, etc.) that are passed as arguments in kuberc
// is appended into the command args.
func (p *Preferences) applyAliases(rootCmd *cobra.Command, kuberc *config.Preference, args []string, errOut io.Writer) ([]string, error) {
	_, _, err := rootCmd.Find(args[1:])
	if err == nil {
		// Command is found, no need to continue for aliasing
		return args, nil
	}

	var aliasArgs *aliasing

	var commandName string // first "non-flag" arguments
	var commandIndex int
	for index, arg := range args[1:] {
		if !strings.HasPrefix(arg, "-") {
			commandName = arg
			commandIndex = index + 1
			break
		}
	}

	for _, alias := range kuberc.Aliases {
		p.aliases[alias.Name] = struct{}{}
		if alias.Name != commandName {
			continue
		}

		// do not allow shadowing built-ins
		if _, _, err := rootCmd.Find([]string{alias.Name}); err == nil {
			fmt.Fprintf(errOut, "Warning: Setting alias %q to a built-in command is not supported\n", alias.Name)
			break
		}

		commands := strings.Fields(alias.Command)
		existingCmd, flags, err := rootCmd.Find(commands)
		if err != nil {
			return args, fmt.Errorf("command %q not found to set alias %q: %v", alias.Command, alias.Name, flags)
		}

		newCmd := *existingCmd
		newCmd.Use = alias.Name
		newCmd.Aliases = []string{}
		aliasCmd := &newCmd

		aliasArgs = &aliasing{
			prependArgs: alias.PrependArgs,
			appendArgs:  alias.AppendArgs,
			flags:       alias.Flags,
			command:     aliasCmd,
		}
		break
	}

	if aliasArgs == nil {
		// pursue with the current behavior.
		// This might be a built-in command, external plugin, etc.
		return args, nil
	}

	rootCmd.AddCommand(aliasArgs.command)

	foundAliasCmd, _, err := rootCmd.Find([]string{commandName})
	if err != nil {
		return args, nil
	}

	// This function triggers merging the persistent flags in the parent commands.
	_ = foundAliasCmd.InheritedFlags()

	allShorthands := make(map[string]struct{})
	foundAliasCmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Shorthand != "" {
			allShorthands[flag.Shorthand] = struct{}{}
		}
	})

	for _, fl := range aliasArgs.flags {
		existingFlag := foundAliasCmd.Flag(fl.Name)
		if existingFlag == nil {
			return args, fmt.Errorf("invalid alias flag %s in alias %s", fl.Name, args[0])
		}
		if searchInArgs(existingFlag.Name, existingFlag.Shorthand, allShorthands, args) {
			// Don't modify the value implicitly, if it is passed in args explicitly
			continue
		}
		err = foundAliasCmd.Flags().Set(fl.Name, fl.Default)
		if err != nil {
			return args, fmt.Errorf("could not apply value %s to flag %s in alias %s err: %w", fl.Default, fl.Name, args[0], err)
		}
	}

	if len(aliasArgs.prependArgs) > 0 {
		// prependArgs defined in kuberc should be inserted after the alias name.
		if commandIndex+1 >= len(args) {
			// command is the last item, we simply append just like appendArgs
			args = append(args, aliasArgs.prependArgs...)
		} else {
			args = append(args[:commandIndex+1], append(aliasArgs.prependArgs, args[commandIndex+1:]...)...)
		}
	}
	if len(aliasArgs.appendArgs) > 0 {
		// appendArgs defined in kuberc should be appended to actual args.
		args = append(args, aliasArgs.appendArgs...)
	}
	// Cobra (command.go#L1078) appends only root command's args into the actual args and ignores the others.
	// We are appending the additional args defined in kuberc in here and
	// expect that it will be passed along to the actual command.
	rootCmd.SetArgs(args[1:])
	return args, nil
}

// DefaultGetPreferences returns KubeRCConfiguration.
// If users sets kuberc file explicitly in --kuberc flag, it has the highest
// priority. If not specified, it looks for in KUBERC environment variable.
// If KUBERC is also not set, it falls back to default .kuberc file at the same location
// where kubeconfig's defaults are residing in.
func DefaultGetPreferences(kuberc string) (*config.Preference, error) {
	kubeRCFile := RecommendedKubeRCFile
	explicitly := false
	if kuberc != "" {
		kubeRCFile = kuberc
		explicitly = true
	}

	if kubeRCFile == "" && os.Getenv("KUBERC") != "" {
		kubeRCFile = os.Getenv("KUBERC")
		explicitly = true
	}

	kubeRCBytes, err := os.ReadFile(kubeRCFile)
	if err != nil {
		if !explicitly {
			// We don't log if the kuberc file does not exist. Because user simply does not
			// specify neither default kuberc file nor explicitly pass it.
			// We'll continue to default behavior without raising any error.
			if !os.IsNotExist(err) {
				klog.V(4).Infof("error reading kuberc file %q: %v", kubeRCFile, err)
			}
			return nil, nil
		}
		// Kuberc is specified by user via flag or env variable and it gets an error.
		return nil, err
	}

	var pref config.Preference
	_, gvk, err := strictCodecs.UniversalDecoder().Decode(kubeRCBytes, nil, &pref)
	if err != nil {
		if explicitly {
			// explicitly specified kuberc can't be decoded and we short cut the process
			return nil, fmt.Errorf("could not be decoded gvk %s, err: %w", gvk, err)
		}
		// default kuberc is incompatible with this version, or it simply is invalid.
		// falling back to lenient decoding to do our best.
		_, gvk2, err2 := lenientCodecs.UniversalDecoder().Decode(kubeRCBytes, nil, &pref)
		if err2 != nil {
			return nil, fmt.Errorf("could not be decoded gvk %s, err: %w", gvk2, err2)
		}
		klog.Warningf("gvk %s could not be decoded with strict decoding, continue with less strict decoding: %s", gvk, err)
		gvk = gvk2
	}

	expectedGK := schema.GroupKind{
		Group: config.SchemeGroupVersion.Group,
		Kind:  "Preference",
	}
	if gvk.GroupKind() != expectedGK {
		return nil, fmt.Errorf("unsupported preference GVK %s", gvk.GroupKind().String())
	}

	return &pref, nil
}

// Normally, we should extract this value directly from kuberc flag.
// However, flag values are set during the command execution and
// we are in very early stages to prepare commands prior to execute them.
// Besides, we only need kuberc flag value in this stage.
func getExplicitKuberc(args []string) (string, error) {
	var kubercPath string
	for i, arg := range args {
		if arg == "--" {
			// flags after "--" does not represent any flag of
			// the command. We should short cut the iteration in here.
			break
		}
		if arg == "--kuberc" {
			if i+1 < len(args) {
				kubercPath = args[i+1]
				break
			}
			return "", fmt.Errorf("kuberc file is not found")
		} else if strings.Contains(arg, "--kuberc=") {
			parg := strings.Split(arg, "=")
			if len(parg) > 1 && parg[1] != "" {
				kubercPath = parg[1]
				break
			}
			return "", fmt.Errorf("kuberc file is not found")
		}
	}

	if kubercPath == "" {
		return "", nil
	}

	return kubercPath, nil
}

// searchInArgs searches the given key in the args and returns
// true, if it finds. Otherwise, it returns false.
func searchInArgs(flagName string, shorthand string, allShorthands map[string]struct{}, args []string) bool {
	for _, arg := range args {
		// if flag is set in args in "--flag value" or "--flag=value" format,
		// we should return it as found
		if fmt.Sprintf("--%s", flagName) == arg || strings.HasPrefix(arg, fmt.Sprintf("--%s=", flagName)) {
			return true
		}
		if shorthand == "" {
			continue
		}
		// shorthand can be in "-n value" or "-nvalue" format
		// it is guaranteed that shorthand is one letter. So that
		// checking just the prefix -oyaml also finds --output.
		if strings.HasPrefix(arg, fmt.Sprintf("-%s", shorthand)) {
			return true
		}

		if !shortHandRegex.MatchString(arg) {
			continue
		}

		// remove prefix "-"
		arg = arg[1:]
		// short hands can be in a combined "-abc" format.
		// First we need to ensure that all the values are shorthand to safely search ours.
		// Because we know that "-abcvalue" is not valid. So that we need to be sure that if we find
		// "b" it correctly refers to the shorthand "b" not arbitrary value "-cargb".
		arbitraryFound := false
		for _, runeValue := range shorthand {
			if _, ok := allShorthands[string(runeValue)]; !ok {
				arbitraryFound = true
				break
			}
		}
		if arbitraryFound {
			continue
		}
		// verified that all values are short hand. Now search ours
		if strings.Contains(arg, shorthand) {
			return true
		}
	}
	return false
}

func validate(plugin *config.Preference) error {
	validateFlag := func(flags []config.CommandOverrideFlag) error {
		for _, flag := range flags {
			if strings.HasPrefix(flag.Name, "-") {
				return fmt.Errorf("flag name %s should be in long form without dashes", flag.Name)
			}
		}
		return nil
	}
	aliases := make(map[string]struct{})
	for _, alias := range plugin.Aliases {
		if !aliasNameRegex.MatchString(alias.Name) {
			return fmt.Errorf("invalid alias name, can only include alphabetical characters")
		}

		if err := validateFlag(alias.Flags); err != nil {
			return err
		}

		if _, ok := aliases[alias.Name]; ok {
			return fmt.Errorf("duplicate alias name %s", alias.Name)
		}
		aliases[alias.Name] = struct{}{}
	}

	for _, override := range plugin.Overrides {
		if err := validateFlag(override.Flags); err != nil {
			return err
		}
	}

	return nil
}
