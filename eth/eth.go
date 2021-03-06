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

//TransferLog transfer event
type TransferLog struct {
	From  common.Address
	To    common.Address
	Value *big.Int
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
func (cli *Cli) GetFrozenTime(addr string) (int64, error) {
	instance, err := ytc.NewYtc(cli.ytaContractAddr, cli.client)
	if err != nil {
		//log.Fatalf("error when create instance of yottacoin stub: %s\n", err.Error())
		return 0, err
	}
	timestamp, err := instance.GetFrozenTimestamp(nil, common.HexToAddress(addr))
	if err != nil {
		//log.Fatalf("error when get frozen timestamp of address %s: %s\n", addr, err.Error())
		return 0, err
	}
	return timestamp.Int64(), nil
}

//GetFrozenTimeRetries get frozen timestamp of one address with retries
func (cli *Cli) GetFrozenTimeRetries(addr string, retries int) int64 {
	if retries == 0 {
		log.Fatalf("error when get frozen timestamp of address: %s\n", addr)
	}
	t, err := cli.GetFrozenTime(addr)
	if err != nil {
		log.Printf("Warning: error when get frozen time: %s\n", err.Error())
		retries--
		return cli.GetFrozenTimeRetries(addr, retries)
	}
	return t
}

//GetFreezedLogsAll get all frozen logs
func (cli *Cli) GetFreezedLogsAll(from, to int) map[string]*dai.Mapping {
	mappings := make(map[string]*dai.Mapping)
	for {
		end := from + 100
		if end >= to {
			end = to
		}
		m := cli.GetFreezedLogsRetry(from, end, 5)
		log.Printf("get freezed logs from %d to %d: %d logs\n", from, end, len(m))
		for k, v := range m {
			mappings[k] = v
		}
		from = end + 1
		if from > to {
			break
		}
	}
	return mappings
}

//GetFreezedLogsRetry fetch all freezed logs with retry
func (cli *Cli) GetFreezedLogsRetry(from, to, retry int) map[string]*dai.Mapping {
	if retry == 0 {
		log.Fatalln("error when get freezed logs")
	}
	m, err := cli.GetFreezedLogs(from, to)
	if err != nil {
		log.Printf("Warning: error when get freezed logs: %s\n", err.Error())
		retry--
		return cli.GetFreezedLogsRetry(from, to, retry)
	}
	return m
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
		return nil, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(ytm.YtmABI)))
	if err != nil {
		return nil, err
	}
	for _, vLog := range logs {
		switch strings.ToLower(vLog.Topics[0].Hex()) {
		case strings.ToLower(cli.eventType.Hex()):
			var freezedEvent FreezedLog
			err := contractAbi.Unpack(&freezedEvent, "Freezed", vLog.Data)
			if err != nil {
				return nil, err
			}
			ytaAccount, err := dai.EthAddrToName(freezedEvent.Sender.Hex())
			if err != nil {
				return nil, err
			}
			m := dai.NewMapping(vLog.TxHash.Hex(), int(vLog.BlockNumber), freezedEvent.Sender.Hex(), freezedEvent.Balance.String(), string(freezedEvent.Data), 0, "", 0, ytaAccount, "", 0, "", "", 0)
			mappings[m.TransactionHash] = m
		}
	}
	return mappings, nil
}

//GetTransferLogsAll get all transfer logs
func (cli *Cli) GetTransferLogsAll(from, to int) []*TransferLog {
	logs := make([]*TransferLog, 0)
	for {
		end := from + 999
		if end >= to {
			end = to
		}
		m := cli.GetTransferLogsRetry(from, end, 5)
		log.Printf("get transfer logs from %d to %d: %d logs\n", from, end, len(m))
		logs = append(logs, m...)
		from = end + 1
		if from > to {
			break
		}
	}
	return logs
}

//GetTransferLogsRetry fetch all transfer logs with retry
func (cli *Cli) GetTransferLogsRetry(from, to, retry int) []*TransferLog {
	if retry == 0 {
		log.Fatalln("error when get transfer logs")
	}
	m, err := cli.GetTransferLogs(from, to)
	if err != nil {
		log.Printf("Warning: error when get transfer logs: %s\n", err.Error())
		retry--
		return cli.GetTransferLogsRetry(from, to, retry)
	}
	return m
}

//GetTransferLogs get transfer events
func (cli *Cli) GetTransferLogs(from, to int) ([]*TransferLog, error) {
	transferLogs := make([]*TransferLog, 0)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
		Addresses: []common.Address{
			cli.ytaContractAddr,
		},
	}
	logs, err := cli.client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}
	contractAbi, err := abi.JSON(strings.NewReader(string(ytc.YtcABI)))
	if err != nil {
		return nil, err
	}

	transferSig := []byte("Transfer(address,address,uint256)")
	transferSigHash := crypto.Keccak256Hash(transferSig)
	for _, vLog := range logs {
		switch strings.ToLower(vLog.Topics[0].Hex()) {
		case strings.ToLower(transferSigHash.Hex()):
			var transferEvent TransferLog
			err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				return nil, err
			}
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
			transferLogs = append(transferLogs, &transferEvent)
		}
	}
	return transferLogs, nil
}
