package eostx

import (
	"fmt"
	"log"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	"github.com/eoscanada/eos-go/system"
	"github.com/eoscanada/eos-go/token"
	ytcrypto "github.com/yottachain/YTCrypto"
)

//EosTX eos operate structure
type EosTX struct {
	API             *eos.API
	adminAccount    string
	adminPK         string
	operatorAccount string
	operatorPK      string
	userPK          string
	cpuStake        int64
	netStake        int64
}

// NewInstance create a new eostx instance contans connect url, contract owner and it's private key
func NewInstance(url, adminAccount, adminPK, operatorAccount, operatorPK, userPK string, cpuStake, netStake int64) *EosTX {
	api := eos.New(url)
	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(adminPK)
	err = keyBag.ImportPrivateKey(operatorPK)
	err = keyBag.ImportPrivateKey(userPK)
	if err != nil {
		log.Fatalf("error when import private key into keybag: %s\n", err.Error())
	}
	api.SetSigner(keyBag)
	api.SetCustomGetRequiredKeys(func(tx *eos.Transaction) ([]ecc.PublicKey, error) {
		publickey1, _ := ytcrypto.GetPublicKeyByPrivateKey(adminPK)
		pubkey1, _ := ecc.NewPublicKey(fmt.Sprintf("%s%s", "YTA", publickey1))
		publickey2, _ := ytcrypto.GetPublicKeyByPrivateKey(operatorPK)
		pubkey2, _ := ecc.NewPublicKey(fmt.Sprintf("%s%s", "YTA", publickey2))
		publickey3, _ := ytcrypto.GetPublicKeyByPrivateKey(userPK)
		pubkey3, _ := ecc.NewPublicKey(fmt.Sprintf("%s%s", "YTA", publickey3))
		return []ecc.PublicKey{pubkey1, pubkey2, pubkey3}, nil
	})
	return &EosTX{API: api, adminAccount: adminAccount, adminPK: adminPK, operatorAccount: operatorAccount, operatorPK: operatorPK, userPK: userPK, cpuStake: cpuStake, netStake: netStake}
}

//MappingTx do mapping
func (eostx *EosTX) MappingTx(userAccount, userPubkey string, balance int64, needVote bool, voteAccount string, blockMap map[int]int64) (string, error) {
	var pubkeyu ecc.PublicKey
	if needVote {
		pk, _ := ytcrypto.GetPublicKeyByPrivateKey(eostx.userPK)
		pubkeyu, _ = ecc.NewPublicKey(fmt.Sprintf("%s%s", "YTA", pk))
	} else {
		pubkeyu, _ = ecc.NewPublicKey(userPubkey)
	}

	//1. create account
	action11 := system.NewNewAccount(eos.AN(eostx.adminAccount), eos.AN(userAccount), pubkeyu)
	action12 := system.NewBuyRAMBytes(eos.AN(eostx.adminAccount), eos.AN(userAccount), 4096)
	action13 := system.NewDelegateBW(eos.AN(eostx.adminAccount), eos.AN(userAccount), NewYTAAsset(eostx.cpuStake), NewYTAAsset(eostx.netStake), true)
	actions := []*eos.Action{action11, action12, action13}
	//2. block
	action2s := make([]*eos.Action, 0)
	for id, bal := range blockMap {
		blockAction := &eos.Action{
			Account: eos.AN("hddlock12345"),
			Name:    eos.ActN("locktransfer"),
			Authorization: []eos.PermissionLevel{
				{Actor: eos.AN(eostx.operatorAccount), Permission: eos.PN("active")},
			},
			ActionData: eos.NewActionData(LockTransfer{LockRuleID: uint64(id), From: eos.AN(eostx.operatorAccount), To: eos.AN(userAccount), Quantity: NewYTAAsset(bal), Amount: NewYTAAsset(0), Memo: "add lock record"}),
		}
		action2s = append(action2s, blockAction)
	}
	if len(action2s) > 0 {
		actions = append(actions, action2s...)
	}
	//3. airdrop
	action31 := token.NewTransfer(eos.AN(eostx.operatorAccount), eos.AN(userAccount), NewYTAAsset(balance-eostx.cpuStake-eostx.netStake), "mapping transfer")
	actions = append(actions, action31)
	if needVote {
		//4. delegate storage
		action41 := &eos.Action{
			Account: eos.AN("hdddeposit12"),
			Name:    eos.ActN("paydeppool"),
			Authorization: []eos.PermissionLevel{
				{Actor: eos.AN(userAccount), Permission: eos.PN("active")},
			},
			ActionData: eos.NewActionData(PayDepPool{User: eos.AN(userAccount), Quant: NewYTAAsset(balance - eostx.cpuStake - eostx.netStake)}),
		}
		//5. vote
		action51 := system.NewVoteProducer(eos.AN(userAccount), eos.AN(""), eos.AN(voteAccount))
		//6. change public key
		pubkeyu, _ := ecc.NewPublicKey(userPubkey)
		au := eos.Authority{
			Threshold: 1,
			Keys: []eos.KeyWeight{
				eos.KeyWeight{
					PublicKey: pubkeyu,
					Weight:    1,
				},
			},
		}
		action61 := system.NewUpdateAuth(eos.AN(userAccount), eos.PN("active"), eos.PN("owner"), au, eos.PN("owner"))
		actions = append(actions, action41, action51, action61)
	}

	txOpts := &eos.TxOptions{}
	if err := txOpts.FillFromChain(eostx.API); err != nil {
		return "", fmt.Errorf("filling tx opts: %s", err)
	}

	tx := eos.NewTransaction(actions, txOpts)
	_, packedTx, err := eostx.API.SignTransaction(tx, txOpts.ChainID, eos.CompressionNone)
	if err != nil {
		return "", fmt.Errorf("sign transaction: %s", err)
	}

	resp, err := eostx.API.PushTransaction(packedTx)
	if err != nil {
		return "", fmt.Errorf("push transaction: %s", err)
	}
	return resp.TransactionID, nil
}
