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

import "testing"

func TestOff(t *testing.T) {
	flag := onOff("")
	err := flag.Set("false")
	if err != nil {
		t.Error("'false' is a valid onOff value")
	}
	if flag.String() != "off" {
		t.Error("flag should have had value 'off'")
	}
	if flag.Bool() {
		t.Error("flag should have had boolean value 'false'")
	}
	err = flag.Set("off")
	if err != nil {
		t.Error("'off' is a valid onOff value")
	}
	if flag.String() != "off" {
		t.Error("flag should have had value 'off'")
	}
}

func TestOn(t *testing.T) {
	var flag onOff
	err := flag.Set("true")
	if err != nil {
		t.Error("'true' is a valid onOff value")
	}
	if !flag.Bool() {
		t.Error("flag should have had boolean value 'true'")
	}
	if flag.String() != "on" {
		t.Error("flag should have had value 'on'")
	}
	err = flag.Set("on")
	if err != nil {
		t.Error("'on' is a valid onOff value")
	}
	s := string(flag)
	if s != "on" {
		t.Error("flag should have had value 'on'")
	}
}

func TestInvalidValues(t *testing.T) {
	var flag onOff
	err := flag.Set("")
	if err == nil {
		t.Error("The empty string is not a valid value")
	}
}

func TestSetPercentage(t *testing.T) {
	var p percentage
	err := p.Set("hello")
	if err == nil {
		t.Error("'hello' is not a valid percentage")
	}
	err = p.Set("36")
	if err != nil {
		t.Error("'36' is a valid percentage")
	}
	if p.Int() != 36 {
		t.Errorf("Expected value 36, got %d\n", p.Int())
	}
	err = p.Set("99.9")
	if err != nil {
		t.Error("percentages can be floating points")
	}
	if p.Int() != 99 {
		t.Errorf("Expected value 99, got %d\n", p.Int())
	}
}
