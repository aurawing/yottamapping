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

// transferCmd represents the transfer command
var transferCmd = &cobra.Command{
	Use:   "transfer",
	Short: "迁移用户账号",
	Long:  `将用户的ERC20账号迁移至YottaChain主网，包括账号创建、代币迁移以及锁仓规则绑定`,
	Run: func(cmd *cobra.Command, args []string) {
		mapper := ym.NewMapper2(dbIP, dbPort, dbUsername, dbPassword, dbName, eosURL, adminAccount, adminPK, lockAccount, lockPK, operatorAccount, operatorPK, userPK, cpuStake, netStake)
		mapper.CreateAccountAndTransfer()
	},
}

var dbIP string
var dbPort int
var dbUsername string
var dbPassword string
var dbName string
var eosURL string
var adminAccount string
var adminPK string
var lockAccount string
var lockPK string
var operatorAccount string
var operatorPK string
var userPK string
var cpuStake int64
var netStake int64

func init() {
	rootCmd.AddCommand(transferCmd)
	// Here you will define your flags
	// Here you will define your flags and configuration settings.
	setParameters(transferCmd)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setParameters(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&dbIP, "dbip", "i", "127.0.0.1", "IP address of local database")
	cmd.Flags().IntVarP(&dbPort, "dbport", "p", 3306, "port number of local database")
	cmd.Flags().StringVarP(&dbUsername, "dbusername", "u", "root", "username of local database")
	cmd.Flags().StringVarP(&dbPassword, "dbpassword", "w", "", "password of local database")
	cmd.Flags().StringVarP(&dbName, "dbname", "n", "test", "name of local database")
	cmd.Flags().StringVarP(&eosURL, "eos-url", "s", "http://47.112.119.28:8888", "eos connection URL")
	cmd.Flags().StringVarP(&adminAccount, "admin-account", "a", "yottacreator", "admin account name")
	cmd.Flags().StringVarP(&adminPK, "admin-pk", "b", "5Kkvt7wy8S3PNJ7LeZxKtFh7ZmNbPSqGmeuLFcbttJ1GhTcZ1Pi", "admin privatekey")
	cmd.Flags().StringVarP(&lockAccount, "lock-account", "c", "lockadminusr", "lock account name")
	cmd.Flags().StringVarP(&lockPK, "lock-pk", "d", "5JcDH48njDbUQLu1R8SWwKsfWLnqBpWXDDiCgxFC3hioDuwLhVx", "lock privatekey")
	cmd.Flags().StringVarP(&operatorAccount, "operator-account", "e", "yottaairdrop", "operator account name")
	cmd.Flags().StringVarP(&operatorPK, "operator-pk", "f", "5K9dhuEfJeMxAB6x8fMJUuw3ntAJoDcVz59SaZVTrKaTmbLAjyh", "operator privatekey")
	cmd.Flags().StringVarP(&userPK, "user-pk", "g", "5KBDZghkb7WQ7jfQoreMxsWCcZi5WYg5HjDfMqHmgNG5pBX1WT9", "user privatekey")
	cmd.Flags().Int64VarP(&cpuStake, "cpu-stake", "k", 40000, "stake for cpu")
	cmd.Flags().Int64VarP(&netStake, "net-stake", "j", 10000, "stake for net")
}
