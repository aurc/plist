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
		stdIn     string
	}{
		{
			"Test Simple Json",
			[]string{"plist", "json", "-i", "../testdata/TestArray.plist"},
			"../testdata/want/TestArray.json",
			"",
		},
		{
			"Test Simple Yaml",
			[]string{"plist", "yaml", "-i", "../testdata/TestArray.plist"},
			"../testdata/want/TestArray.yaml",
			"",
		},
		{
			"Test Simple Yaml via StdIn aka pipe",
			[]string{"plist", "yaml"},
			"../testdata/want/TestArray.yaml",
			"../testdata/TestArray.plist",
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
			}, test.stdIn)
			be, err := ioutil.ReadFile(test.expect)
			assert.NoError(t, err)
			expected := string(be)
			assert.Equal(t, expected, got)
		})
	}
}

func captureOut(test func(), stdInFile string) string {
	oldIn := os.Stdin
	oldOut := os.Stdout
	defer func() {
		os.Stdin = oldIn
		os.Stdout = oldOut
	}()
	r, w, _ := os.Pipe()
	os.Stdout = w
	if stdInFile != "" {
		in, err := os.Open(stdInFile)
		defer func() {
			_ = in.Close()
		}()
		if err != nil {
			panic(err)
		}
		os.Stdin = in
	}

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
	out := <-outC
	return out
}
