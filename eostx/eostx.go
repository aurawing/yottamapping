package eostx

import (
	"fmt"
	"log"
	"strings"

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
	lockAccount     string
	lockPK          string
	operatorAccount string
	operatorPK      string
	userPK          string
	cpuStake        int64
	netStake        int64
}

// New create a new eostx instance contans connect url, contract owner and it's private key
func New(url, adminAccount, adminPK, lockAccount, lockPK, operatorAccount, operatorPK, userPK string, cpuStake, netStake int64) *EosTX {
	api := eos.New(url)
	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(adminPK)
	err = keyBag.ImportPrivateKey(lockPK)
	err = keyBag.ImportPrivateKey(operatorPK)
	err = keyBag.ImportPrivateKey(userPK)
	if err != nil {
		log.Fatalf("error when import private key into keybag: %s\n", err.Error())
	}
	api.SetSigner(keyBag)
	return &EosTX{API: api, adminAccount: adminAccount, adminPK: adminPK, lockAccount: lockAccount, lockPK: lockPK, operatorAccount: operatorAccount, operatorPK: operatorPK, userPK: userPK, cpuStake: cpuStake, netStake: netStake}
}

// //GetAvailableResources get value of all available resources
// func (eostx *EosTX) GetAvailableResources(account string) (int64, int64, int64) {
// 	info, err := eostx.API.GetAccount(eos.AN(account))
// 	if err != nil {
// 		panic(fmt.Sprintf("get available resources error: %s\n", err.Error()))
// 	}
// 	return int64(info.CPULimit.Available), int64(info.NetLimit.Available), int64(info.RAMQuota - info.RAMUsage)
// }

//IfTransactioIrreversible check if transaction is irreversiblen
func (eostx *EosTX) IfTransactioIrreversible(txid string) bool {
	resp, err := eostx.API.GetTransaction(txid)
	if err != nil {
		log.Fatalf("error happens when get transaction information: %s\n", err.Error())
	}
	if resp.Receipt.Status != eos.TransactionStatusExecuted {
		return false
	}
	blockResp, err := eostx.API.GetBlockByNum(resp.BlockNum)
	if err != nil {
		log.Fatalf("error happens when get block by number from transaction information: %s\n", err.Error())
	}
	for _, tx := range blockResp.Transactions {
		if tx.Transaction.ID.String() == txid {
			return resp.LastIrreversibleBlock > resp.BlockNum
		}
	}
	return false
}

