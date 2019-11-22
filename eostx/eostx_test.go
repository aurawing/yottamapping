package eostx

import (
	"os"
	"testing"
)

var etx *EosTX

func TestMain(m *testing.M) {
	etx = NewInstance("http://47.112.119.28:8888", "yottacreator", "5Kkvt7wy8S3PNJ7LeZxKtFh7ZmNbPSqGmeuLFcbttJ1GhTcZ1Pi", "lockadminusr", "5JcDH48njDbUQLu1R8SWwKsfWLnqBpWXDDiCgxFC3hioDuwLhVx", "yottaairdrop", "5K9dhuEfJeMxAB6x8fMJUuw3ntAJoDcVz59SaZVTrKaTmbLAjyh", "5KBDZghkb7WQ7jfQoreMxsWCcZi5WYg5HjDfMqHmgNG5pBX1WT9", 40000, 10000)
	os.Exit(m.Run())
}

func TestMappingTx(t *testing.T) {
	m := make(map[int]int64)
	m[100002] = 100000000
	//txid, err := etx.CreateAccountTx("yottachain32", "EOS6JidcEehqdQ3Gtq1c3EkRroY6Usc4NckqEeiusCvgovZgLfodY", 1000000, 1576911600, true /*"producer1",*/, m)
	txid, err := etx.VoteTx("yottachain32", "EOS6JidcEehqdQ3Gtq1c3EkRroY6Usc4NckqEeiusCvgovZgLfodY", 1000000, 1576911600, "producer1", m)
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("do mapping successful: %s\n", txid)
	}
}
