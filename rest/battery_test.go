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

func Test_CapacityAh(t *testing.T) {
	f, err := os.Open("testdata/battery.json")
	if err != nil {
		t.Fatal(err)
	}
	battery := Battery{&SynkObject{}}
	data, err := io.ReadAll(f)
	if err != nil {
		t.Fatal(err)
	}
	err = json.Unmarshal(data, battery.SynkObject)
	if err != nil {
		t.Fatal(err)
	}
	cap, err := battery.CapacityAh()
	if err != nil {
		t.Fatal(err)
	}
	exp := 100.0
	if cap != exp {
		t.Errorf("expected %f, got %f", exp, cap)
	}
}