//CreateAccountTx create a new account and add block rules
func (eostx *EosTX) CreateAccountTx(userAccount, userPubkey string, balance int64, frozenTime int64, needVote bool, blockMap map[int]int64) (string, error) {
	if strings.HasPrefix(userPubkey, "YTA") {
		userPubkey = fmt.Sprintf("%s%s", "EOS", strings.TrimLeft(userPubkey, "YTA"))
	}
	eostx.API.SetCustomGetRequiredKeys(func(tx *eos.Transaction) ([]ecc.PublicKey, error) {
		publickey1, err := ytcrypto.GetPublicKeyByPrivateKey(eostx.adminPK)
		pubkey1, err := ecc.NewPublicKey(fmt.Sprintf("%s%s", "EOS", publickey1))
		publickey2, err := ytcrypto.GetPublicKeyByPrivateKey(eostx.lockPK)
		pubkey2, err := ecc.NewPublicKey(fmt.Sprintf("%s%s", "EOS", publickey2))
		publickey3, err := ytcrypto.GetPublicKeyByPrivateKey(eostx.operatorPK)
		pubkey3, err := ecc.NewPublicKey(fmt.Sprintf("%s%s", "EOS", publickey3))
		if err != nil {
			log.Fatalf("error when get required public keys: %s\n", err.Error())
		}
		if frozenTime > 0 {
			return []ecc.PublicKey{pubkey1, pubkey2, pubkey3}, nil
		}
		return []ecc.PublicKey{pubkey1, pubkey3}, nil
	})
	var pubkeyu ecc.PublicKey
	if needVote {
		pk, _ := ytcrypto.GetPublicKeyByPrivateKey(eostx.userPK)
		pubkeyu, _ = ecc.NewPublicKey(fmt.Sprintf("%s%s", "EOS", pk))
	} else {
		pubkeyu, _ = ecc.NewPublicKey(userPubkey)
	}

	//1. create account
	action11 := system.NewNewAccount(eos.AN(eostx.adminAccount), eos.AN(userAccount), pubkeyu)
	action12 := system.NewBuyRAMBytes(eos.AN(eostx.adminAccount), eos.AN(userAccount), 4096)
	action13 := system.NewDelegateBW(eos.AN(eostx.operatorAccount), eos.AN(userAccount), NewYTAAsset(eostx.cpuStake), NewYTAAsset(eostx.netStake), true)
	actions := []*eos.Action{action11, action12, action13}
	//2. block
	action2s := make([]*eos.Action, 0)
	if frozenTime > 0 {
		action21 := &eos.Action{
			Account: eos.AN("hddlock12345"),
			Name:    eos.ActN("frozenuser"),
			Authorization: []eos.PermissionLevel{
				{Actor: eos.AN(eostx.lockAccount), Permission: eos.PN("active")},
			},
			ActionData: eos.NewActionData(FrozenUser{AccountName: eos.AN(userAccount), Time: frozenTime}),
		}
		action2s = append(action2s, action21)
	}
	for id, bal := range blockMap {
		blockAction := &eos.Action{
			Account: eos.AN("hddlock12345"),
			Name:    eos.ActN("locktransfer"),
			Authorization: []eos.PermissionLevel{
				{Actor: eos.AN(eostx.operatorAccount), Permission: eos.PN("active")},
			},
			ActionData: eos.NewActionData(LockTransfer{LockRuleID: int64(id), From: eos.AN(eostx.operatorAccount), To: eos.AN(userAccount), Quantity: NewYTAAsset(bal), Amount: NewYTAAsset(0), Memo: "add lock record"}),
		}
		action2s = append(action2s, blockAction)
	}
	if len(action2s) > 0 {
		actions = append(actions, action2s...)
	}
	//3. airdrop
	action31 := token.NewTransfer(eos.AN(eostx.operatorAccount), eos.AN(userAccount), NewYTAAsset(balance-eostx.cpuStake-eostx.netStake), "mapping transfer")
	actions = append(actions, action31)

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

//VoteTx do voting
func (eostx *EosTX) VoteTx(userAccount, userPubkey string, balance int64, frozenTime int64, voteAccount string) (string, error) {
	if strings.HasPrefix(userPubkey, "YTA") {
		userPubkey = fmt.Sprintf("%s%s", "EOS", strings.TrimLeft(userPubkey, "YTA"))
	}
	eostx.API.SetCustomGetRequiredKeys(func(tx *eos.Transaction) ([]ecc.PublicKey, error) {
		publickey1, err := ytcrypto.GetPublicKeyByPrivateKey(eostx.lockPK)
		pubkey1, err := ecc.NewPublicKey(fmt.Sprintf("%s%s", "EOS", publickey1))
		publickey2, err := ytcrypto.GetPublicKeyByPrivateKey(eostx.userPK)
		pubkey2, err := ecc.NewPublicKey(fmt.Sprintf("%s%s", "EOS", publickey2))
		if err != nil {
			log.Fatalf("error when get required public keys: %s\n", err.Error())
		}
		if frozenTime > 0 {
			return []ecc.PublicKey{pubkey1, pubkey2}, nil
		}
		return []ecc.PublicKey{pubkey2}, nil
	})
	//use default user privatekey
	pk, _ := ytcrypto.GetPublicKeyByPrivateKey(eostx.userPK)
	pubkeyu, _ := ecc.NewPublicKey(fmt.Sprintf("%s%s", "EOS", pk))
	actions := make([]*eos.Action, 0)
	//1. unfrozen
	if frozenTime > 0 {
		action1 := &eos.Action{
			Account: eos.AN("hddlock12345"),
			Name:    eos.ActN("frozenuser"),
			Authorization: []eos.PermissionLevel{
				{Actor: eos.AN(eostx.lockAccount), Permission: eos.PN("active")},
			},
			ActionData: eos.NewActionData(FrozenUser{AccountName: eos.AN(userAccount), Time: 0}),
		}
		actions = append(actions, action1)
	}
	//2. delegate storage
	action2 := &eos.Action{
		Account: eos.AN("hdddeposit12"),
		Name:    eos.ActN("paydeppool"),
		Authorization: []eos.PermissionLevel{
			{Actor: eos.AN(userAccount), Permission: eos.PN("active")},
		},
		ActionData: eos.NewActionData(PayDepPool{User: eos.AN(userAccount), Quant: NewYTAAsset(balance - eostx.cpuStake - eostx.netStake)}),
	}
	actions = append(actions, action2)
	//3. vote
	action3 := system.NewVoteProducer(eos.AN(userAccount), eos.AN(""), eos.AN(voteAccount))
	actions = append(actions, action3)
	//4. frozen
	if frozenTime > 0 {
		action4 := &eos.Action{
			Account: eos.AN("hddlock12345"),
			Name:    eos.ActN("frozenuser"),
			Authorization: []eos.PermissionLevel{
				{Actor: eos.AN(eostx.lockAccount), Permission: eos.PN("active")},
			},
			ActionData: eos.NewActionData(FrozenUser{AccountName: eos.AN(userAccount), Time: frozenTime}),
		}
		actions = append(actions, action4)
	}
	//5. change public key
	pubkeyu, err := ecc.NewPublicKey(userPubkey)
	if err != nil {
		log.Fatalf(err.Error())
	}
	au := eos.Authority{
		Threshold: 1,
		Keys: []eos.KeyWeight{
			eos.KeyWeight{
				PublicKey: pubkeyu,
				Weight:    1,
			},
		},
	}
	action5 := system.NewUpdateAuth(eos.AN(userAccount), eos.PN("active"), eos.PN("owner"), au, eos.PN("owner"))
	action6 := system.NewUpdateAuth(eos.AN(userAccount), eos.PN("owner"), eos.PN(""), au, eos.PN("owner"))
	actions = append(actions, action5, action6)

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
