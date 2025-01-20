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

import "strconv"

type onOff string

func (flag *onOff) String() string {
	return string(*flag)
}

func (flag *onOff) Type() string {
	return "on/off"
}

func (flag *onOff) Set(v string) error {
	switch v {
	case "on":
		*flag = "on"
	case "off":
		*flag = "off"
	default:
		b, err := strconv.ParseBool(v)
		if err != nil {
			return ErrInvalidFlag
		}
		if b {
			*flag = "on"
		} else {
			*flag = "off"
		}
	}
	return nil
}

type percentage string

func (p *percentage) String() string {
	return string(*p)
}

func (p *percentage) Type() string {
	return "percentage"
}

func (p *percentage) Set(v string) error {
	_, err := strconv.Atoi(v)
	if err != nil {
		return err
	}
	*p = percentage(v)
	return nil
}

func (p *percentage) Int() int {
	v, _ := strconv.Atoi(p.String())
	return v
}
