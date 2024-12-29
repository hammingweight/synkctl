package synk

import (
	"io"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Configuration struct {
	Endpoint string `yaml:"endpoint"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func createAndOpenFile(filename string) (*os.File, error) {
	dir := filepath.Dir(filename)
	err := os.MkdirAll(dir, 0700)
	if err != nil {
		return nil, err
	}
	return os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0600)
}

func writeConfiguration(writer io.Writer, configuration Configuration) error {
	marshalledData, err := yaml.Marshal(configuration)
	if err != nil {
		return nil
	}
	_, err = writer.Write(marshalledData)
	return err
}

func CreateConfigurationFile(fileName string, configuration Configuration) error {
	writer, err := createAndOpenFile(fileName)
	if err != nil {
		return err
	}
	defer writer.Close()
	return writeConfiguration(writer, configuration)
}
