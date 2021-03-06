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

// The user fields describe information about the user that is relevant to the
// event.
// Fields can have one entry or multiple entries. If a user has more than one
// id, provide an array that includes all of them.
type User struct {
	// Unique identifier of the user.
	ID string `ecs:"id"`

	// Short name or login of the user.
	Name string `ecs:"name"`

	// User's full name, if available.
	FullName string `ecs:"full_name"`

	// User email address.
	Email string `ecs:"email"`

	// Unique user hash to correlate information for a user in anonymized form.
	// Useful if `user.id` or `user.name` contain confidential information and
	// cannot be used.
	Hash string `ecs:"hash"`

	// Name of the directory the user is a member of.
	// For example, an LDAP or Active Directory domain name.
	Domain string `ecs:"domain"`

	// Array of user roles at the time of the event.
	Roles []string `ecs:"roles"`
}
