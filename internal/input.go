package internal

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadInput(input string) ([]byte, error) {
	info, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}
	if info.Mode()&os.ModeCharDevice != 0 || info.Size() <= 0 {
		if input == "" {
			return nil, fmt.Errorf("bad input")
		} else {
			inputPayload, err := ioutil.ReadFile(input)
			if err != nil {
				return nil, err
			}
			return inputPayload, nil
		}
	}

	reader := bufio.NewReader(os.Stdin)

	inputPayload, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return inputPayload, nil
}
