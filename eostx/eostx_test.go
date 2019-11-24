package eostx

import (
	"os"
	"testing"
)

var etx *EosTX

func TestMain(m *testing.M) {
	etx = New("http://47.112.119.28:8888", "yottacreator", "5Kkvt7wy8S3PNJ7LeZxKtFh7ZmNbPSqGmeuLFcbttJ1GhTcZ1Pi", "lockadminusr", "5JcDH48njDbUQLu1R8SWwKsfWLnqBpWXDDiCgxFC3hioDuwLhVx", "yottaairdrop", "5K9dhuEfJeMxAB6x8fMJUuw3ntAJoDcVz59SaZVTrKaTmbLAjyh", "5KBDZghkb7WQ7jfQoreMxsWCcZi5WYg5HjDfMqHmgNG5pBX1WT9", 40000, 10000)
	os.Exit(m.Run())
}

func TestMappingTx(t *testing.T) {
	m := make(map[int]int64)
	m[100002] = 100000000
	txid, err := etx.CreateAccountTx("yottachain34", "EOS6JidcEehqdQ3Gtq1c3EkRroY6Usc4NckqEeiusCvgovZgLfodY", 1000000, 0, false /*"producer1",*/, nil)
	//txid, err := etx.VoteTx("yottachain34", "EOS6JidcEehqdQ3Gtq1c3EkRroY6Usc4NckqEeiusCvgovZgLfodY", 1000000, 1576911600, "producer1")
	if err != nil {
		t.Error(err.Error())
	} else {
		t.Logf("do mapping successful: %s\n", txid)
	}
}

func TestIfTransactioIrreversible(t *testing.T) {
	b := etx.IfTransactioIrreversible("d8815a3c582bc38469885a82fcd7a916f55ccef216d26234481b06678a409a42")
	t.Logf("transaction is irreversible: %t\n", b)
}
