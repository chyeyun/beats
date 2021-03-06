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

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package redis

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "redis", asset.ModuleFieldsPri, AssetRedis); err != nil {
		panic(err)
	}
}

// AssetRedis returns asset data.
// This is the base64 encoded zlib format compressed contents of module/redis.
func AssetRedis() string {
	return "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"
}
