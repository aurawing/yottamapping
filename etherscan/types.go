package etherscan

//RuleMap mapping of ERC20 lock rule address and EOS lock rule ID
var RuleMap map[string]int

func init() {
	RuleMap = make(map[string]int)
	RuleMap["0x7702b3113e2851d7e5312d67431faecd128807e7"] = 100001
	RuleMap["0x0e11d9b9b159002e4c4d657d42ae3ad60f519cf6"] = 100002
	RuleMap["0xe7636a1aad84d194a95db2b0c70507f6c9b37789"] = 100003
	RuleMap["0x0ece6f913eb73c063f28bda8cc540ab419833b63"] = 100004
	RuleMap["0x10ba852d94c0fb8cfc189ed6f13d2de08c8ce2ab"] = 100005
	RuleMap["0x3421dcf836287b4cc48a3620512b871ea603289a"] = 100005
	RuleMap["0xf30b7276e79f92f710a1447e573fd5a4f2c933af"] = 100005
	RuleMap["0x65506ce6f06a588ec692593a0a0069dffe459055"] = 100005
	RuleMap["0x8ec2667b3fdb5945f1a1f712512f2959ce6054fe"] = 100005
	RuleMap["0xe76cec63e051da14666308cb2e1860a0ca481eb7"] = 100005
	RuleMap["0x3c3f6cf57d8b87243209b3136a9ce34a49b9f912"] = 100005
	RuleMap["0x4df97a71e72e250a75ddcd24d914e7f2fb595d93"] = 100006
	RuleMap["0x8981d66c9d72a06f300e45f6c482be3b001587fe"] = 100007

	// RuleMap["0xaa3b2078fb5d73c71d52810fcb89b6d6c6ab87c7"] = 100010
	// RuleMap["0x63c00b8187afcfff857fedbb7834893c13605b1e"] = 100011
	// RuleMap["0x23a1f5f0e7d596da1b7560852331988487c8721f"] = 100012
	// RuleMap["0xd50567e1c0a7b1232edbfbb109ab9109d4d26422"] = 100013
	// RuleMap["0xb85fc4a6a95d299e2c2a67e81c5af297703a7e42"] = 100014
	// RuleMap["0xff24eeb602d24a3542267cdde50919261c24d45c"] = 100015
	// RuleMap["0x75b22d0afd341854e671b58ebea5b5b009858fe1"] = 100016
}

//Response response by etherscan API
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Result  []*Tx  `json:"result"`
}

//Tx transaction infomation returned by etherscan API
type Tx struct {
	BlockNumber       string `json:"blockNumber"`
	TimeStamp         string `json:"timeStamp"`
	Hash              string `json:"hash"`
	Nonce             string `json:"nonce"`
	BlockHash         string `json:"blockHash"`
	From              string `json:"from"`
	ContractAddress   string `json:"contractAddress"`
	To                string `json:"to"`
	Value             string `json:"value"`
	TokenName         string `json:"tokenName"`
	TokenSymbol       string `json:"tokenSymbol"`
	TokenDecimal      string `json:"tokenDecimal"`
	TransactionIndex  string `json:"transactionIndex"`
	Gas               string `json:"gas"`
	GasPrice          string `json:"gasPrice"`
	GasUsed           string `json:"gasUsed"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`
	Input             string `json:"input"`
	Confirmations     string `json:"confirmations"`
}
