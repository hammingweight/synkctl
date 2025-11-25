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

package configuration

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-yaml/yaml"
)

func TestReadConfiguration(t *testing.T) {
	configData := `endpoint: http://example.com
user: carl
password: secret`
	b := bytes.NewBufferString(configData)
	configuration, err := readConfiguration(b)
	if err != nil {
		t.Fatal("error: ", err)
	}
	api := "http://example.com"
	user := "carl"
	password := "secret"
	if configuration.Endpoint != api {
		t.Errorf("expected Endpoint: %s, got %s", api, configuration.Endpoint)
	}
	if configuration.User != user {
		t.Errorf("expected User: %s, got %s", user, configuration.User)
	}
	if configuration.Password != password {
		t.Errorf("expected Password: %s, got %s", password, configuration.Password)
	}
}

func TestReadConfigurationFromFile(t *testing.T) {
	configuration, err := ReadConfigurationFromFile("./testdata/testconfig")
	if err != nil {
		t.Fatal("failed with error: ", err)
	}
	api := "https://example.com/"
	user := "foo@example.com"
	password := "mypassword"
	if configuration.Endpoint != api {
		t.Errorf("expected Endpoint: %s, got %s", api, configuration.Endpoint)
	}
	if configuration.User != user {
		t.Errorf("expected User: %s, got %s", user, configuration.User)
	}
	if configuration.Password != password {
		t.Errorf("expected Password: %s, got %s", password, configuration.Password)
	}
	if configuration.DefaultInverterSN != "" {
		t.Errorf("expected DefaultInverterId: \"\", got %s", configuration.DefaultInverterSN)
	}
}

func TestReadConfigurationFromFileWithInverterSerialNumber(t *testing.T) {
	configuration, err := ReadConfigurationFromFile("./testdata/testconfig_with_inverter_id")
	if err != nil {
		t.Fatal("failed with error: ", err)
	}
	sn := "12345678"
	if configuration.DefaultInverterSN != sn {
		t.Errorf("expected DefaultInverterSN: %s got %s", sn, configuration.DefaultInverterSN)
	}
}

func TestReadConfigurationNoEndpoint(t *testing.T) {
	configuration, err := ReadConfigurationFromFile("./testdata/testconfig_no_endpoint")
	if err != nil {
		t.Fatal("failed with error: ", err)
	}
	if configuration.Endpoint != DefaultEndpoint {
		t.Errorf("expected endpoint: %s got %s", DefaultEndpoint, configuration.Endpoint)
	}
}

func TestReadUnknownKeyInConfigurationFile(t *testing.T) {
	_, err := ReadConfigurationFromFile("./testdata/testconfig_with_bad_key")
	if err == nil {
		t.Error("accepted key 'bad_key'")
	}
}

func TestWriteConfiguration(t *testing.T) {
	api := "https://example.com/"
	user := "hamming"
	password := "weight"
	configuration := &Configuration{
		Endpoint: api,
		User:     user,
		Password: password,
	}
	w := &bytes.Buffer{}
	if err := writeConfiguration(w, configuration); err != nil {
		t.Fatal("error:", err)
	}
	configMap := map[string]string{}
	if err := yaml.Unmarshal(w.Bytes(), configMap); err != nil {
		t.Fatal("error:", err)
	}
	if len(configMap) != 4 {
		t.Errorf("expected 4 elements, got %d\n", len(configMap))
	}
	if configuration.Endpoint != api {
		t.Errorf("expected Endpoint: %s, got %s", api, configuration.Endpoint)
	}
	if configuration.User != user {
		t.Errorf("expected User: %s, got %s", user, configuration.User)
	}
	if configuration.Password != password {
		t.Errorf("expected Password: %s, got %s", password, configuration.Password)
	}
}

func TestWriteConfigurationToFile(t *testing.T) {
	api := "https://api.sunsynk.net/"
	user := "carl"
	password := "Foobar"
	configuration := &Configuration{
		Endpoint: api,
		User:     user,
		Password: password,
	}
	filename := filepath.Join(t.TempDir(), "config")
	if err := WriteConfigurationToFile(filename, configuration); err != nil {
		t.Fatal("error: ", err)
	}
	f, err := os.Open(filename)
	if err != nil {
		t.Fatal("error: ", err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal("error: ", err)
	}
	configMap := map[string]string{}
	if err = yaml.Unmarshal(data, configMap); err != nil {
		t.Fatal("error: ", err)
	}
	if len(configMap) != 4 {
		t.Errorf("expected 4 elements, got %d\n", len(configMap))
	}
	if configuration.Endpoint != api {
		t.Errorf("expected Endpoint: %s, got %s", api, configuration.Endpoint)
	}
	if configuration.User != user {
		t.Errorf("expected User: %s, got %s", user, configuration.User)
	}
	if configuration.Password != password {
		t.Errorf("expected Password: %s, got %s", password, configuration.Password)
	}
}
