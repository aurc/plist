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
package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pretty bool

// jsonCmd represents the json command
var jsonCmd = &cobra.Command{
	Use:   "json",
	Short: "Converts plist into JSON",
	Long:  `Outputs a JSON format payload converted from the given input plist.`,
	Run: func(cmd *cobra.Command, args []string) {
		in, err := ReadInput(input)
		if err != nil {
			panic(err)
		}
		output, err := Convert(in, &Config{
			Target:       Json,
			HighFidelity: highFidelity,
			Beatify:      pretty,
		})
		if err != nil {
			panic(err)
		}
		fmt.Print(string(output))
	},
}

func init() {
	rootCmd.AddCommand(jsonCmd)

	jsonCmd.PersistentFlags().BoolVarP(&pretty, "pretty-print", "p", false,
		"Pretty print (indent) output , e.g. --pretty true")

}
