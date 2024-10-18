// Copyright (C) 2019 SIESGSTarena

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package uuidapikey contains generator, validator, and converter that transforms UUIDs into human-readable Base32-Crockford encoded API Keys
package uuidapikey

// Dependencies
import (
	"encoding/json"
	"errors"
	guid "github.com/google/uuid"
	"regexp"
	"strconv"
	"strings"
)

// checkDashes to verify dashes in given UUID
func checkDashes(positions []int, uuid string) bool {
	test := true
	for index := range positions {
		chr := uuid[positions[index]]
		test = test && chr == '-'
	}
	return test
}

// IsUUID will verify given UUID
func IsUUID(uuid string) (bool, error) {
	if uuid == "" {
		return false, errors.New("The required parameter UUID is undefined")
	}
	//_, error := goid.GetUUIDFromString(uuid)
	err := guid.Validate(uuid)
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsAPIKey will verify given API Key
func IsAPIKey(apiKey string) (bool, error) {
	if apiKey == "" {
		return false, errors.New("The required parameter API Key is undefined")
	}
	apiKey = strings.ToUpper(apiKey)
	var re = regexp.MustCompile("-")
	apiKey = re.ReplaceAllString(apiKey, "")
	re = regexp.MustCompile("[0-9A-Z]{28}")
	return len(apiKey) == 28 && re.MatchString(apiKey), nil
}

// ToAPIKey will convert given UUID to API Key
func ToAPIKey(uuid string) (string, error) {
	if uuid == "" {
		return "", errors.New("The required parameter UUID is undefined")
	}
	uuidValid, err := IsUUID(uuid)
	if err != nil {
		return "", err
	}

	if uuidValid {
		var re = regexp.MustCompile("-")
		uuid = re.ReplaceAllString(uuid, "")
		s1 := uuid[0:8]
		s2 := uuid[8:16]
		s3 := uuid[16:24]
		s4 := uuid[24:32]
		// fmt.Println(s1, s2, s3, s4)
		n1, _ := strconv.ParseUint(s1, 16, 32)
		n2, _ := strconv.ParseUint(s2, 16, 32)
		n3, _ := strconv.ParseUint(s3, 16, 32)
		n4, _ := strconv.ParseUint(s4, 16, 32)
		// fmt.Println(n1, n2, n3, n4)
		e1 := Encode(n1)
		e2 := Encode(n2)
		e3 := Encode(n3)
		e4 := Encode(n4)
		return e1 + "-" + e2 + "-" + e3 + "-" + e4, nil
	}
	return "", errors.New("Invalid UUID string")
}

// ToUUID will convert given API Key to UUID
func ToUUID(apiKey string) (string, error) {
	if apiKey == "" {
		return "", errors.New("The required parameter API Key is undefined")
	}
	apiKeyValid, err := IsAPIKey(apiKey)
	if err != nil {
		return "", err
	}
	if apiKeyValid {
		var re = regexp.MustCompile("-")
		apiKey = re.ReplaceAllString(apiKey, "")
		e1 := apiKey[0:7]
		e2 := apiKey[7:14]
		e3 := apiKey[14:21]
		e4 := apiKey[21:28]
		// fmt.Println(e1, e2, e3, e4)
		d1 := Decode(e1)
		d2 := Decode(e2)
		d3 := Decode(e3)
		d4 := Decode(e4)
		// fmt.Println(d1, d2, d3, d4)
		d2a := d2[0:4]
		d2b := d2[4:8]
		d3a := d3[0:4]
		d3b := d3[4:8]
		return d1 + "-" + d2a + "-" + d2b + "-" + d3a + "-" + d3b + d4, nil
	}
	return "", errors.New("Invalid API Key string")
}

// Check will verify both UUID and API Key, with either given
func Check(uuid, apiKey string) (bool, error) {
	if uuid == "" {
		return false, errors.New("The required parameter UUID is undefined")
	}
	if apiKey == "" {
		return false, errors.New("The required parameter API Key is undefined")
	}

	uuidValid, err := IsUUID(uuid)
	if err != nil {
		return false, err
	}
	apiKeyValid, err := IsAPIKey(apiKey)
	if err != nil {
		return false, err
	}

	if uuidValid && apiKeyValid {
		uuidCheck, err := ToUUID(apiKey)
		if err != nil {
			return false, err
		}
		return uuidCheck == uuid, nil
	}
	return false, errors.New("Invalid UUID or API Key string")
}

// Create will create new UUID and API Key
func Create() (string, error) {
	uuid := guid.New().String()
	apiKey, err := ToAPIKey(uuid)
	if err != nil {
		return "", err
	}
	pair := map[string]string{"uuid": uuid, "apiKey": apiKey}
	jsonPair, err := json.Marshal(pair)
	if err != nil {
		return "", errors.New("Error creating a new pair of keys")
	}
	return string(jsonPair), nil
}
