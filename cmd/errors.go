package cmd

import (
	"errors"
)

var (
	ErrCantAuthenticateUser   = errors.New("can't authenticate user")
	ErrCantCreateConfigFile   = errors.New("can't create config file")
	ErrNoInverterSerialNumber = errors.New("no inverter serial number (--inverter) supplied")
	ErrUnexpectedArgument     = errors.New("unexpected argument")
)
