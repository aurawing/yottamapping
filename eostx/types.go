package eostx

import (
	"github.com/eoscanada/eos-go"
)

//PayDepPool creates account
type PayDepPool struct {
	User  eos.AccountName `json:"user"`
	Quant eos.Asset       `json:"quant"`
}

//LockTransfer block rule
type LockTransfer struct {
	LockRuleID int64           `json:"lockruleid"`
	From       eos.AccountName `json:"from"`
	To         eos.AccountName `json:"to"`
	Quantity   eos.Asset       `json:"quantity"`
	Amount     eos.Asset       `json:"amount"`
	Memo       string          `json:"memo"`
}

//FrozenUser add fronzen rule to user
type FrozenUser struct {
	AccountName eos.AccountName `json:"account_name"`
	Time        int64           `json:"time"`
}
