package etherscan

import (
	"os"
	"testing"
)

var cli *Cli

func TestMain(m *testing.M) {
	cli = NewCli("https://api-ropsten.etherscan.io/api", "0x103b8c9e072193F2D0E3B77cEfb6d5DC9294d0d2")
	os.Exit(m.Run())
}

func TestMappingRuleID(t *testing.T) {
	mappings := cli.MappingRuleID("0xA17C6339d3C43afC40bd6982731F2CD39c3C00Fc", 999999999)
	t.Logf("find %d rules: %v\n", len(mappings), mappings)
}
