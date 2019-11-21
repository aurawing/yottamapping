package etherscan

import (
	"os"
	"testing"
)

var cli *Cli

func TestMain(m *testing.M) {
	cli = NewCli("http://api-cn.etherscan.com/api", "0x5edc1a266e8b2c5e8086d373725df0690af7e3ea")
	os.Exit(m.Run())
}

func TestMappingRuleID(t *testing.T) {
	mappings := cli.MappingRuleID("0x3b832892e7a59c19052d7B8Ca4Ef23C461a9BEAa", 999999999)
	t.Logf("find %d rules: %v\n", len(mappings), mappings)
}
