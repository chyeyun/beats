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

package monitors

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/elastic/beats/v7/libbeat/beat"
	"github.com/elastic/beats/v7/libbeat/beat/events"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/common/fmtstr"
	"github.com/elastic/beats/v7/libbeat/processors/add_data_stream"
)

func TestPreProcessors(t *testing.T) {
	binfo := beat.Info{
		Beat:        "heartbeat",
		IndexPrefix: "heartbeat",
		Version:     "8.0.0",
	}
	tests := map[string]struct {
		settings           publishSettings
		expectedIndex      string
		expectedDatastream *add_data_stream.DataStream
		monitorType        string
		wantProc           bool
		wantErr            bool
	}{
		"no settings should yield no processor": {
			publishSettings{},
			"",
			nil,
			"browser",
			false,
			false,
		},
		"exact index should be used exactly": {
			publishSettings{Index: *fmtstr.MustCompileEvent("test")},
			"test",
			nil,
			"browser",
			true,
			false,
		},
		"data stream should be type-namespace-dataset": {
			publishSettings{
				DataStream: &add_data_stream.DataStream{
					Namespace: "myNamespace",
					Dataset:   "myDataset",
					Type:      "myType",
				},
			},
			"myType-myDataset-myNamespace",
			&add_data_stream.DataStream{
				Namespace: "myNamespace",
				Dataset:   "myDataset",
				Type:      "myType",
			},
			"myType",
			true,
			false,
		},
		"data stream should use defaults": {
			publishSettings{
				DataStream: &add_data_stream.DataStream{},
			},
			"synthetics-browser-default",
			&add_data_stream.DataStream{
				Namespace: "default",
				Dataset:   "browser",
				Type:      "synthetics",
			},
			"browser",
			true,
			false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			e := beat.Event{Meta: common.MapStr{}, Fields: common.MapStr{}}
			procs, err := preProcessors(binfo, tt.settings, tt.monitorType)
			if tt.wantErr == true {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)

			// If we're just setting event.dataset we only get the 1
			// otherwise we get a second add_data_stream processor
			if !tt.wantProc {
				require.Len(t, procs.List, 1)
				return
			}
			require.Len(t, procs.List, 2)

			_, err = procs.Run(&e)

			t.Run("index name should be set", func(t *testing.T) {
				require.NoError(t, err)
				require.Equal(t, tt.expectedIndex, e.Meta[events.FieldMetaRawIndex])
			})

			eventDs, err := e.GetValue("event.dataset")
			require.NoError(t, err)

			t.Run("event.dataset should always be present, preferring data_stream", func(t *testing.T) {
				dataset := tt.monitorType
				if tt.settings.DataStream != nil && tt.settings.DataStream.Dataset != "" {
					dataset = tt.settings.DataStream.Dataset
				}
				require.Equal(t, dataset, eventDs, "event.dataset be computed correctly")
				require.Regexp(t, regexp.MustCompile(`^.+`), eventDs, "should be a string > 1 char")
			})

			t.Run("event.data_stream", func(t *testing.T) {
				dataStreamRaw, _ := e.GetValue("data_stream")
				if tt.expectedDatastream != nil {
					dataStream := dataStreamRaw.(add_data_stream.DataStream)
					require.Equal(t, eventDs, dataStream.Dataset, "event.dataset be identical to data_stream.dataset")

					require.Equal(t, *tt.expectedDatastream, dataStream)
				}
			})
		})
	}
}
