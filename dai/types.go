package dai

//Verifier verify mapping data
type Verifier func(*Mapping) bool

//Mapping structure
type Mapping struct {
	TransactionHash string
	BlockNumber     int
	EthAddress      string
	Balance         string
	Param           string
	IsVote          int
	NodeAccount     string
	Status          int
	YtaAccount      string
	BlockRule       string
	FronzenTime     int64
	TxID1           string
	TxID2           string
	ModifyTime      int64
}

//BkRange structure
type BkRange struct {
	ID  int
	End int
}

//NewMapping create a new mapping structure
func NewMapping(transactionHash string, blockNumber int, ethAddress, balance, param string, isVote int, nodeAccount string, status int, ytaAccount, blockRule string, fronzenTime int64, txid1, txid2 string, modifyTime int64) *Mapping {
	return &Mapping{
		TransactionHash: transactionHash,
		BlockNumber:     blockNumber,
		EthAddress:      ethAddress,
		Balance:         balance,
		Param:           param,
		IsVote:          isVote,
		NodeAccount:     nodeAccount,
		Status:          status,
		YtaAccount:      ytaAccount,
		BlockRule:       blockRule,
		FronzenTime:     fronzenTime,
		TxID1:           txid1,
		TxID2:           txid2,
		ModifyTime:      modifyTime,
	}
}

//NewBkRange create a new BkRange structure
func NewBkRange(id, end int) *BkRange {
	return &BkRange{
		ID:  id,
		End: end,
	}
}
