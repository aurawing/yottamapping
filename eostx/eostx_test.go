package eostx

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/eoscanada/eos-go"
)

var etx *EosTX

func TestMain(m *testing.M) {
	etx = New("https://wallet.yottachain.net", "yottacreator", "5Kkvt7wy8S3PNJ7LeZxKtFh7ZmNbPSqGmeuLFcbttJ1GhTcZ1Pi", "lockadminusr", "5JcDH48njDbUQLu1R8SWwKsfWLnqBpWXDDiCgxFC3hioDuwLhVx", "yottaairdrop", "5K9dhuEfJeMxAB6x8fMJUuw3ntAJoDcVz59SaZVTrKaTmbLAjyh", "5KBDZghkb7WQ7jfQoreMxsWCcZi5WYg5HjDfMqHmgNG5pBX1WT9", 40000, 10000)
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
	b := etx.IfTransactioIrreversible("60b1e47b37548c9100dcaf8e6f54e55d0345acf11a707346e7e92eb37d85bf88")
	t.Logf("transaction is irreversible: %t\n", b)
}

func TestGetAccount(t *testing.T) {
	account := eos.AccountName("yottaairdrop")
	info, err := etx.API.GetAccount(account)
	if err != nil {
		if err == eos.ErrNotFound {
			fmt.Printf("unknown account: %s", account)
			return
		}

		panic(fmt.Errorf("get account: %s", err))
	}

	bytes, err := json.Marshal(info)
	if err != nil {
		panic(fmt.Errorf("json marshal response: %s", err))
	}

	fmt.Println(string(bytes))
}
