/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	b64 "encoding/base64"

	"github.com/karnowsa/gologic"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// stopCmd represents the stop command
var (
	// Used for flags.
	force string

	stopCmd = &cobra.Command{
		Use:   "stop",
		Short: "Stop Weblogic Managed Server",
		Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Run: func(cmd *cobra.Command, args []string) {
			force, _ := cmd.Flags().GetBool("force")
			passwordBase64Decode, _ := b64.StdEncoding.DecodeString(viper.GetString("password"))
			var admin gologic.AdminServer = gologic.LoginAdminServer(
				viper.GetString("ip"),
				viper.GetInt("port"),
				viper.GetString("username"),
				string(passwordBase64Decode))
			admin.Stop(args, force)
		},
	}
)

func init() {
	rootCmd.AddCommand(stopCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// stopCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	stopCmd.Flags().BoolP("force", "f", false, "Force Stop Weblogic Servers")
}
