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
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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

		var fromIP string
		var fromPort uint64
		var fromUsername string
		var fromPassword string
		var fromDbname string
		var toIP string
		var toPort uint64
		var toUsername string
		var toPassword string
		var toDbname string
		var etherscanAPIURL string
		var ethURL string
		var ytaContractAddr string
		var mapContractAddr string
		var eosURL string
		var adminAccount string
		var adminPK string
		var lockAccount string
		var lockPK string
		var operatorAccount string
		var operatorPK string
		var userPK string
		var cpuStake uint64
		var netStake uint64
		var ramStake uint64
		var password string

		fmt.Println("###### 1. 以下为上游数据库相关设置 ######")
		fmt.Print("请输入上游同步数据库IP或域名（默认为rm-2ze0y94f9up0nyhvrto.mysql.rds.aliyuncs.com）: ")
		fromIP, _ = reader.ReadString('\n')
		fromIP = RemoveBlank(fromIP)
		if fromIP == "" {
			fromIP = "rm-2ze0y94f9up0nyhvrto.mysql.rds.aliyuncs.com"
		}

		fmt.Print("请输入上游同步数据库端口（默认为3306）: ")
		for {
			fromPortStr, _ := reader.ReadString('\n')
			fromPortStr = RemoveBlank(fromPortStr)
			if fromPortStr == "" {
				fromPortStr = "3306"
			}
			fromPort, err = strconv.ParseUint(fromPortStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if fromPort <= 0 || fromPort > 65535 {
				fmt.Print("格式错误，请重新输入: ")
			} else {
				break
			}
		}

		fmt.Print("请输入上游同步数据库用户名（默认为root）: ")
		fromUsername, _ = reader.ReadString('\n')
		fromUsername = RemoveBlank(fromUsername)
		if fromUsername == "" {
			fromUsername = "root"
		}

		fmt.Print("请输入上游同步数据库密码（默认为admin）: ")
		fromPassword, _ = reader.ReadString('\n')
		fromPassword = RemoveBlank(fromPassword)
		if fromPassword == "" {
			fromPassword = "admin"
		}

		fmt.Print("请输入上游同步数据库名（默认为mapping）: ")
		fromDbname, _ = reader.ReadString('\n')
		fromDbname = RemoveBlank(fromDbname)
		if fromDbname == "" {
			fromDbname = "mapping"
		}

		fmt.Println("###### 2. 以下为本地数据库相关设置 ######")
		fmt.Print("请输入本地数据库IP或域名（默认为127.0.0.1）: ")
		toIP, _ = reader.ReadString('\n')
		toIP = RemoveBlank(toIP)
		if toIP == "" {
			toIP = "127.0.0.1"
		}

		fmt.Print("请输入本地数据库端口（默认为3306）: ")
		for {
			toPortStr, _ := reader.ReadString('\n')
			toPortStr = RemoveBlank(toPortStr)
			if toPortStr == "" {
				toPortStr = "3306"
				break
			}
			toPort, err = strconv.ParseUint(toPortStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if toPort <= 0 || toPort > 65535 {
				fmt.Print("格式错误，请重新输入: ")
			} else {
				break
			}
		}

		fmt.Print("请输入本地数据库用户名（默认为root）: ")
		toUsername, _ = reader.ReadString('\n')
		toUsername = RemoveBlank(toUsername)
		if toUsername == "" {
			toUsername = "root"
		}

		fmt.Print("请输入本地数据库密码（默认为admin）: ")
		toPassword, _ = reader.ReadString('\n')
		toPassword = RemoveBlank(toPassword)
		if toPassword == "" {
			toPassword = "admin"
		}

		fmt.Print("请输入本地数据库名（默认为mapping）: ")
		toDbname, _ = reader.ReadString('\n')
		toDbname = RemoveBlank(toDbname)
		if toDbname == "" {
			toDbname = "mapping"
		}

		fmt.Println("###### 3. 以下为以太坊相关设置 ######")
		fmt.Print("请输入Etherscan API URL（默认为http://api-cn.etherscan.com/api）: ")
		etherscanAPIURL, _ = reader.ReadString('\n')
		etherscanAPIURL = RemoveBlank(etherscanAPIURL)
		if etherscanAPIURL == "" {
			etherscanAPIURL = "http://api-cn.etherscan.com/api"
		}

		fmt.Print("请输入以太坊API URL（默认为https://mainnet.infura.io）: ")
		ethURL, _ = reader.ReadString('\n')
		ethURL = RemoveBlank(ethURL)
		if ethURL == "" {
			ethURL = "https://mainnet.infura.io"
		}

		fmt.Print("请输入YottaCoin合约地址（默认为0x5edc1a266e8b2c5e8086d373725df0690af7e3ea）: ")
		ytaContractAddr, _ = reader.ReadString('\n')
		ytaContractAddr = RemoveBlank(ytaContractAddr)
		if ytaContractAddr == "" {
			ytaContractAddr = "0x5edc1a266e8b2c5e8086d373725df0690af7e3ea"
		}

		fmt.Print("请输入YottaMapping合约地址（默认为0x56E5B0BDaea3c33a8EC723E2E038453E84C14bA7）: ")
		mapContractAddr, _ = reader.ReadString('\n')
		mapContractAddr = RemoveBlank(mapContractAddr)
		if mapContractAddr == "" {
			mapContractAddr = "0x56E5B0BDaea3c33a8EC723E2E038453E84C14bA7"
		}

		fmt.Println("###### 4. 以下为EOS相关设置 ######")
		fmt.Print("请输入EOS连接URL（默认为http://47.112.119.28:8888）: ")
		eosURL, _ = reader.ReadString('\n')
		eosURL = RemoveBlank(eosURL)
		if eosURL == "" {
			eosURL = "http://47.112.119.28:8888"
		}

		fmt.Print("请输入管理员账号（默认为yottacreator）: ")
		adminAccount, _ = reader.ReadString('\n')
		adminAccount = RemoveBlank(adminAccount)
		if adminAccount == "" {
			adminAccount = "yottacreator"
		}

		fmt.Print("请输入管理员账号私钥: ")
		for {
			byteAdminPK, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(byteAdminPK) == "" {
				fmt.Print("格式错误，请重新输入: ")
				continue
			}

			fmt.Print("请再次输入管理员账号私钥: ")
			var byteAdminPK2 []byte
			for {
				byteAdminPK2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(byteAdminPK2) == "" {
					fmt.Print("格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(byteAdminPK) != string(byteAdminPK2) {
				fmt.Print("两次输入不一致，请重新输入管理员账号私钥: ")
				continue
			}
			adminPK = string(byteAdminPK)
			adminPK = RemoveBlank(adminPK)
			break
		}

		fmt.Print("请输入锁定账号（默认为lockadminusr）: ")
		lockAccount, _ = reader.ReadString('\n')
		lockAccount = RemoveBlank(lockAccount)
		if lockAccount == "" {
			lockAccount = "lockadminusr"
		}

		fmt.Print("请输入锁定账号私钥: ")
		for {
			byteLockPK, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(byteLockPK) == "" {
				fmt.Print("格式错误，请重新输入: ")
				continue
			}

			fmt.Print("请再次输入锁定账号私钥: ")
			var byteLockPK2 []byte
			for {
				byteLockPK2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(byteLockPK2) == "" {
					fmt.Print("格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(byteLockPK) != string(byteLockPK2) {
				fmt.Print("两次输入不一致，请重新输入锁定账号私钥: ")
				continue
			}
			lockPK = string(byteLockPK)
			lockPK = RemoveBlank(lockPK)
			break
		}

		fmt.Print("请输入操作账号（默认为yottaairdrop）: ")
		operatorAccount, _ = reader.ReadString('\n')
		operatorAccount = RemoveBlank(operatorAccount)
		if operatorAccount == "" {
			operatorAccount = "yottaairdrop"
		}

		fmt.Print("请输入操作账号私钥: ")
		for {
			byteOperatorPK, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(byteOperatorPK) == "" {
				fmt.Print("格式错误，请重新输入: ")
				continue
			}

			fmt.Print("请再次输入操作账号私钥: ")
			var byteOperatorPK2 []byte
			for {
				byteOperatorPK2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(byteOperatorPK2) == "" {
					fmt.Print("格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(byteOperatorPK) != string(byteOperatorPK2) {
				fmt.Print("两次输入不一致，请重新输入操作账号私钥: ")
				continue
			}
			operatorPK = string(byteOperatorPK)
			operatorPK = RemoveBlank(operatorPK)
			break
		}

		fmt.Print("请输入用户账号临时私钥: ")
		for {
			byteUserPK, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(byteUserPK) == "" {
				fmt.Print("格式错误，请重新输入: ")
				continue
			}

			fmt.Print("请再次输入用户账号临时私钥: ")
			var byteUserPK2 []byte
			for {
				byteUserPK2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(byteUserPK2) == "" {
					fmt.Print("格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(byteUserPK) != string(byteUserPK2) {
				fmt.Print("两次输入不一致，请重新输入用户账号临时私钥: ")
				continue
			}
			userPK = string(userPK)
			userPK = RemoveBlank(userPK)
			break
		}

		fmt.Print("请输入CPU抵押额度（默认为40000）: ")
		for {
			cpuStakeStr, _ := reader.ReadString('\n')
			cpuStakeStr = RemoveBlank(cpuStakeStr)
			if cpuStakeStr == "" {
				cpuStakeStr = "40000"
				break
			}
			cpuStake, err = strconv.ParseUint(cpuStakeStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if cpuStake <= 0 {
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
				break
			}
			netStake, err = strconv.ParseUint(netStakeStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if netStake <= 0 {
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
				ramStake = 4096
				break
			}
			ramStake, err = strconv.ParseUint(ramStakeStr, 10, 64)
			if err != nil {
				fmt.Print("格式错误，请重新输入: ")
			} else if ramStake <= 0 {
				fmt.Print("格式错误，请重新输入: ")
			} else {
				break
			}
		}

		fmt.Print("请输入配置文件加密密码: ")
		for {
			bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil || string(bytePassword) == "" {
				fmt.Print("格式错误，请重新输入: ")
				continue
			}

			fmt.Print("请再次输入配置文件加密密码: ")
			var bytePassword2 []byte
			for {
				bytePassword2, err = terminal.ReadPassword(int(syscall.Stdin))
				if err != nil || string(bytePassword2) == "" {
					fmt.Print("格式错误，请重新输入: ")
					continue
				}
				break
			}
			if string(bytePassword) != string(bytePassword2) {
				fmt.Print("两次输入不一致，请重新输入配置文件加密密码: ")
				continue
			}
			password = string(bytePassword)
			password = RemoveBlank(password)
			break
		}

		if delConfig {
			err := os.Remove(configPath)
			if err != nil {
				panic(err.Error())
			}
		}

		fmt.Println(adminPK, lockPK, operatorPK)
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

// Home returns the home directory for the executing user.
//
// This uses an OS-specific method for discovering the home directory.
// An error is returned if a home directory cannot be detected.
func Home() (string, error) {
	user, err := user.Current()
	if nil == err {
		return user.HomeDir, nil
	}

	// cross compile support

	if "windows" == runtime.GOOS {
		return homeWindows()
	}

	// Unix-like system, so just assume Unix
	return homeUnix()
}

func homeUnix() (string, error) {
	// First prefer the HOME environmental variable
	if home := os.Getenv("HOME"); home != "" {
		return home, nil
	}

	// If that fails, try the shell
	var stdout bytes.Buffer
	cmd := exec.Command("sh", "-c", "eval echo ~$USER")
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return "", err
	}

	result := strings.TrimSpace(stdout.String())
	if result == "" {
		return "", errors.New("blank output when reading home directory")
	}

	return result, nil
}

func homeWindows() (string, error) {
	drive := os.Getenv("HOMEDRIVE")
	path := os.Getenv("HOMEPATH")
	home := drive + path
	if drive == "" || path == "" {
		home = os.Getenv("USERPROFILE")
	}
	if home == "" {
		return "", errors.New("HOMEDRIVE, HOMEPATH, and USERPROFILE are blank")
	}

	return home, nil
}

//Exists 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

//IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

//IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

//RemoveBlank 去掉空白和换行符
func RemoveBlank(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	return str
}
