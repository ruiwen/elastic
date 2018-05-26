// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	// "fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var output string

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Saves the current configuration",
	Long: `Writes the current configuration to disk.

By default, writes to the file passed to -c/--config (default: ~/.elastic).
However, can also write to a different file if passed the -o/--output option`,
	Run: func(cmd *cobra.Command, args []string) {

		settings := viper.AllSettings()
		delete(settings, "config") // Don't self reference, since we're writing a config file

		settings_output, err := yaml.Marshal(&settings)
		if err != nil {
			log.Fatalf("Error marshalling settings to YAML: %s\n", err)
		}
		if output == "" {
			output = viper.GetString("config")
		}
		err = ioutil.WriteFile(output, settings_output, 0644)
		if err != nil {
			log.Fatalf("Failed to write config file to %s: %s\n", output, err)
		}
	},
}

func init() {
	rootCmd.AddCommand(saveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	saveCmd.Flags().StringVarP(&output, "output", "o", "", "Write configuration to this file")
}
