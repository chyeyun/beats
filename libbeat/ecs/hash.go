// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package ecs

// The hash fields represent different bitwise hash algorithms and their
// values.
// Field names for common hashes (e.g. MD5, SHA1) are predefined. Add fields
// for other hashes by lowercasing the hash algorithm name and using underscore
// separators as appropriate (snake case, e.g. sha3_512).
// Note that this fieldset is used for common hashes that may be computed over
// a range of generic bytes. Entity-specific hashes such as ja3 or imphash are
// placed in the fieldsets to which they relate (tls and pe, respectively).
type Hash struct {
	// MD5 hash.
	Md5 string `ecs:"md5"`

	// SHA1 hash.
	Sha1 string `ecs:"sha1"`

	// SHA256 hash.
	Sha256 string `ecs:"sha256"`

	// SHA512 hash.
	Sha512 string `ecs:"sha512"`

	// SSDEEP hash.
	Ssdeep string `ecs:"ssdeep"`
}
