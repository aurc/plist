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
package plistparser

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name      string
		inputFile string
		config    *Config
		expect    string
		isStdIn   bool
	}{
		{
			name:      "Test Array As Root to Json",
			inputFile: "../../testdata/TestArraySimple.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "../../testdata/want/TestArraySimple.json",
		},
		{
			name:      "Test Array As Root to Json with StdIn",
			inputFile: "../../testdata/TestArraySimple.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect:  "../../testdata/want/TestArraySimple.json",
			isStdIn: true,
		},
		{
			name:      "Test Simple Array As Root to Json",
			inputFile: "../../testdata/TestArray.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "../../testdata/want/TestArray.json",
		},
		{
			name:      "Test Dictionary As Root to Json",
			inputFile: "../../testdata/TestDict.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "../../testdata/want/TestDict.json",
		},
		{
			name:      "Test App Info Plist as Json",
			inputFile: "../../testdata/Info.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "../../testdata/want/Info.json",
		},
		{
			name:      "Test App Info Plist as Yaml as high fidelity",
			inputFile: "../../testdata/Info.plist",
			config: &Config{
				Target:       Yaml,
				HighFidelity: true,
				Beatify:      false,
			},
			expect: "../../testdata/want/InfoHF.yaml",
		},
		{
			name:      "Test Complex Example to Json",
			inputFile: "../../testdata/Info.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "../../testdata/want/Info.json",
		},
		{
			name:      "Test Complex Large File to Json",
			inputFile: "../../testdata/Power.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "../../testdata/want/Power.json",
		},
		{
			name:      "Test Simple Array As Root to Json with High Fidelity",
			inputFile: "../../testdata/TestArraySimple.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: true,
				Beatify:      false,
			},
			expect: "../../testdata/want/TestArraySimpleHF.json",
		},
		{
			name:      "Test Array As Root to Json with High Fidelity",
			inputFile: "../../testdata/TestArray.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: true,
				Beatify:      false,
			},
			expect: "../../testdata/want/TestArrayHF.json",
		},
		{
			name:      "Test Dictionary As Root to Json with High Fidelity, pretty print",
			inputFile: "../../testdata/TestDict.plist",
			config: &Config{
				Target:       Json,
				HighFidelity: true,
				Beatify:      true,
			},
			expect: "../../testdata/want/TestDictHF.json",
		},
		{
			name:      "Test Array As Root to Yaml",
			inputFile: "../../testdata/TestArray.plist",
			config: &Config{
				Target:       Yaml,
				HighFidelity: false,
				Beatify:      false,
			},
			expect: "../../testdata/want/TestArray.yaml",
		},
		{
			name:      "Test Dictionary As Root to Yaml",
			inputFile: "../../testdata/TestDict.plist",
			config: &Config{
				Target:       Yaml,
				HighFidelity: false,
			},
			expect: "../../testdata/want/TestDict.yaml",
		},
		{
			name:      "Test Array As Root to Yaml with high fidelity",
			inputFile: "../../testdata/TestArray.plist",
			config: &Config{
				Target:       Yaml,
				HighFidelity: true,
			},
			expect: "../../testdata/want/TestArrayHF.yaml",
		},
		{
			name:      "Test Dictionary As Root to Yaml with High Fidelity",
			inputFile: "../../testdata/TestDict.plist",
			config: &Config{
				Target:       Yaml,
				HighFidelity: true,
			},
			expect: "../../testdata/want/TestDictHF.yaml",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			func() {
				inputFile := test.inputFile
				oldIn := os.Stdin
				defer func() {
					os.Stdin = oldIn
				}()
				if test.isStdIn {
					inputFile = ""
					in, err := os.Open(test.inputFile)
					if err != nil {
						panic(err)
					}
					defer func() {
						_ = in.Close()
					}()
					os.Stdin = in
				}
				in, err := ReadInput(inputFile)
				assert.NoError(t, err)
				be, err := ioutil.ReadFile(test.expect)
				assert.NoError(t, err)
				expected := string(be)
				got, err := Convert(in, test.config)
				assert.NoError(t, err)
				gotStr := string(got)
				assert.NoError(t, err)
				assert.Equal(t, expected, gotStr)
			}()
		})
	}
}
