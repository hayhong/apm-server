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

package error

import (
	"testing"

	"github.com/elastic/apm-server/config"
	"github.com/elastic/apm-server/tests/loader"
)

func BenchmarkEventWithFileLoading(b *testing.B) {
	processor := NewProcessor()
	for i := 0; i < b.N; i++ {
		data, _ := loader.LoadValidData("error")
		err := processor.Validate(data)
		if err != nil {
			panic(err)
		}

		payload, err := processor.Decode(data)
		if err != nil {
			b.Fatalf("Error: %v", err)
		}
		payload.Transform(config.Config{})
	}
}

func BenchmarkEventFileLoadingOnce(b *testing.B) {
	processor := NewProcessor()
	data, _ := loader.LoadValidData("error")
	for i := 0; i < b.N; i++ {
		err := processor.Validate(data)
		if err != nil {
			panic(err)
		}

		payload, err := processor.Decode(data)
		if err != nil {
			b.Fatalf("Error: %v", err)
		}
		payload.Transform(config.Config{})
	}
}