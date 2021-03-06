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

// Geo fields can carry data about a specific location related to an event.
// This geolocation information can be derived from techniques such as Geo IP,
// or be user-supplied.
type Geo struct {
	// Longitude and latitude.
	Location string `ecs:"location"`

	// Two-letter code representing continent's name.
	ContinentCode string `ecs:"continent_code"`

	// Name of the continent.
	ContinentName string `ecs:"continent_name"`

	// Country name.
	CountryName string `ecs:"country_name"`

	// Region name.
	RegionName string `ecs:"region_name"`

	// City name.
	CityName string `ecs:"city_name"`

	// Country ISO code.
	CountryIsoCode string `ecs:"country_iso_code"`

	// Postal code associated with the location.
	// Values appropriate for this field may also be known as a postcode or ZIP
	// code and will vary widely from country to country.
	PostalCode string `ecs:"postal_code"`

	// Region ISO code.
	RegionIsoCode string `ecs:"region_iso_code"`

	// The time zone of the location, such as IANA time zone name.
	Timezone string `ecs:"timezone"`

	// User-defined description of a location, at the level of granularity they
	// care about.
	// Could be the name of their data centers, the floor number, if this
	// describes a local physical entity, city names.
	// Not typically used in automated geolocation.
	Name string `ecs:"name"`
}
