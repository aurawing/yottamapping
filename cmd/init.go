package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "初始化配置文件",
	Long:  `对各种密钥和设置进行初始化，并生成配置文件`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		home, err := Home()
		if err != nil {
			panic(err.Error())
		}
		configPath := filepath.Join(home, ".mapping_config")
		var delConfig bool
		if Exists(configPath) {
			fmt.Print("配置文件已存在，是否重新进行设置(Y/N): ")
			for {
				yn, _ := reader.ReadString('\n')
				yn = RemoveBlank(yn)
				if yn == "Y" || yn == "y" {
					delConfig = true
					break
				} else if yn == "N" || yn == "n" {
					return
				} else {
					fmt.Print("请输入Y或N: ")
				}
			}
		}

		config := new(Config)
		var password string

		fmt.Println("\n###### 1. 以下为上游数据库相关设置 ######")
		fmt.Print("请输入上游同步数据库IP或域名（默认为rm-2ze0y94f9up0nyhvrto.mysql.rds.aliyuncs.com）: ")
		config.FromIP, _ = reader.ReadString('\n')
		config.FromIP = RemoveBlank(config.FromIP)
		if config.FromIP == "" {
			config.FromIP = "rm-2ze0y94f9up0nyhvrto.mysql.rds.aliyuncs.com"
		}

		fmt.Print("请输入上游同步数据库端口（默认为3306）: ")
		for {
			fromPortStr, _ := reader.ReadString('\n')
			fromPortStr = RemoveBlank(fromPortStr)
			if fromPortStr == "" {
				fromPortStr = "3306"
			}
			config.FromPort, err = strconv.ParseUint(fromPortStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if config.FromPort <= 0 || config.FromPort > 65535 {
				fmt.Print("格式错误，请重新输入: ")
			} else {
				break
			}
		}

		fmt.Print("请输入上游同步数据库用户名（默认为root）: ")
		config.FromUsername, _ = reader.ReadString('\n')
		config.FromUsername = RemoveBlank(config.FromUsername)
		if config.FromUsername == "" {
			config.FromUsername = "root"
		}

		fmt.Print("请输入上游同步数据库密码（默认为admin）: ")
		config.FromPassword, _ = reader.ReadString('\n')
		config.FromPassword = RemoveBlank(config.FromPassword)
		if config.FromPassword == "" {
			config.FromPassword = "admin"
		}

		fmt.Print("请输入上游同步数据库名（默认为mapping）: ")
		config.FromDbname, _ = reader.ReadString('\n')
		config.FromDbname = RemoveBlank(config.FromDbname)
		if config.FromDbname == "" {
			config.FromDbname = "mapping"
		}

		fmt.Println("\n###### 2. 以下为本地数据库相关设置 ######")
		fmt.Print("请输入本地数据库IP或域名（默认为127.0.0.1）: ")
		config.ToIP, _ = reader.ReadString('\n')
		config.ToIP = RemoveBlank(config.ToIP)
		if config.ToIP == "" {
			config.ToIP = "127.0.0.1"
		}

		fmt.Print("请输入本地数据库端口（默认为3306）: ")
		for {
			toPortStr, _ := reader.ReadString('\n')
			toPortStr = RemoveBlank(toPortStr)
			if toPortStr == "" {
				toPortStr = "3306"
			}
			config.ToPort, err = strconv.ParseUint(toPortStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if config.ToPort <= 0 || config.ToPort > 65535 {
				fmt.Print("格式错误，请重新输入: ")
			} else {
				break
			}
		}

		fmt.Print("请输入本地数据库用户名（默认为root）: ")
		config.ToUsername, _ = reader.ReadString('\n')
		config.ToUsername = RemoveBlank(config.ToUsername)
		if config.ToUsername == "" {
			config.ToUsername = "root"
		}

		fmt.Print("请输入本地数据库密码（默认为admin）: ")
		config.ToPassword, _ = reader.ReadString('\n')
		config.ToPassword = RemoveBlank(config.ToPassword)
		if config.ToPassword == "" {
			config.ToPassword = "admin"
		}

		fmt.Print("请输入本地数据库名（默认为mapping）: ")
		config.ToDbname, _ = reader.ReadString('\n')
		config.ToDbname = RemoveBlank(config.ToDbname)
		if config.ToDbname == "" {
			config.ToDbname = "mapping"
		}

		fmt.Println("\n###### 3. 以下为以太坊相关设置 ######")
		fmt.Print("请输入Etherscan API URL（默认为http://api-cn.etherscan.com/api）: ")
		config.EtherscanAPIURL, _ = reader.ReadString('\n')
		config.EtherscanAPIURL = RemoveBlank(config.EtherscanAPIURL)
		if config.EtherscanAPIURL == "" {
			config.EtherscanAPIURL = "http://api-cn.etherscan.com/api"
		}

		fmt.Print("请输入以太坊API URL（默认为https://mainnet.infura.io）: ")
		config.EthURL, _ = reader.ReadString('\n')
		config.EthURL = RemoveBlank(config.EthURL)
		if config.EthURL == "" {
			config.EthURL = "https://mainnet.infura.io"
		}

		fmt.Print("请输入YottaCoin合约地址（默认为0x5edc1a266e8b2c5e8086d373725df0690af7e3ea）: ")
		config.YtaContractAddr, _ = reader.ReadString('\n')
		config.YtaContractAddr = RemoveBlank(config.YtaContractAddr)
		if config.YtaContractAddr == "" {
			config.YtaContractAddr = "0x5edc1a266e8b2c5e8086d373725df0690af7e3ea"
		}

		fmt.Print("请输入YottaMapping合约地址（默认为0x56E5B0BDaea3c33a8EC723E2E038453E84C14bA7）: ")
		config.MapContractAddr, _ = reader.ReadString('\n')
		config.MapContractAddr = RemoveBlank(config.MapContractAddr)
		if config.MapContractAddr == "" {
			config.MapContractAddr = "0x56E5B0BDaea3c33a8EC723E2E038453E84C14bA7"
		}

		fmt.Println("\n###### 4. 以下为EOS相关设置 ######")
		fmt.Print("请输入EOS连接URL（默认为http://47.112.119.28:8888）: ")
		config.EosURL, _ = reader.ReadString('\n')
		config.EosURL = RemoveBlank(config.EosURL)
		if config.EosURL == "" {
			config.EosURL = "http://47.112.119.28:8888"
		}

		fmt.Print("请输入管理员账号（默认为yottacreator）: ")
		config.AdminAccount, _ = reader.ReadString('\n')
		config.AdminAccount = RemoveBlank(config.AdminAccount)
		if config.AdminAccount == "" {
			config.AdminAccount = "yottacreator"
		}

		fmt.Print("请输入管理员账号私钥: ")
		for {
			byteAdminPK, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(byteAdminPK) == "" {
				fmt.Print("\n格式错误，请重新输入: ")
				continue
			}

			fmt.Print("\n请再次输入管理员账号私钥: ")
			var byteAdminPK2 []byte
			for {
				byteAdminPK2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(byteAdminPK2) == "" {
					fmt.Print("\n格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(byteAdminPK) != string(byteAdminPK2) {
				fmt.Print("\n两次输入不一致，请重新输入管理员账号私钥: ")
				continue
			}
			config.AdminPK = string(byteAdminPK)
			config.AdminPK = RemoveBlank(config.AdminPK)
			break
		}

		fmt.Print("\n请输入锁定账号（默认为lockadminusr）: ")
		config.LockAccount, _ = reader.ReadString('\n')
		config.LockAccount = RemoveBlank(config.LockAccount)
		if config.LockAccount == "" {
			config.LockAccount = "lockadminusr"
		}

		fmt.Print("请输入锁定账号私钥: ")
		for {
			byteLockPK, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(byteLockPK) == "" {
				fmt.Print("\n格式错误，请重新输入: ")
				continue
			}

			fmt.Print("\n请再次输入锁定账号私钥: ")
			var byteLockPK2 []byte
			for {
				byteLockPK2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(byteLockPK2) == "" {
					fmt.Print("\n格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(byteLockPK) != string(byteLockPK2) {
				fmt.Print("\n两次输入不一致，请重新输入锁定账号私钥: ")
				continue
			}
			config.LockPK = string(byteLockPK)
			config.LockPK = RemoveBlank(config.LockPK)
			break
		}

		fmt.Print("\n请输入操作账号（默认为yottaairdrop）: ")
		config.OperatorAccount, _ = reader.ReadString('\n')
		config.OperatorAccount = RemoveBlank(config.OperatorAccount)
		if config.OperatorAccount == "" {
			config.OperatorAccount = "yottaairdrop"
		}

		fmt.Print("请输入操作账号私钥: ")
		for {
			byteOperatorPK, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(byteOperatorPK) == "" {
				fmt.Print("\n格式错误，请重新输入: ")
				continue
			}

			fmt.Print("\n请再次输入操作账号私钥: ")
			var byteOperatorPK2 []byte
			for {
				byteOperatorPK2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(byteOperatorPK2) == "" {
					fmt.Print("\n格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(byteOperatorPK) != string(byteOperatorPK2) {
				fmt.Print("\n两次输入不一致，请重新输入操作账号私钥: ")
				continue
			}
			config.OperatorPK = string(byteOperatorPK)
			config.OperatorPK = RemoveBlank(config.OperatorPK)
			break
		}

		fmt.Print("\n请输入用户账号临时私钥: ")
		for {
			byteUserPK, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(byteUserPK) == "" {
				fmt.Print("\n格式错误，请重新输入: ")
				continue
			}

			fmt.Print("\n请再次输入用户账号临时私钥: ")
			var byteUserPK2 []byte
			for {
				byteUserPK2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(byteUserPK2) == "" {
					fmt.Print("\n格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(byteUserPK) != string(byteUserPK2) {
				fmt.Print("\n两次输入不一致，请重新输入用户账号临时私钥: ")
				continue
			}
			config.UserPK = string(byteUserPK)
			config.UserPK = RemoveBlank(config.UserPK)
			break
		}

		fmt.Print("\n请输入CPU抵押额度（默认为40000）: ")
		for {
			cpuStakeStr, _ := reader.ReadString('\n')
			cpuStakeStr = RemoveBlank(cpuStakeStr)
			if cpuStakeStr == "" {
				cpuStakeStr = "40000"
			}
			config.CPUStake, err = strconv.ParseUint(cpuStakeStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if config.CPUStake <= 0 {
				fmt.Print("格式错误，请重新输入: ")
			} else {
				break
			}
		}

		fmt.Print("请输入带宽抵押额度（默认为10000）: ")
		for {
			netStakeStr, _ := reader.ReadString('\n')
			netStakeStr = RemoveBlank(netStakeStr)
			if netStakeStr == "" {
				netStakeStr = "10000"
			}
			config.NetStake, err = strconv.ParseUint(netStakeStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if config.NetStake <= 0 {
				fmt.Print("格式错误，请重新输入: ")
			} else {
				break
			}
		}

		fmt.Print("请输入RAM采购额度（默认为4096）: ")
		for {
			ramStakeStr, _ := reader.ReadString('\n')
			ramStakeStr = RemoveBlank(ramStakeStr)
			if ramStakeStr == "" {
				ramStakeStr = "4096"
			}
			config.RAMStake, err = strconv.ParseUint(ramStakeStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if config.RAMStake <= 0 {
				fmt.Print("格式错误，请重新输入: ")
			} else {
				break
			}
		}

		fmt.Print("请输入允许自动转账的额度上限（默认为5000000000）: ")
		for {
			balThresholdStr, _ := reader.ReadString('\n')
			balThresholdStr = RemoveBlank(balThresholdStr)
			if balThresholdStr == "" {
				balThresholdStr = "5000000000"
			}
			config.BalThreshold, err = strconv.ParseInt(balThresholdStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if config.BalThreshold <= 0 {
				fmt.Print("格式错误，请重新输入: ")
			} else {
				break
			}
		}

		fmt.Print("请输入配置文件加密密码(长度不小于10): ")
		for {
			bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(bytePassword) == "" || len(string(bytePassword)) < 10 {
				fmt.Print("\n格式错误，请重新输入: ")
				continue
			}

			fmt.Print("\n请再次输入配置文件加密密码: ")
			var bytePassword2 []byte
			for {
				bytePassword2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(bytePassword2) == "" || len(string(bytePassword2)) < 10 {
					fmt.Print("\n格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(bytePassword) != string(bytePassword2) {
				fmt.Print("\n两次输入不一致，请重新输入配置文件加密密码: ")
				continue
			}
			password = string(bytePassword)
			password = RemoveBlank(password)
			break
		}

		c, err := json.Marshal(config)
		if err != nil {
			panic(fmt.Sprintf("\n配置文件序列化JSON失败: %s", err.Error()))
		}

		encConfig, err := AES256GCMEncrypt(c, password)
		if err != nil {
			panic(fmt.Sprintf(fmt.Sprintf("\n加密配置文件失败: %s", err.Error())))
		}

		if delConfig {
			err := os.Remove(configPath)
			if err != nil {
				panic(err.Error())
			}
		}

		err = ioutil.WriteFile(configPath, encConfig, 0666)
		if err != nil {
			panic(fmt.Sprintf("\n写入配置文件错误：%s\n", err.Error()))
		}
		fmt.Printf("\n初始化完毕，已写入加密配置文件：%s\n", configPath)

		//fmt.Println(string(AES256GCMDecrypt(encConfig, password)))
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
