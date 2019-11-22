package eth

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/aurawing/yottamapping/dai"
	"github.com/aurawing/yottamapping/eth/ytc"
	"github.com/aurawing/yottamapping/eth/ytm"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

//Cli ethereum client
type Cli struct {
	client          *ethclient.Client
	ytaContractAddr common.Address
	mapContractAddr common.Address
	eventType       common.Hash
}

//FreezedLog matches freezed log structure
type FreezedLog struct {
	Sender  common.Address
	Balance *big.Int
	Data    []byte
}

//NewClient create a new ethereum client
func NewClient(url, ytaContractAddr, mapContractAddr string) *Cli {
	freezedSig := []byte("Freezed(address,uint256,bytes)")
	freezedSigHash := crypto.Keccak256Hash(freezedSig)
	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatalf("error when create ethereum client: %s\n", err.Error())
	}
	return &Cli{client: client, ytaContractAddr: common.HexToAddress(ytaContractAddr), mapContractAddr: common.HexToAddress(mapContractAddr), eventType: freezedSigHash}
}

//GetFrozenTime get frozen timestamp of one address
func (cli *Cli) GetFrozenTime(addr string) int64 {
	address := common.HexToAddress(addr)
	instance, err := ytc.NewYtc(address, cli.client)
	if err != nil {
		log.Fatalf("error when create instance of yottacoin stub: %s\n", err.Error())
	}
	timestamp, err := instance.GetFrozenTimestamp(nil, common.HexToAddress(addr))
	if err != nil {
		log.Fatalf("error when get frozen timestamp of address %s: %s\n", addr, err.Error())
	}
	return timestamp.Int64()
}

//GetFreezedLogs fetch all freezed logs
func (cli *Cli) GetFreezedLogs(from, to int) (map[string]*dai.Mapping, error) {
	mappings := make(map[string]*dai.Mapping)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: []common.Address{
			cli.mapContractAddr,
		},
	}
	logs, err := cli.client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatalf("error when filter logs: %s\n", err.Error())
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(ytm.YtmABI)))
	if err != nil {
		log.Fatalf("error when decode YottaCoin ABI: %s\n", err.Error())
	}
	for _, vLog := range logs {
		switch strings.ToLower(vLog.Topics[0].Hex()) {
		case strings.ToLower(cli.eventType.Hex()):
			var freezedEvent FreezedLog
			err := contractAbi.Unpack(&freezedEvent, "Freezed", vLog.Data)
			if err != nil {
				log.Fatalf("error when decode freezed event ABI: %s\n", err.Error())
			}
			ytaAccount, err := dai.EthAddrToName(freezedEvent.Sender.Hex())
			if err != nil {
				log.Fatalf("error when generate YTA account from ethereum address: %s\n", err.Error())
			}
			m := dai.NewMapping(vLog.TxHash.Hex(), int(vLog.BlockNumber), freezedEvent.Sender.Hex(), freezedEvent.Balance.String(), string(freezedEvent.Data), 0, "", 0, ytaAccount, "", 0, "", "", 0)
			mappings[m.TransactionHash] = m
		}
	}
	return mappings, nil
}
