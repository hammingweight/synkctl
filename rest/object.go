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
	"fmt"
)

// SynkObject is an abstraction of JSON objects returned by the SunSynk REST API. Concrete objects like
// Grid, Load or Battery encapsulate an instance of SynkObject that maps attributes of the object to values.
// The SynkObject type is "read only" and types derived from it should not be updated.
//
// Note: The Inverter abstraction does not encapsulate an instance of this type. The Inverter type is
// special since it can be updated (e.g. to set the battery discharge threshold.)
type SynkObject map[string]any

// Get returns the value of an attribute for an object returned by the SunSynk REST API.
func (s *SynkObject) Get(key string) (any, bool) {
	v, ok := (*s)[key]
	return v, ok
}

// String returns a JSON representation of a SynkObject (grid, battery, load, etc.)
func (s SynkObject) String() string {
	m, err := json.MarshalIndent(s, "", "    ")
	if err != nil {
		panic(err)
	}
	return string(m)
}

// ExtractKeys returns a subset of the fields of a SynkObject
func (s *SynkObject) ExtractKeys(keys []string) (*SynkObject, error) {
	res := &SynkObject{}
	for _, k := range keys {
		v, ok := (*s)[k]
		if !ok {
			return nil, fmt.Errorf("no such key: %s", k)
		}
		(*res)[k] = v
	}
	return res, nil
}
