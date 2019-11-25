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

// voteCmd represents the vote command
var voteCmd = &cobra.Command{
	Use:   "vote",
	Short: "平移投票",
	Long:  `将用户账号在ERC20上进行的投票平移至YottaCahin主网`,
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
		mapper := ym.NewMapper2(config.ToIP, int(config.ToPort), config.ToUsername, config.ToPassword, config.ToDbname, config.EosURL, config.AdminAccount, config.AdminPK, config.LockAccount, config.LockPK, config.OperatorAccount, config.OperatorPK, config.UserPK, int64(config.CPUStake), int64(config.NetStake), config.BalThreshold)
		mapper.Vote()
	},
}

func init() {
	rootCmd.AddCommand(voteCmd)
	// Here you will define your flags
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// transferCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// transferCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
