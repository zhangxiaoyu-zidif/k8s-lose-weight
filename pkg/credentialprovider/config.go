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
	"fmt"
	"net/http"
)

// DockerConfigJson represents ~/.docker/config.json file info
// see https://github.com/docker/docker/pull/12009
type DockerConfigJson struct {
	Auths DockerConfig `json:"auths"`
	// +optional
	HttpHeaders map[string]string `json:"HttpHeaders,omitempty"`
}

// DockerConfig represents the config file used by the docker CLI.
// This config that represents the credentials that should be used
// when pulling images from specific image repositories.
type DockerConfig map[string]DockerConfigEntry

type DockerConfigEntry struct {
	Username string
	Password string
	Email    string
	Provider DockerConfigProvider
}


func SetPreferredDockercfgPath(path string) {

}

func GetPreferredDockercfgPath() string {
	return ""
}

//DefaultDockercfgPaths returns default search paths of .dockercfg
func DefaultDockercfgPaths() []string {
	return []string{GetPreferredDockercfgPath(), "", "", ""}
}

//DefaultDockerConfigJSONPaths returns default search paths of .docker/config.json
func DefaultDockerConfigJSONPaths() []string {
	return []string{GetPreferredDockercfgPath(), "", "", ""}
}

// ReadDockercfgFile attempts to read a legacy dockercfg file from the given paths.
// if searchPaths is empty, the default paths are used.
func ReadDockercfgFile(searchPaths []string) (cfg DockerConfig, err error) {

	return nil, fmt.Errorf("couldn't find valid .dockercfg after checking in %v", searchPaths)
}

// ReadDockerConfigJSONFile attempts to read a docker config.json file from the given paths.
// if searchPaths is empty, the default paths are used.
func ReadDockerConfigJSONFile(searchPaths []string) (cfg DockerConfig, err error) {

	return nil, fmt.Errorf("couldn't find valid .docker/config.json after checking in %v", searchPaths)
}

func ReadDockerConfigFile() (cfg DockerConfig, err error) {
	if cfg, err := ReadDockerConfigJSONFile(nil); err == nil {
		return cfg, nil
	}
	// Can't find latest config file so check for the old one
	return ReadDockercfgFile(nil)
}

// HttpError wraps a non-StatusOK error code as an error.
type HttpError struct {
	StatusCode int
	Url        string
}

// Error implements error
func (he *HttpError) Error() string {
	return fmt.Sprintf("http status code: %d while fetching url %s",
		he.StatusCode, he.Url)
}

func ReadUrl(url string, client *http.Client, header *http.Header) (body []byte, err error) {
	return nil, fmt.Errorf("couldn't find valid .docker/config.json after checking")

}

func ReadDockerConfigFileFromUrl(url string, client *http.Client, header *http.Header) (cfg DockerConfig, err error) {
	return nil, fmt.Errorf("couldn't find valid .docker/config.json after checking")
}


func (ident *DockerConfigEntry) UnmarshalJSON(data []byte) error {

	return nil
}

func (ident DockerConfigEntry) MarshalJSON() ([]byte, error) {
	return nil,nil
}
