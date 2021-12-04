package cmd

import (
	"os"
	"testing"

	"github.com/spf13/cobra"

	"github.com/stretchr/testify/assert"
)

func TestFileNotFoundPanic(t *testing.T) {
	tests := []struct {
		name    string
		cmdFunc func(cmd *cobra.Command, args []string)
	}{
		{
			"Bad Json Input",
			jsonCmd.Run,
		},
		{
			"Bad Yaml Input",
			yamlCmd.Run,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Panics(t, func() {
				input = ""
				test.cmdFunc(nil, nil)
			})
		})
	}
}

func TestBadPListFormat(t *testing.T) {
	tests := []struct {
		name      string
		cmdFunc   func(cmd *cobra.Command, args []string)
		inputFile string
	}{
		{
			"Bad Json Input",
			jsonCmd.Run,
			"../testdata/want/TestArray.json",
		},
		{
			"Bad Yaml Input",
			yamlCmd.Run,
			"../testdata/want/TestArray.yaml",
		},
	}
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Panics(t, func() {
				input = test.inputFile
				test.cmdFunc(nil, nil)
			})
		})
	}
}
