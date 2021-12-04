package cmd

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJsonCMD(t *testing.T) {
	tests := []struct {
		name      string
		arguments []string
		expect    string
	}{
		{
			"Test Simple Json",
			[]string{"plist", "json", "-i", "../testdata/TestArray.plist"},
			"../testdata/want/TestArray.json",
		},
		{
			"Test Simple Yaml",
			[]string{"plist", "yaml", "-i", "../testdata/TestArray.plist"},
			"../testdata/want/TestArray.yaml",
		},
	}
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := captureOut(func() {
				os.Args = test.arguments
				Execute()
			})
			be, err := ioutil.ReadFile(test.expect)
			assert.NoError(t, err)
			expected := string(be)
			assert.Equal(t, expected, got)
		})
	}
}

func captureOut(test func()) string {
	old := os.Stdout // keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	test()

	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old // restoring the real stdout
	out := <-outC
	return out
}
