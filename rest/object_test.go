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
	"testing"
)

func TestExtractOk(t *testing.T) {
	o := SynkObject{}
	o["foo"] = 3
	o["bar"] = 7
	o["hello"] = "world"
	keys := []string{"foo", "hello"}
	subset, err := o.ExtractKeys(keys)
	if err != nil {
		t.Errorf("test failed with reason %v", err)
	}
	if len(*subset) != 2 {
		t.Errorf("Expected 2 attributes, but got %d", len(*subset))
	}
	if (*subset)["foo"] != 3 {
		t.Errorf("Expected 3, but got %v", (*subset)["foo"])
	}
	if (*subset)["hello"] != "world" {
		t.Errorf("Expected hello, but got %v", (*subset)["hello"])
	}
}

func TestNoSuchKey(t *testing.T) {
	o := SynkObject{"foo": 3}
	subset, err := o.ExtractKeys([]string{"bar"})
	if err == nil {
		t.Error("expected to get an error")
	}
	if subset != nil {
		t.Error("no SynkObject should have been returned")
	}
}

func TestGet(t *testing.T) {
	o := SynkObject{"foo": 3, "bar": "baz"}
	v, ok := o.Get("foo")
	if !ok {
		t.Error("Couldn't retrieve key foo")
	}
	if v != 3 {
		t.Errorf("Expected value 3, got %d", v)
	}
	_, ok = o.Get("baz")
	if ok {
		t.Error("Retrieved non-existent key baz")
	}
}

func TestString(t *testing.T) {
	// Test that the String method returns valid JSON
	s := SynkObject{"foo": 3, "bar": "baz"}.String()
	data := map[string]any{}
	err := json.Unmarshal([]byte(s), &data)
	if err != nil {
		t.Fatal(err)
	}
	if len(data) != 2 {
		t.Errorf("Expected 2 keys, got %d", len(data))
	}
	_, ok := data["foo"]
	if !ok {
		t.Error("Couldn't retrieve key foo")
	}
}
