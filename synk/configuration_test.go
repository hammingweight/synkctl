package synk

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/go-yaml/yaml"
)

func TestRead(t *testing.T) {
	configData := `endpoint: http://example.com
user: carl
password: secret`
	b := bytes.NewBufferString(configData)
	configuration := &Configuration{}
	err := configuration.read(b)
	if err != nil {
		t.Fatal("error: ", err)
	}
	if configuration.Endpoint != "http://example.com" {
		t.Errorf("expected Endpoint: http://example.com, got %s", configuration.Endpoint)
	}
	if configuration.User != "carl" {
		t.Errorf("expected User: carl, got %s", configuration.User)
	}
	if configuration.Password != "secret" {
		t.Errorf("expected Password: secret, got %s", configuration.Password)
	}
}

func TestReadFromFile(t *testing.T) {
	configuration := &Configuration{}
	err := configuration.ReadFromFile("./testdata/testconfig")
	if err != nil {
		t.Fatal("failed with error: ", err)
	}
	if configuration.Endpoint != "http://example.com" {
		t.Errorf("expected Endpoint: http://example.com, got %s", configuration.Endpoint)
	}
	if configuration.User != "foo" {
		t.Errorf("expected User: foo, got %s", configuration.User)
	}
	if configuration.Password != "bar" {
		t.Errorf("expected Password: bar, got %s", configuration.Password)
	}
}

func TestWrite(t *testing.T) {
	api := "https://example.com/"
	user := "hamming"
	password := "weight"
	configuration := &Configuration{
		Endpoint: api,
		User:     user,
		Password: password,
	}
	w := &bytes.Buffer{}
	err := configuration.write(w)
	if err != nil {
		t.Fatal("error:", err)
	}
	configMap := map[string]string{}
	err = yaml.Unmarshal(w.Bytes(), configMap)
	if err != nil {
		t.Fatal("error:", err)
	}
	if len(configMap) != 3 {
		t.Errorf("expected 3 elements, got %d\n", len(configMap))
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

func TestWriteToFile(t *testing.T) {
	api := "https://api.sunsynk.net/"
	user := "carl"
	password := "Foobar"
	configuration := &Configuration{
		Endpoint: api,
		User:     user,
		Password: password,
	}
	filename := filepath.Join(t.TempDir(), "config")
	err := configuration.WriteToFile(filename)
	if err != nil {
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
	err = yaml.Unmarshal(data, configMap)
	if err != nil {
		t.Fatal("error: ", err)
	}
	if len(configMap) != 3 {
		t.Errorf("expected 3 elements, got %d\n", len(configMap))
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
