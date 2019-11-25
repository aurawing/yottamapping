package yottamapping

import (
	"github.com/aurawing/yottamapping/dai"
	"github.com/aurawing/yottamapping/eostx"
	"github.com/aurawing/yottamapping/eth"
	"github.com/aurawing/yottamapping/etherscan"
)

//Mapper structure
type Mapper struct {
	from         *dai.Dai
	to           *dai.Dai
	ethcli       *eth.Cli
	etherscanCli *etherscan.Cli
	etx          *eostx.EosTX
	balThreshold int64
}
