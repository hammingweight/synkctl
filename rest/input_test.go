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

package rest

import (
	"encoding/json"
	"io"
	"os"
	"testing"
)

func TestReadPVOK(t *testing.T) {
	f, err := os.Open("testdata/input.json")
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	i := Input{&SynkObject{}}
	err = json.Unmarshal(data, i.SynkObject)
	if err != nil {
		t.Fatal(err)
	}
	pv, _ := i.PV(0)
	exp := "2025-01-22 17:55:06"
	if pv["time"] != exp {
		t.Errorf("Expected %s, got %v", exp, pv["time"])
	}
	pv, _ = i.PV(1)
	if pv["time"] != exp {
		t.Errorf("Expected %s, got %v", exp, pv["time"])
	}
}

func TestReadPVOutOfRange(t *testing.T) {
	f, err := os.Open("testdata/input.json")
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	i := Input{&SynkObject{}}
	err = json.Unmarshal(data, i.SynkObject)
	if err != nil {
		t.Fatal(err)
	}
	_, ok := i.PV(2)
	if ok {
		t.Error("expected value to be out of range")
	}
}
