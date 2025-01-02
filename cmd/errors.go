package cmd

import (
	"errors"
)

var (
	ErrCantAuthenticateUser       = errors.New("can't authenticate user")
	ErrCantCreateConfigFile       = errors.New("can't create config file")
	ErrCantReadBatteryState       = errors.New("can't read battery state")
	ErrCantReadGridState          = errors.New("can't read grid state")
	ErrCantReadInputState         = errors.New("can't read input state")
	ErrCantReadInverterSettings   = errors.New("can't read inverter settings")
	ErrCantReadLoadStatistics     = errors.New("can't read load statistics")
	ErrCantUpdateInverterSettings = errors.New("can't update inverter settings")
	ErrNoInverterSerialNumber     = errors.New("no inverter serial number (--inverter) supplied")
	ErrUnexpectedArgument         = errors.New("unexpected argument")
)
