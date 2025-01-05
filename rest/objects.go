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
)

// SynkObject is an abstraction of JSON objects returned by the SunSynk REST API. Concrete objects like
// Grid, Load or Battery encapsulate an instance of SynkObject that maps attributes of the object to values.
// The SynkObject type is "read only" and types derived from it should not be updated.

// Note: The Inverter abstraction does not encapsulate an instance of this type. The Inverter type is
// special since it can be updated (e.g. to set the battery discharge threshold.)
type SynkObject map[string]any

// Get returns the value of an attribute for an object returned by the SunSynk REST API.
func (s *SynkObject) Get(key string) (any, bool) {
	v, ok := (*s)[key]
	return v, ok
}

// String returns a JSON representation of an SynkObject (inverter, battery, load, etc.)
func (s SynkObject) String() string {
	m, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(m)
}
