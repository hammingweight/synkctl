/*
Copyright 2025 Carl Meijer.

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

// Package configuration configures credentials for accessing the
// SunSynk API.
package configuration

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

// DefaultEndpoint is SunSynk's region 1 endpoint.
const DefaultEndpoint = "https://api.sunsynk.net"

// Configuration stores the credentials to access the SunSynk API
type Configuration struct {
	Endpoint          string `yaml:"endpoint"`
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	DefaultInverterSN string `yaml:"default_inverter_sn,omitempty"`
}

func createAndOpenFile(filename string) (*os.File, error) {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
}

func writeConfiguration(writer io.Writer, configuration *Configuration) error {
	marshalledData, err := yaml.Marshal(configuration)
	if err != nil {
		return nil
	}
	_, err = writer.Write(marshalledData)
	return err
}

// WriteConfigurationToFile writes a marshalled configuration object to a YAML file
func WriteConfigurationToFile(fileName string, configuration *Configuration) error {
	writer, err := createAndOpenFile(fileName)
	if err != nil {
		return err
	}
	defer writer.Close()
	return writeConfiguration(writer, configuration)
}

func readConfiguration(reader io.Reader) (*Configuration, error) {
	configuration := &Configuration{}
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	err = yaml.UnmarshalStrict(data, configuration)
	return configuration, err
}

// ReadConfigurationFromFile unnmarshals a configuration object from a YAML file
func ReadConfigurationFromFile(fileName string) (*Configuration, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	configuration, err := readConfiguration(f)
	return configuration, err
}

// New constructs a configuration object with a specified username/password
// combination using the default region 1 (London) endpoint
func New(user string, password string) (*Configuration, error) {
	if user == "" {
		return nil, errors.New("'user' cannot be empty")
	}
	if password == "" {
		return nil, errors.New("'password' cannot be empty")
	}
	return &Configuration{User: user, Password: password, Endpoint: DefaultEndpoint}, nil
}

// NewWithEndpoint constructs a configuration object with username/password credentials and an
// overridden API endpoint
func NewWithEndpoint(user string, password string, endpoint string) (*Configuration, error) {
	config, err := New(user, password)
	if err != nil {
		return nil, err
	}
	if endpoint == "" {
		return nil, errors.New("'endpoint' cannot be empty")
	}
	config.Endpoint = endpoint
	return config, err
}

// DefaultConfigurationFile gets the path to the default location of the synk config file ($HOME/.synk/config on Linux).
func DefaultConfigurationFile() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".synk", "config"), nil
}
