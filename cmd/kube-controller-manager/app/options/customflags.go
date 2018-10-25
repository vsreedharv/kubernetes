/*
Copyright 2018 The Kubernetes Authors.

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

package options

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"

	// ensure libs have a chance to globally register their flags
	_ "k8s.io/kubernetes/pkg/cloudprovider/providers"
)

// AddGenericFlags explicitly registers flags that internal packages register
// against the global flagsets from "flag". We do this in order to prevent
// unwanted flags from leaking into the kube-controller-manager's flagset.
func AddGenericFlags(fs *pflag.FlagSet) {
	addCloudProviderFlags(fs)
}

// addCloudProviderFlags adds flags from k8s.io/kubernetes/pkg/cloudprovider/providers.
func addCloudProviderFlags(fs *pflag.FlagSet) {
	// lookup flags in global flag set and re-register the values with our flagset
	global := flag.CommandLine
	local := pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)

	register(global, local, "cloud-provider-gce-lb-src-cidrs")

	fs.AddFlagSet(local)
}

// register adds a flag to local that targets the Value associated with the Flag named globalName in global
func register(global *flag.FlagSet, local *pflag.FlagSet, globalName string) {
	if f := global.Lookup(globalName); f != nil {
		local.AddGoFlag(f)
	} else {
		panic(fmt.Sprintf("failed to find flag in global flagset (flag): %s", globalName))
	}
}
