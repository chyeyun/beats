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

// The OS fields contain information about the operating system.
type Os struct {
	// Use the `os.type` field to categorize the operating system into one of
	// the broad commercial families.
	// One of these following values should be used (lowercase): linux, macos,
	// unix, windows.
	// If the OS you're dealing with is not in the list, the field should not
	// be populated. Please let us know by opening an issue with ECS, to
	// propose its addition.
	Type string `ecs:"type"`

	// Operating system platform (such centos, ubuntu, windows).
	Platform string `ecs:"platform"`

	// Operating system name, without the version.
	Name string `ecs:"name"`

	// Operating system name, including the version or code name.
	Full string `ecs:"full"`

	// OS family (such as redhat, debian, freebsd, windows).
	Family string `ecs:"family"`

	// Operating system version as a raw string.
	Version string `ecs:"version"`

	// Operating system kernel version as a raw string.
	Kernel string `ecs:"kernel"`
}
