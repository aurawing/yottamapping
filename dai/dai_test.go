package dai

import (
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
