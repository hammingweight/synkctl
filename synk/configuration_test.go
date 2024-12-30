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

func TestReadFromFile(t *testing.T) {
	configuration := &Configuration{}
	err := configuration.ReadFromFile("./testdata/testconfig")
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

func TestReadFromFileWithInverterSerialNumber(t *testing.T) {
	configuration := &Configuration{}
	err := configuration.ReadFromFile("./testdata/testconfig_with_inverter_id")
	if err != nil {
		t.Fatal("failed with error: ", err)
	}
	sn := "12345678"
	if configuration.DefaultInverterSN != sn {
		t.Errorf("expected DefaultInverterSN: %s got %s", sn, configuration.DefaultInverterSN)
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
