/*
Copyright 2025 Carl Meijer.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
	ErrUnexpectedArguments        = errors.New("unexpected argument(s)")
)
