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
	cli = NewClient("https://ropsten.infura.io", "0x103b8c9e072193F2D0E3B77cEfb6d5DC9294d0d2", "0x56E5B0BDaea3c33a8EC723E2E038453E84C14bA7")
	os.Exit(m.Run())
}

func TestGetFreezedLogs(t *testing.T) {
	m, err := cli.GetFreezedLogs(6849506, 6849506)
	if err != nil {
		t.Error(err.Error())
	}
	event := m["0x549f6b834b7e9bf2c0d7fc414406eb8e51b900de568fc909a277d8a0d58e2c6f"]
	t.Logf("get freezed log successful: %s -> %s\n", event.EthAddress, event.Balance)
}

func TestLatestBlockNumber(t *testing.T) {
	header, err := cli.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744
}
