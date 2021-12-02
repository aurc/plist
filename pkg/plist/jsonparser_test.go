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
		expect    string
	}{
		{
			name:      "Test Array As Root",
			inputFile: "testdata/TestArray.plist",
			expect:    "testdata/want/TestArray.json",
		},
		{
			name:      "Test Dictionary As Root",
			inputFile: "testdata/TestDict.plist",
			expect:    "testdata/want/TestDict.json",
		},
		{
			name:      "Test Complex Example",
			inputFile: "testdata/Info.plist",
			expect:    "testdata/want/Info.json",
		},
		{
			name:      "Test Complex Large File",
			inputFile: "testdata/Power.plist",
			expect:    "testdata/want/Power.json",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			in, err := ioutil.ReadFile(test.inputFile)
			assert.NoError(t, err)
			be, err := ioutil.ReadFile(test.expect)
			assert.NoError(t, err)
			expected := string(be)
			got, err := Parse(in)
			assert.NoError(t, err)
			fmt.Println(got)

			assert.Equal(t, expected, got)
		})
	}
}
