package cmd

//Config configurations of yottamapping
type Config struct {
	FromIP          string `json:"fromIP"`
	FromPort        uint64 `json:"fromPort"`
	FromUsername    string `json:"fromUsername"`
	FromPassword    string `json:"fromPassword"`
	FromDbname      string `json:"fromDbname"`
	ToIP            string `json:"toIP"`
	ToPort          uint64 `json:"toPort"`
	ToUsername      string `json:"toUsername"`
	ToPassword      string `json:"toPassword"`
	ToDbname        string `json:"toDbname"`
	EtherscanAPIURL string `json:"etherscanAPIURL"`
	EthURL          string `json:"ethURL"`
	YtaContractAddr string `json:"ytaContractAddr"`
	MapContractAddr string `json:"mapContractAddr"`
	EosURL          string `json:"eosURL"`
	AdminAccount    string `json:"adminAccount"`
	AdminPK         string `json:"adminPK"`
	LockAccount     string `json:"lockAccount"`
	LockPK          string `json:"lockPK"`
	OperatorAccount string `json:"operatorAccount"`
	OperatorPK      string `json:"operatorPK"`
	UserPK          string `json:"userPK"`
	CPUStake        uint64 `json:"cpuStake"`
	NetStake        uint64 `json:"netStake"`
	RAMStake        uint64 `json:"ramStake"`
	BalThreshold    int64  `json:"balThreshold"`
}
