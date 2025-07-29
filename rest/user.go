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
	"context"
	"errors"
)

// User is a model of a user.
type User struct{ *SynkObject }

// User calls the SunSynk REST API to get the user's details.
func (synkClient *SynkClient) User(ctx context.Context) (*User, error) {
	path := []string{"user"}
	queryParams := map[string]string{"lan": "en"}
	o := &SynkObject{}
	err := synkClient.readAPIV1(ctx, o, queryParams, path...)
	return &User{o}, err
}

// ID returns the user's identifier.
func (user *User) User() (int, error) {
	v, ok := user.Get("idc")
	if ok {
		switch v := v.(type) {
		case float64:
			return int(v), nil
		case int:
			return v, nil
		}
	}
	return 0, errors.New("cannot read user's ID")
}
