// Copyright 2016 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package tool

import (
	"crypto/sha1"
	"encoding/hex"

	"github.com/Unknwon/com"
	"github.com/satori/go.uuid"
)

// EncodeSHA1 encodes string to SHA1 hex value.
func EncodeSHA1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// NewSecretToekn generates and returns a random secret token based on SHA1.
func NewSecretToekn() string {
	return EncodeSHA1(uuid.NewV4().String())
}

// Int64sToStrings converts a slice of int64 to a slice of string.
func Int64sToStrings(ints []int64) []string {
	strs := make([]string, len(ints))
	for i := range ints {
		strs[i] = com.ToStr(ints[i])
	}
	return strs
}
