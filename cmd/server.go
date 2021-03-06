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

	"github.com/funayman/aomori-library/controller"
	"github.com/funayman/aomori-library/controller/admin"
	"github.com/funayman/aomori-library/router"
	"github.com/funayman/aomori-library/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	port      int
	ip        string
	withAdmin bool
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "run in server mode",
	// Long: `This is the log desc`,
	PreRun: func(cmd *cobra.Command, args []string) {
		port = viper.GetInt("server.port")
		ip = viper.GetString("server.ip")
		withAdmin = viper.GetBool("server.admin")

		fmt.Println("loading controllers")
		controller.Load()

		if withAdmin {
			admin.Load()
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("starting server...")
		server.Start(router.Instance(), server.Server{Port: port})
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	serverCmd.Flags().Int("port", 8081, "port for the server to run on (default 8081)")
	serverCmd.Flags().String("ip", "", "IP address for the server (default is localhost)")
	serverCmd.Flags().BoolP("admin", "", false, "allows all admin routes/API calls to be loaded")

	viper.BindPFlag("server.port", serverCmd.Flags().Lookup("port"))
	viper.BindPFlag("server.ip", serverCmd.Flags().Lookup("ip"))
	viper.BindPFlag("server.admin", serverCmd.Flags().Lookup("admin"))
}
