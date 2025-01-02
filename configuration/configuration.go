package configuration

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

const DefaultEndpoint = "https://api.sunsynk.net"

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

func ReadConfigurationFromFile(fileName string) (*Configuration, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	configuration, err := readConfiguration(f)
	return configuration, err
}

func New(user string, password string) (*Configuration, error) {
	if user == "" {
		return nil, errors.New("'user' cannot be empty")
	}
	if password == "" {
		return nil, errors.New("'password' cannot be empty")
	}
	return &Configuration{User: user, Password: password, Endpoint: DefaultEndpoint}, nil
}

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
