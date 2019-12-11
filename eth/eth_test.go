package eth

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"
)

var cli *Cli

func TestMain(m *testing.M) {
	cli = NewClient("https://mainnet.infura.io", "0x5edc1a266e8b2c5e8086d373725df0690af7e3ea", "0xBb99a5B15fC7e5E8273Cc3e380741F1899485Aea")
	os.Exit(m.Run())
}

func TestGetFreezedLogs(t *testing.T) {
	m, err := cli.GetFreezedLogs(9082718, 99999999999999)
	if err != nil {
		t.Error(err.Error())
	}
	l := len(m)
	// event := m["0x549f6b834b7e9bf2c0d7fc414406eb8e51b900de568fc909a277d8a0d58e2c6f"]
	t.Logf("get freezed log successful: %d\n", l)
}

func TestGetTransferLogs(t *testing.T) {
	logs, err := cli.GetTransferLogs(6849506, 6849606)
	if err != nil {
		t.Error(err.Error())
	}
	t.Logf("get %d transfer events\n", len(logs))
}

func TestLatestBlockNumber(t *testing.T) {
	header, err := cli.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744
}
