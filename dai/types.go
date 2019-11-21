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
	FronzenTime     uint64
}

//BkRange structure
type BkRange struct {
	ID  int
	End int
}

//NewMapping create a new mapping structure
func NewMapping(transactionHash string, blockNumber int, ethAddress, balance, param string, isVote int, nodeAccount string, status int, ytaAccount, blockRule string, fronzenTime uint64) *Mapping {
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
	}
}

//NewBkRange create a new BkRange structure
func NewBkRange(id, end int) *BkRange {
	return &BkRange{
		ID:  id,
		End: end,
	}
}
