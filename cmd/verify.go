/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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

	ym "github.com/aurawing/yottamapping"
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("verify called")
	},
}

var ethURL string
var contractAddr string
var fromIP string
var fromPort int
var fromUsername string
var fromPassword string
var fromDbname string
var toIP string
var toPort int
var toUsername string
var toPassword string
var toDbname string

func init() {
	rootCmd.AddCommand(verifyCmd)
	// Here you will define your flags and configuration settings.
	verifyCmd.Flags().StringVarP(&ethURL, "eth-url", "e", "https://ropsten.infura.io", "ethereum connection URL")
	verifyCmd.Flags().StringVarP(&contractAddr, "contract-addr", "c", "0x56E5B0BDaea3c33a8EC723E2E038453E84C14bA7", "contract address of yottamapping contract")
	verifyCmd.Flags().StringVarP(&fromIP, "from-ip", "i", "127.0.0.1", "IP address of source database")
	verifyCmd.Flags().IntVarP(&fromPort, "from-port", "p", 3306, "port number of source database")
	verifyCmd.Flags().StringVarP(&fromUsername, "from-username", "u", "root", "username of source database")
	verifyCmd.Flags().StringVarP(&fromPassword, "from-password", "w", "", "password of source database")
	verifyCmd.Flags().StringVarP(&fromDbname, "from-dbname", "n", "yottamapping", "name of source database")
	verifyCmd.Flags().StringVarP(&toIP, "to-ip", "j", "127.0.0.1", "IP address of destination database")
	verifyCmd.Flags().IntVarP(&toPort, "to-port", "q", 3306, "port number of destination database")
	verifyCmd.Flags().StringVarP(&toUsername, "to-username", "v", "root", "username of destination database")
	verifyCmd.Flags().StringVarP(&toPassword, "to-password", "x", "", "password of destination database")
	verifyCmd.Flags().StringVarP(&toDbname, "to-dbname", "o", "yottamapping", "name of destination database")

	mapper := ym.NewMapper(ethURL, contractAddr, fromIP, fromPort, fromUsername, fromPassword, fromDbname, toIP, toPort, toUsername, toPassword, toDbname)
	mapper.PullData()

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
