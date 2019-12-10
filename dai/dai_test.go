package dai

import (
	"fmt"
	"os"
	"testing"
)

var d *Dai

func TestMain(m *testing.M) {
	d = New("127.0.0.1", 3306, "root", "", "test")
	os.Exit(m.Run())
}

func TestGetBkRange(t *testing.T) {
	r := d.GetBkRange()
	t.Logf("Get block range: %d\n", r.End)
}

func TestBrowserAndModify(t *testing.T) {
	d.BrowserAndModify(func(m *Mapping) {
		t.Logf("get mapping data: %s\n", m.TransactionHash)
		return
	}, 1)
}

func TestNameToString(t *testing.T) {
	n, _ := EthAddrToName("0x941Aa12a96012f6F1A3CC0050Fa6537974C105ea")
	fmt.Println(n)
}
