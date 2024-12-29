package synk

import (
	"bytes"
	"testing"
)

func TestRead(t *testing.T) {
	configData := `endpoint: http://example.com
user: carl
password: secret`
	b := bytes.NewBufferString(configData)
	configuration := &Configuration{}
	err := configuration.read(b)
	if err != nil {
		t.Fatal("failed with error: ", err)
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
