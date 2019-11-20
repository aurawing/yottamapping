package dai

import (
	"os"
	"testing"
)

var d *Dai

func TestMain(m *testing.M) {
	d = New("127.0.0.1", 3306, "root", "admin", "yottamapping")
	os.Exit(m.Run())
}

func TestGetBkRange(t *testing.T) {
	r := d.GetBkRange()
	if r.From != 6900000 {
		t.Error("Generate public key by private key failed.")
	}
}
