package yottamapping

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"

	dai "github.com/aurawing/yottamapping/dai"
	eostx "github.com/aurawing/yottamapping/eostx"
)

//NewMapper2 create a new mapper structure
func NewMapper2(dbIP string, dbPort int, dbUsername, dbPassword, dbName, eosURL, adminAccount, adminPK, lockAccount, lockPK, operatorAccount, operatorPK, userPK string, cpuStake, netStake int64) *Mapper {
	return &Mapper{
		to:  dai.New(dbIP, dbPort, dbUsername, dbPassword, dbName),
		etx: eostx.New(eosURL, adminAccount, adminPK, lockAccount, lockPK, operatorAccount, operatorPK, userPK, cpuStake, netStake),
	}
}

//CreateAccountAndTransfer create a new account on YTA mainnet and do airdrop
func (mapper *Mapper) CreateAccountAndTransfer() {
	mapper.to.BrowserAndModify(func(m *dai.Mapping) {
		var s string
		if len(m.Balance) <= 14 {
			s = "0"
		} else {
			s = m.Balance[0 : len(m.Balance)-14]
		}
		bal, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalf("transfer: error when convert balance string to int64: %s\n", err.Error())
		}
		if bal > 10000 {
			var input string
			fmt.Printf("地址%s余额为%d, 是否转账?(Y/N): ", m.EthAddress, bal)
			for {
				fmt.Scanln(&input)
				if input == "Y" || input == "y" {
					break
				} else if input == "N" || input == "n" {
					log.Printf("取消该账号的转账操作")
					return
				} else {
					fmt.Printf("请输入Y或N: ")
				}
			}
		}
		var needVote bool
		var status int
		if m.IsVote == 1 {
			needVote = true
			status = 2
		} else {
			needVote = false
			status = 9
		}
		var blockMap map[int]int64
		if m.BlockRule != "" {
			err := json.Unmarshal([]byte(m.BlockRule), &blockMap)
			if err != nil {
				log.Fatalf("transfer: error when unmarshal block rule to map: %s\n", err.Error())
			}
		}
		txid, err := mapper.etx.CreateAccountTx(m.YtaAccount, string(m.Param[1:]), bal, m.FrozenTime, needVote, blockMap)
		if err != nil {
			if strings.Contains(err.Error(), "Account name already exists") {
				log.Printf("transfer: account has been created: %s\n", m.YtaAccount)
			} else {
				log.Fatalf("transfer: error happens when create YTA account: %s\n", err.Error())
			}
		} else {
			log.Printf("transfer: create account %s successful, TXID: %s\n", m.YtaAccount, txid)
		}
		mapper.to.UpdateTxid1(m.TransactionHash, txid, status)
		log.Printf("transfer: update database successful, account: %s, TXID: %s\n", m.YtaAccount, txid)
	}, 1)
}

//Vote do voting
func (mapper *Mapper) Vote() {
	mapper.to.BrowserAndModify(func(m *dai.Mapping) {
		if m.TxID1 != "" && !mapper.etx.IfTransactioIrreversible(m.TxID1) {
			log.Printf("vote: transaction %s is not irreversible, skip\n", m.TxID1)
			return
		}
		var s string
		if len(m.Balance) <= 14 {
			s = "0"
		} else {
			s = m.Balance[0 : len(m.Balance)-14]
		}
		bal, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalf("vote: error when convert balance string to int64: %s\n", err.Error())
		}
		txid, err := mapper.etx.VoteTx(m.YtaAccount, string(m.Param[1:]), bal, m.FrozenTime, m.NodeAccount)
		if err != nil {
			if strings.Contains(err.Error(), "does not have signatures for it under a provided delay") {
				log.Printf("voting has been done: %s\n", m.YtaAccount)
			} else {
				log.Fatalf("error happens when create YTA account: %s\n", err.Error())
			}
		} else {
			log.Printf("account %s do voting successful, TXID: %s\n", m.YtaAccount, txid)
		}
		mapper.to.UpdateTxid2(m.TransactionHash, txid, 9)
		log.Printf("update database successful, account: %s, TXID: %s\n", m.YtaAccount, txid)
	}, 2)
}
