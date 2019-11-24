package dai

import (
	"os"
	"testing"
)

var d *Dai

func TestMain(m *testing.M) {
	d = New("127.0.0.1", 3306, "root", "", "test", "http://api-cn.etherscan.com/api", "0x103b8c9e072193F2D0E3B77cEfb6d5DC9294d0d2")
	os.Exit(m.Run())
}

func TestGetBkRange(t *testing.T) {
	r := d.GetBkRange()
	t.Logf("Get block range: %d\n", r.End)
}

func TestBrowserAndModify(t *testing.T) {
	d.BrowserAndModify(func(m *Mapping) bool {
		t.Logf("get mapping data: %s\n", m.TransactionHash)
		return true
	})
}
