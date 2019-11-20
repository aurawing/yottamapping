package eth

import (
	"os"
	"testing"
)

var cli *Cli

func TestMain(m *testing.M) {
	cli = NewClient("https://ropsten.infura.io", "0x56E5B0BDaea3c33a8EC723E2E038453E84C14bA7")
	os.Exit(m.Run())
}

func TestGetFreezedLogs(t *testing.T) {
	m, err := cli.GetFreezedLogs(6812240, 6812240)
	if err != nil {
		t.Error(err.Error())
	}
	event := m["0x549f6b834b7e9bf2c0d7fc414406eb8e51b900de568fc909a277d8a0d58e2c6f"]
	t.Logf("get freezed log successful: %s -> %s\n", event.EthAddress, event.Balance)
}