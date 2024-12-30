package synk

import (
	"io"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

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

func (configuration *Configuration) write(writer io.Writer) error {
	marshalledData, err := yaml.Marshal(configuration)
	if err != nil {
		return nil
	}
	_, err = writer.Write(marshalledData)
	return err
}

func (configuration *Configuration) WriteToFile(fileName string) error {
	writer, err := createAndOpenFile(fileName)
	if err != nil {
		return err
	}
	defer writer.Close()
	return configuration.write(writer)
}

func (configuration *Configuration) read(reader io.Reader) error {
	data, err := io.ReadAll(reader)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, configuration)
}

func (configuration *Configuration) ReadFromFile(fileName string) error {
	f, err := os.Open(fileName)
	if err != nil {
		return err
	}
	return configuration.read(f)
}
