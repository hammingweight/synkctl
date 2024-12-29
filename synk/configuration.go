package synk

type Configuration struct {
	Endpoint string
	User     string
	Password string
}

func CreateConfigurationFile(fileName string, configuration Configuration) error {
	return nil
}
