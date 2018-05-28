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
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clusterCmd represents the cluster command
var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "Cluster status. Use `help cluster` for more commands",
	Long: `Cluster information

See help for more subcommands.`,
	Args: cobra.MinimumNArgs(1),
}

var clusterHealthCmd = &cobra.Command{
	Use:   "health",
	Short: "Show cluster health",
	Run: func(cmd *cobra.Command, args []string) {
		es, err := client.New(client.WithHost(viper.GetString("config")))
		if err != nil {
			log.Fatalf("Unable to establish connection")
		}

		res, err := es.Cluster.Health()
		fmt.Printf("Status: %s\n", res.Response.Status)

	},
}

func init() {
	rootCmd.AddCommand(clusterCmd)

	clusterCmd.AddCommand(clusterHealthCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// clusterCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// clusterCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
