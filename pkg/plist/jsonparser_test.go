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
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in, err := ioutil.ReadFile(test.inputFile)
			assert.NoError(t, err)
			be, err := ioutil.ReadFile(test.expect)
			assert.NoError(t, err)
			expected := string(be)
			got, err := Parse(in, test.config)
			gotStr := string(got)
			fmt.Println(gotStr)
			assert.NoError(t, err)
			assert.Equal(t, expected, gotStr)
		})
	}
}
