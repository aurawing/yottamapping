package main

import (
	"fmt"
	"io/ioutil"

	"github.com/aurawing/yottamapping/cmd"
)

func main() {
	cmd.Execute()
}

func main1() {
	encConfig, err := ioutil.ReadFile("C:\\.mapping_config")
	if err != nil {
		panic(fmt.Sprintf("读取配置文件出错：%s", err.Error()))
	}
	config, err := cmd.AES256GCMDecrypt(encConfig, "")
	if err != nil {
		panic(fmt.Sprintf(fmt.Sprintf("\n解密配置文件失败: %s", err.Error())))
	}

	err = ioutil.WriteFile("c:\\.mapping_config_dec", config, 0666)
	if err != nil {
		panic(fmt.Sprintf("\n写入配置文件错误：%s\n", err.Error()))
	}
	fmt.Printf("\n配置文件解密完毕：%s\n")
}

func main2() {
	config, err := ioutil.ReadFile("C:\\.mapping_config_dec")
	if err != nil {
		panic(fmt.Sprintf("读取配置文件出错：%s", err.Error()))
	}
	encConfig, err := cmd.AES256GCMEncrypt(config, "")
	if err != nil {
		panic(fmt.Sprintf(fmt.Sprintf("\n加密配置文件失败: %s", err.Error())))
	}

	err = ioutil.WriteFile("C:\\.mapping_config_enc", encConfig, 0666)
	if err != nil {
		panic(fmt.Sprintf("\n写入加密配置文件错误：%s\n", err.Error()))
	}
	fmt.Printf("\n配置文件加密完毕：%s\n")
}
