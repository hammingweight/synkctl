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

import "testing"

func TestUpdateOK(t *testing.T) {
	object := SynkObject{"foo": 3, "bar": "baz"}
	err := object.Update("foo", 5)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if object["foo"] != 5 {
		t.Errorf("Expected foo=5, but got %v", object["foo"])
	}
	err = object.Update("bar", "hello")
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if object["bar"] != "hello" {
		t.Errorf("Expected bar=baz, but got %v", object["bar"])
	}
}

func TestUpdateInvalidKey(t *testing.T) {
	object := SynkObject{"foo": 3, "bar": "baz"}
	err := object.Update("foobar", 5)
	if err == nil {
		t.Error("updated non-existent key")
	}
}

func TestUpdateWrongTypes(t *testing.T) {
	object := SynkObject{"foo": 3, "bar": "baz"}
	err := object.Update("foo", "string")
	if err == nil {
		t.Errorf("updated foo value with wrong type")
	}
	err = object.Update("bar", 12345678)
	if err == nil {
		t.Errorf("updated bar value with wrong type")
	}
}

func TestGetNotOk(t *testing.T) {
	object := SynkObject{"foo": 3, "bar": "baz"}
	v, ok := object.Get("key")
	if ok {
		t.Error("unexpectedly got value")
	}
	if v != nil {
		t.Errorf("Expected nil, got %v", v)
	}
}
