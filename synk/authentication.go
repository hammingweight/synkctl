package synk

type Tokens struct {
	Bearer  string
	Refresh string
}

type AuthenticationRequest struct {
	GrantType string `json:"grant_type"`
	User      string `json:"username"`
	Password  string `json:"password"`
}

func Authenticate(configFile string) (*Tokens, error) {
	config := &Configuration{}
	err := config.ReadFromFile(configFile)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
