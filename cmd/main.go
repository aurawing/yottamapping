package main

import (
	"fmt"

	ym "github.com/aurawing/yottamapping"

	dai "github.com/aurawing/yottamapping/dai"
)

func main() {
	_ = dai.New("127.0.0.1", 3306, "root", "admin", "test")
	fmt.Println(ym.EthAddrToName("0xe41d2489571d322189246dafa5ebde1f4699f498"))
}
