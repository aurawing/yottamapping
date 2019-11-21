package eostx

import (
	"fmt"

	"github.com/eoscanada/eos-go"
	"github.com/eoscanada/eos-go/ecc"
	_ "github.com/eoscanada/eos-go/system"
	_ "github.com/eoscanada/eos-go/token"
	ytcrypto "github.com/yottachain/YTCrypto"
)

//EoxTX eos operate structure
type EosTX struct {
	API            *eos.API
	adminAccount   string
	adminPK        string
	opreateAccount string
	opreatePK      string
}

// NewInstance create a new eostx instance contans connect url, contract owner and it's private key
func NewInstance(url, adminAccount, adminPK, opreateAccount, opreatePK string) (*EosTX, error) {
	api := eos.New(url)
	keyBag := &eos.KeyBag{}
	err := keyBag.ImportPrivateKey(adminPK)
	if err != nil {
		return nil, fmt.Errorf("import private key: %s", err)
	}
	api.SetSigner(keyBag)
	api.SetCustomGetRequiredKeys(func(tx *eos.Transaction) ([]ecc.PublicKey, error) {
		publickey, _ := ytcrypto.GetPublicKeyByPrivateKey(privateKey)
		pubkey, _ := ecc.NewPublicKey(fmt.Sprintf("%s%s", "YTA", publickey))
		return []ecc.PublicKey{pubkey}, nil
	})
	return &EosTX{API: api, BpAccount: bpAccount, ContractOwnerM: contractOwnerM, ContractOwnerD: contractOwnerD, ShadowAccount: shadowAccount}, nil
}
