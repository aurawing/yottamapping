package cmd

import (
	ym "github.com/aurawing/yottamapping"
	"github.com/spf13/cobra"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "从同步数据库拉取并校验数据",
	Long:  `当上游的同步数据库增量获取映射数据并入库后，本程序将同步这部分数据，同时在以太坊主链上校验其真实性，校验通过的数据会写入本地数据库`,
	Run: func(cmd *cobra.Command, args []string) {
		mapper := ym.NewMapper1(etherscanAPIURL, ethURL, ytaContractAddr, mapContractAddr, fromIP, fromPort, fromUsername, fromPassword, fromDbname, toIP, toPort, toUsername, toPassword, toDbname)
		mapper.PullData()
	},
}

var etherscanAPIURL string
var ethURL string
var ytaContractAddr string
var mapContractAddr string
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
	verifyCmd.Flags().StringVarP(&etherscanAPIURL, "api-url", "a", "http://ropsten.etherscan.io/api", "etherscan API URL")
	verifyCmd.Flags().StringVarP(&ethURL, "eth-url", "e", "https://ropsten.infura.io", "ethereum connection URL")
	verifyCmd.Flags().StringVarP(&ytaContractAddr, "yta-contract-addr", "y", "0x103b8c9e072193F2D0E3B77cEfb6d5DC9294d0d2", "contract address of yottamapping contract")
	verifyCmd.Flags().StringVarP(&mapContractAddr, "map-contract-addr", "m", "0x56E5B0BDaea3c33a8EC723E2E038453E84C14bA7", "contract address of yottacoin contract")
	verifyCmd.Flags().StringVarP(&fromIP, "from-ip", "i", "rm-2ze0y94f9up0nyhvrto.mysql.rds.aliyuncs.com", "IP address of source database")
	verifyCmd.Flags().IntVarP(&fromPort, "from-port", "p", 3306, "port number of source database")
	verifyCmd.Flags().StringVarP(&fromUsername, "from-username", "u", "mapping", "username of source database")
	verifyCmd.Flags().StringVarP(&fromPassword, "from-password", "w", "Yotta123", "password of source database")
	verifyCmd.Flags().StringVarP(&fromDbname, "from-dbname", "n", "test_mapping", "name of source database")
	verifyCmd.Flags().StringVarP(&toIP, "to-ip", "j", "127.0.0.1", "IP address of destination database")
	verifyCmd.Flags().IntVarP(&toPort, "to-port", "q", 3306, "port number of destination database")
	verifyCmd.Flags().StringVarP(&toUsername, "to-username", "v", "root", "username of destination database")
	verifyCmd.Flags().StringVarP(&toPassword, "to-password", "x", "", "password of destination database")
	verifyCmd.Flags().StringVarP(&toDbname, "to-dbname", "o", "test", "name of destination database")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
