package cmd

import (
	"errors"
)

var (
	ErrCantCreateConfigFile = errors.New("can't create config file")
	ErrCantAuthenticateUser = errors.New("can't authenticate user")
)
