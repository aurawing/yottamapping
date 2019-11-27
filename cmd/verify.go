package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"syscall"

	ym "github.com/aurawing/yottamapping"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// verifyCmd represents the verify command
var verifyCmd = &cobra.Command{
	Use:   "verify",
	Short: "从同步数据库拉取并校验数据",
	Long:  `当上游的同步数据库增量获取映射数据并入库后，本程序将同步这部分数据，同时在以太坊主链上校验其真实性，校验通过的数据会写入本地数据库`,
	Run: func(cmd *cobra.Command, args []string) {
		config := new(Config)
		home, err := Home()
		if err != nil {
			panic(err.Error())
		}
		configPath := filepath.Join(home, ".mapping_config")
		if !Exists(configPath) {
			log.Println("配置文件不存在, 请先执行init子命令创建一个配置文件")
			return
		}

		encConfig, err := ioutil.ReadFile(configPath)
		if err != nil {
			panic(fmt.Sprintf("读取配置文件出错：%s", err.Error()))
		}
		fmt.Print("请输入配置文件加密密码: ")
		for {
			bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(bytePassword) == "" {
				fmt.Print("\n格式错误，请重新输入密码: ")
				continue
			}
			password := string(bytePassword)
			password = RemoveBlank(password)
			configJSON, err := AES256GCMDecrypt(encConfig, password)
			if err != nil {
				fmt.Printf("\n解密失败，请重新输入密码：")
				continue
			}
			err = json.Unmarshal(configJSON, config)
			if err != nil {
				panic(fmt.Sprintf("\n反序列化配置文件失败：%s\n", err.Error()))
			}
			break
		}
		fmt.Println()
		fmt.Println()
		mapper := ym.NewMapper1(config.EtherscanAPIURL, config.EthURL, config.YtaContractAddr, config.MapContractAddr, config.FromIP, int(config.FromPort), config.FromUsername, config.FromPassword, config.FromDbname, config.ToIP, int(config.ToPort), config.ToUsername, config.ToPassword, config.ToDbname, config.TxTimeGap)
		mapper.PullData()
	},
}

func init() {
	rootCmd.AddCommand(verifyCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// verifyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// verifyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
