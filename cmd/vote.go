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
	ym "github.com/aurawing/yottamapping"
	"github.com/spf13/cobra"
)

// voteCmd represents the vote command
var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "平移投票",
	Long:  `将用户账号在ERC20上进行的投票平移至YottaCahin主网`,
	Run: func(cmd *cobra.Command, args []string) {
		mapper := ym.NewMapper2(dbIP, dbPort, dbUsername, dbPassword, dbName, eosURL, adminAccount, adminPK, lockAccount, lockPK, operatorAccount, operatorPK, userPK, cpuStake, netStake)
		mapper.Vote()
	},
}

func init() {
	rootCmd.AddCommand(voteCmd)
	// Here you will define your flags
	// Here you will define your flags and configuration settings.
	setParameters(voteCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
