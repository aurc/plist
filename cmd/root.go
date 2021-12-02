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
	"github.com/spf13/cobra"
)

var cfgFile string
var input string
var inputPayload []byte
var pretty bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "plist",
	Short: "Convert Apple's Property Lists",
	Long: `This tool converts Apple's Property List (.plist) inputs into several useful
formats, such as JSON, YAML and HTML.

It supports both a file name as input and a piped ('|') input which might be useful
on more involved shell scripts.

For example:
    ./plist json -i myfile.plist
    ./plist json | cat myfile.plist

For individual commands instructions run:
	./plist [command] -h
	./plist json -h
`,
	TraverseChildren: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVarP(&input, "input", "i", "",
		"Specifies a input file, e.g. --input myFile.plist")
	rootCmd.PersistentFlags().BoolVarP(&pretty, "pretty-print", "p", false,
		"Pretty print (indent) output , e.g. --pretty true")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}
