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

package uuidapikey

import (
	"reflect"
	"testing"
)

var tests = []struct {
	uuid   string
	apiKey string
}{
	{
		"f06815b9-8808-4b4e-b466-0eb2ad36c1ac",
		"3R6G5DS-240GJTE-2T6C3NJ-2PKDGDC",
	},
	{
		"79dc9766-91a1-4d4b-8f16-2bdaa30cb03c",
		"1WXS5V6-28T2KAB-27HCAYT-2HGSC1W",
	},
	{
		"d1756360-5da0-40df-9926-a76abff5601d",
		"38QARV0-1ET0G6Z-2CJD9VA-2ZZAR0X",
	},
	{
		"317e31c0-5f78-4ad4-aa0e-f273b807321f",
		"0RQWCE0-1FQGJPM-2N0XWKK-2W0ECGZ",
	},
	{
		"a8ef8666-8d1b-4934-b115-28f70be3b762",
		"2MEZ1K6-26HPJ9M-2RHAA7Q-05Y7DV2",
	},
}

func TestIsUUID(t *testing.T) {
	for _, test := range tests {
		uuidValid, err := IsUUID(test.uuid)
		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}
		if !reflect.DeepEqual(true, uuidValid) {
			t.Errorf("Got %v, expected %v", uuidValid, true)
		}
	}
}

func TestIsAPIKey(t *testing.T) {
	for _, test := range tests {
		apiKeyValid, err := IsAPIKey(test.apiKey)
		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}
		if !reflect.DeepEqual(true, apiKeyValid) {
			t.Errorf("Got %v, expected %v", apiKeyValid, true)
		}
	}
}

func TestToUUID(t *testing.T) {
	for _, test := range tests {
		createdUUID, err := ToUUID(test.apiKey)
		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}
		if !reflect.DeepEqual(test.uuid, createdUUID) {
			t.Errorf("Got %v, expected %v", createdUUID, test.uuid)
		}
	}
}

func TestToAPIKey(t *testing.T) {
	for _, test := range tests {
		createdAPIKey, err := ToAPIKey(test.uuid)
		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}
		if !reflect.DeepEqual(test.apiKey, createdAPIKey) {
			t.Errorf("Got %v, expected %v", createdAPIKey, test.apiKey)
		}
	}
}

func TestCheck(t *testing.T) {
	for _, test := range tests {
		checkValid, err := Check(test.uuid, test.apiKey)
		if err != nil {
			t.Errorf("Error occurred: %v", err)
		}
		if !reflect.DeepEqual(true, checkValid) {
			t.Errorf("Got %v, expected %v", checkValid, true)
		}
	}
}
