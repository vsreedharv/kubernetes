/*
Copyright 2017 The Kubernetes Authors.

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

/*
Package bootstrap provides a token authenticator for TLS bootstrap secrets.
*/
package bootstrap

import (
	"strings"
	"time"

	"github.com/golang/glog"

	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apiserver/pkg/authentication/user"
	"k8s.io/kubernetes/pkg/api"
	bootstrapapi "k8s.io/kubernetes/pkg/bootstrap/api"
	"k8s.io/kubernetes/pkg/client/listers/core/internalversion"
)

const tokenSeparator = "."

// NewTokenAuthenticator initializes a bootstrap token authenticator.
//
// Lister is expected to be for the "kube-system" namespace.
func NewTokenAuthenticator(lister internalversion.SecretNamespaceLister) *TokenAuthenticator {
	return &TokenAuthenticator{lister}
}

// TokenAuthenticator authenticates bootstrap tokens from secrets in the API server.
type TokenAuthenticator struct {
	lister internalversion.SecretNamespaceLister
}

// AuthenticateToken tries to match the provided token to a bootstrap token secret
// in a given namespace. If found, it authenticates the token in the
// "system:bootstrappers" group and with the "system:bootstrap:(token-id)" username.
//
// All secrets must be of type "bootstrap.kubernetes.io/token". An example secret:
//
//     apiVersion: v1
//     kind: Secret
//     metadata:
//       # Name is arbitrary and isn't used for authentication purposes.
//       name: bootstrap-token-( token id )
//       namespace: kube-system
//     # Only secrets of this type will be evaluated.
//     type: bootstrap.kubernetes.io/token
//     data:
//       token-secret: ( private part of token )
//       token-id: ( token id )
//       # Required key usage.
//       usage-bootstrap-authentication: true
//       # May also contain an expiry.
//
// Tokens are expected to be of the form:
//
//     ( token-id ).( token-secret )
//
func (t *TokenAuthenticator) AuthenticateToken(token string) (user.Info, bool, error) {
	// If token isn't of form "x.x" then ignore it.
	if !strings.Contains(token, tokenSeparator) {
		return nil, false, nil
	}

	secrets, err := t.lister.List(labels.Everything())
	if err != nil {
		return nil, false, err
	}

	for _, secret := range secrets {
		if string(secret.Type) != string(bootstrapapi.SecretTypeBootstrapToken) || secret.Data == nil {
			continue
		}

		ts := getSecretString(secret, bootstrapapi.BootstrapTokenSecretKey)
		if len(ts) == 0 {
			continue
		}
		id := getSecretString(secret, bootstrapapi.BootstrapTokenIDKey)
		if len(id) == 0 {
			continue
		}

		if token != id+tokenSeparator+ts {
			continue
		}

		if isSecretExpired(secret) {
			continue
		}

		if getSecretString(secret, bootstrapapi.BootstrapTokenUsageAuthentication) != "true" {
			glog.V(3).Infof("Bearer token matching bootstrap Secret %s/%s not marked %s=true.",
				secret.Namespace, secret.Name, bootstrapapi.BootstrapTokenUsageAuthentication)
			continue
		}

		return &user.DefaultInfo{
			Name:   bootstrapapi.BootstrapUserPrefix + string(id),
			Groups: []string{bootstrapapi.BootstrapGroup},
		}, true, nil
	}
	return nil, false, nil
}

// Everything bellow this also copied from k8s.io/kubernetes/pkg/bootstrap/api.

func getSecretString(secret *api.Secret, key string) string {
	if secret.Data == nil {
		return ""
	}
	if val, ok := secret.Data[key]; ok {
		return string(val)
	}
	return ""
}

// isSecretExpired returns true if the Secret is expired.
func isSecretExpired(secret *api.Secret) bool {
	expiration := getSecretString(secret, bootstrapapi.BootstrapTokenExpirationKey)
	if len(expiration) > 0 {
		expTime, err2 := time.Parse(time.RFC3339, expiration)
		if err2 != nil {
			glog.V(3).Infof("Unparseable expiration time (%s) in %s/%s Secret: %v. Treating as expired.",
				expiration, secret.Namespace, secret.Name, err2)
			return true
		}
		if time.Now().After(expTime) {
			glog.V(3).Infof("Expired bootstrap token in %s/%s Secret: %v",
				secret.Namespace, secret.Name, expiration)
			return true
		}
	}
	return false
}
