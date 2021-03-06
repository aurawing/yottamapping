package yottamapping

import (
	"encoding/json"
	"log"
	"time"

	dai "github.com/aurawing/yottamapping/dai"
)

//Verify verify correctness of data
func (mapper *Mapper) Verify(mappings map[string]*dai.Mapping) dai.Verifier {
	return func(m *dai.Mapping) bool {
		mv := mappings[m.TransactionHash]
		if mv == nil {
			log.Printf("!!! event with transaction hash %s is not exist\n", m.TransactionHash)
			return false
		}
		if m.BlockNumber != mv.BlockNumber {
			log.Printf("!!! block number in DB is not equal to event:%s %d %d\n", m.TransactionHash, m.BlockNumber, mv.BlockNumber)
			return false
		}
		if m.EthAddress != mv.EthAddress {
			log.Printf("!!! ERC20 address in DB is not equal to event: %s %s %s\n", m.TransactionHash, m.EthAddress, mv.EthAddress)
			return false
		}
		if m.Balance != mv.Balance {
			log.Printf("!!! balance in DB is not equal to event: %s %s %s\n", m.TransactionHash, m.Balance, mv.Balance)
			return false
		}
		if m.Param != mv.Param {
			log.Printf("!!! param in DB is not equal to event: %s %s %s\n", m.TransactionHash, m.Param, mv.Param)
			return false
		}
		ruleMap := mapper.etherscanCli.MappingRuleID(m.EthAddress, m.BlockNumber)
		if len(ruleMap) > 0 {
			ruleMapJSON, err := json.Marshal(ruleMap)
			if err != nil {
				log.Fatalf("error when marshal rule map to json: %s\n", err.Error())
			}
			m.BlockRule = string(ruleMapJSON)
			log.Printf("fetched rule map for etheraddr %s: %s\n", m.EthAddress, m.BlockRule)
		}
		timestamp := mapper.ethcli.GetFrozenTimeRetries(m.EthAddress, 5)
		if timestamp > 1574319600 {
			m.FrozenTime = timestamp
			log.Printf("fetched frozen time for etheraddr %s: %d\n", m.EthAddress, m.FrozenTime)
		}
		if mapper.txTimeGap > 0 {
			time.Sleep(time.Duration(mapper.txTimeGap) * time.Millisecond)
		}
		return true
	}
}
