/*
Copyright 2014 The Kubernetes Authors.

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

package credentialprovider

import (


)

// All registered credential providers.
var providers = make(map[string]DockerConfigProvider)

// RegisterCredentialProvider is called by provider implementations on
// initialization to register themselves, like so:
//   func init() {
//    	RegisterCredentialProvider("name", &myProvider{...})
//   }
func RegisterCredentialProvider(name string, provider DockerConfigProvider) {

}

// NewDockerKeyring creates a DockerKeyring to use for resolving credentials,
// which lazily draws from the set of registered credential providers.
func NewDockerKeyring() DockerKeyring {
	keyring := &lazyDockerKeyring{
		Providers: make([]DockerConfigProvider, 0),
	}


	return keyring
}
