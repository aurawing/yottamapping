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
}

//BkRange structure
type BkRange struct {
	ID   int
	From int
	To   int
}

//NewMapping create a new mapping structure
func NewMapping(transactionHash string, blockNumber int, ethAddress, balance, param string, isVote int, nodeAccount string, status int) *Mapping {
	return &Mapping{
		TransactionHash: transactionHash,
		BlockNumber:     blockNumber,
		EthAddress:      ethAddress,
		Balance:         balance,
		Param:           param,
		IsVote:          isVote,
		NodeAccount:     nodeAccount,
		Status:          status,
	}
}

//NewBkRange create a new BkRange structure
func NewBkRange(id, from, to int) *BkRange {
	return &BkRange{
		ID:   id,
		From: from,
		To:   to,
	}
}
