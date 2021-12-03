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
package cmd

import (
	"fmt"

	"github.com/aurc/plist"

	"github.com/aurc/plist/internal"
	"github.com/spf13/cobra"
)

// yamlCmd represents the yaml command
var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Converts plist into YAML",
	Long:  `Outputs a YAML format payload converted from the given input plist.`,
	Run: func(cmd *cobra.Command, args []string) {
		in, err := internal.ReadInput(input)
		if err != nil {
			panic(err)
		}
		output, err := plist.Convert(in, &plist.Config{
			Target:       plist.Yaml,
			HighFidelity: highFidelity,
			Beatify:      false,
		})
		if err != nil {
			panic(err)
		}
		fmt.Print(string(output))
	},
}

func init() {
	rootCmd.AddCommand(yamlCmd)
}
