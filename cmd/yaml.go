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

	"github.com/aurc/plist/pkg/plistparser"

	"github.com/spf13/cobra"
)

// yamlCmd represents the yaml command
var yamlCmd = &cobra.Command{
	Use:   "yaml",
	Short: "Converts plist into YAML",
	Long:  `Outputs a YAML format payload converted from the given input plist.`,
	Run: func(cmd *cobra.Command, args []string) {
		in, err := plistparser.ReadInput(input)
		if err != nil {
			panic(err)
		}
		output, err := plistparser.Convert(in, &plistparser.Config{
			Target:       plistparser.Yaml,
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
