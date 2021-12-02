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

	"github.com/aurc/plist/internal"
	"github.com/aurc/plist/pkg/plist"
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
		output, err := plist.Parse(in, &plist.Config{
			Target:       plist.Yaml,
			HighFidelity: pretty,
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// yamlCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// yamlCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
