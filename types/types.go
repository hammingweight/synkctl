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

package types

import (
	"errors"
	"strconv"
	"strings"
)

// OnOff represents a setting that is on (true) or off (false) or
// undefined.
type OnOff string

func (flag *OnOff) String() string {
	return string(*flag)
}

// Type returns a description of the OnOff type.
func (flag *OnOff) Type() string {
	return "on/off"
}

// Set converts a string to the value "on" or "off" or returns
// an error if the value cannot be converted. For example, both
// "on" and "true" will set the flag to the value "on".
func (flag *OnOff) Set(v string) error {
	switch v {
	case "on":
		*flag = "on"
	case "off":
		*flag = "off"
	default:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return errors.New("must be on/true or off/false")

		}
		if b {
			*flag = "on"
		} else {
			*flag = "off"
		}
	}
	return nil
}

// Bool returns true if the flag valye is "on", otherwise it
// returns false.
func (flag *OnOff) Bool() bool {
	return *flag == "on"
}

// NewOnOff converts a bool to the values "on" (true) or
// "off" (false).
func NewOnOff(b bool) OnOff {
	if b {
		return "on"
	}
	return "off"
}

// Percentage represents a percentage.
type Percentage string

func (p *Percentage) String() string {
	return string(*p)
}

// Type returns a desccription of the percentage type.
func (p *Percentage) Type() string {
	return "percentage"
}

// Set will update the value of a percentage type if the value
// can be parsed as a float, otherwise an error is returned.
func (p *Percentage) Set(v string) error {
	_, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return err
	}
	*p = Percentage(v)
	return nil
}

// Int returns the integer part of a percentage.
func (p *Percentage) Int() int {
	v, _ := strconv.ParseFloat(p.String(), 64)
	return int(v)
}

// CSV is a string with comma-separated values.
type CSV string

func (csv *CSV) String() string {
	return string(*csv)
}

// Type returns a description of the CSV type.
func (csv *CSV) Type() string {
	return "csv"
}

// Set sets the value.
func (csv *CSV) Set(v string) error {
	*csv = CSV(v)
	return nil
}

// Values returns the values as a slice
func (csv *CSV) Values() []string {
	values := strings.Split(string(*csv), ",")
	r := []string{}
	for _, v := range values {
		if v != "" {
			v = strings.TrimSpace(v)
			r = append(r, v)
		}
	}
	return r
}
