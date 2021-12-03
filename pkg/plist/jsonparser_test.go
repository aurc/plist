/*
Copyright © 2021 Aurelio Calegari

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package plist

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		config    *Config
		expect    string
	}{
		{
			name:      "Test Array As Root to Json",
			inputFile: "testdata/TestArraySimple.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "testdata/want/TestArraySimple.json",
		},
		{
			name:      "Test Simple Array As Root to Json",
			inputFile: "testdata/TestArray.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "testdata/want/TestArray.json",
		},
		{
			name:      "Test Dictionary As Root to Json",
			inputFile: "testdata/TestDict.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "testdata/want/TestDict.json",
		},
		{
			name:      "Test Complex Example to Json",
			inputFile: "testdata/Info.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "testdata/want/Info.json",
		},
		{
			name:      "Test Complex Large File to Json",
			inputFile: "testdata/Power.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "testdata/want/Power.json",
		},
		{
			name:      "Test Simple Array As Root to Json with High Fidelity",
			inputFile: "testdata/TestArraySimple.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: true,
				Beatify:      false,
			},
			expect: "testdata/want/TestArraySimpleHF.json",
		},
		{
			name:      "Test Array As Root to Json with High Fidelity",
			inputFile: "testdata/TestArray.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: true,
				Beatify:      false,
			},
			expect: "testdata/want/TestArrayHF.json",
		},
		{
			name:      "Test Dictionary As Root to Json with High Fidelity, pretty print",
			inputFile: "testdata/TestDict.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: true,
				Beatify:      false,
			},
			expect: "testdata/want/TestDictHF.json",
		},
		{
			name:      "Test Array As Root to Yaml",
			inputFile: "testdata/TestArray.plist",
			config: &Config{
				Target:       Yaml,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "testdata/want/TestArray.yaml",
		},
		{
			name:      "Test Dictionary As Root to Yaml",
			inputFile: "testdata/TestDict.plist",
			config: &Config{
				Target:       Yaml,
				HighFidelity: false,
			},
			expect: "testdata/want/TestDict.yaml",
		},
		{
			name:      "Test Array As Root to Yaml with high fidelity",
			inputFile: "testdata/TestArray.plist",
			config: &Config{
				Target:       Yaml,
				HighFidelity: true,
			},
			expect: "testdata/want/TestArrayHF.yaml",
		},
		{
			name:      "Test Dictionary As Root to Yaml with High Fidelity",
			inputFile: "testdata/TestDict.plist",
			config: &Config{
				Target:       Yaml,
				HighFidelity: true,
			},
			expect: "testdata/want/TestDictHF.yaml",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in, err := ioutil.ReadFile(test.inputFile)
			assert.NoError(t, err)
			be, err := ioutil.ReadFile(test.expect)
			assert.NoError(t, err)
			expected := string(be)
			got, err := Parse(in, test.config)
			assert.NoError(t, err)
			gotStr := string(got)
			fmt.Println(gotStr)
			assert.NoError(t, err)
			assert.Equal(t, expected, gotStr)
		})
	}
}
