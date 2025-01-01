package cmd

import (
	"errors"
)

var (
	ErrCantAuthenticateUser       = errors.New("can't authenticate user")
	ErrCantCreateConfigFile       = errors.New("can't create config file")
	ErrCantReadInverterSettings   = errors.New("can't read inverter settings")
	ErrCantUpdateInverterSettings = errors.New("can't update inverter settings")
	ErrNoInverterSerialNumber     = errors.New("no inverter serial number (--inverter) supplied")
	ErrUnexpectedArgument         = errors.New("unexpected argument")
)
