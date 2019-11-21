package dai

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aurawing/yottamapping/etherscan"
	_ "github.com/go-sql-driver/mysql" // driver auto register
)

//Dai data access interface
type Dai struct {
	db           *sql.DB
	etherscanCli *etherscan.Cli
}

//New create new Dao instance
func New(ip string, port int, username, password, dbname, apiURL, contractAddr string) *Dai {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, ip, port, dbname))
	checkErr(err)
	return &Dai{db: db, etherscanCli: etherscan.NewCli(apiURL, contractAddr)}
}

//Close close driver
func (dai *Dai) Close() error {
	return dai.db.Close()
}

//FetchNewData fetch new data from upstream database
func (dai *Dai) FetchNewData(from, to int) []*Mapping {
	mappings := make([]*Mapping, 0)
	rows, err := dai.db.Query("select transactionHash,blockNumber,ethAddress,balance,param,isVote,nodeAccount from mapping where blockBumber between ? and ?", from, to)
	checkErr(err)
	for rows.Next() {
		var transactionHash string
		var blockNumber int
		var ethAddress string
		var balance string
		var param string
		var isVote int
		var nodeAccount string
		err := rows.Scan(&transactionHash, &blockNumber, &ethAddress, &balance, &param, &isVote, &nodeAccount)
		checkErr(err)
		ytaAccount, err := EthAddrToName(ethAddress)
		checkErr(err)
		mappings = append(mappings, NewMapping(transactionHash, blockNumber, ethAddress, balance, param, isVote, nodeAccount, 0, ytaAccount, ""))
	}
	return mappings
}

//GetBkRange get range of blocks
func (dai *Dai) GetBkRange() *BkRange {
	var id, end int
	err := dai.db.QueryRow("select `id`,`end` from bkrange where id=?", 1).Scan(&id, &end)
	checkErr(err)
	return NewBkRange(id, end)
}

//UpdateLocalData update local database of mapping
func (dai *Dai) UpdateLocalData(mappings []*Mapping, v Verifier, from, to int) error {
	tx, err := dai.db.Begin()
	checkErr(err)
	stmt, err := tx.Prepare("insert into mapping (transactionHash, blockNumber, ethAddress, balance, param, isVote, nodeAccount, status, ytaAccount, blockRule) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		checkErr(err)
	}
	for _, m := range mappings {
		if v(m) {
			m.Status = 1
		}
		ruleMap := dai.etherscanCli.MappingRuleID(m.EthAddress, m.BlockNumber)
		ruleMapJSON, err := json.Marshal(ruleMap)
		if err != nil {
			log.Fatalf("error when marshal rule map to json: %s\n", err.Error())
		}
		_, err = stmt.Exec(m.TransactionHash, m.BlockNumber, m.EthAddress, m.Balance, m.Param, m.IsVote, m.NodeAccount, m.Status, m.YtaAccount, string(ruleMapJSON))
		if err != nil {
			tx.Rollback()
			checkErr(err)
		}
	}
	stmt, err = tx.Prepare("update bkrange set end=? where id=?")
	if err != nil {
		tx.Rollback()
		checkErr(err)
	}
	res, err := stmt.Exec(to, 1)
	if err != nil {
		tx.Rollback()
		checkErr(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		tx.Rollback()
		checkErr(err)
	}
	if affect != 1 {
		tx.Rollback()
		log.Fatalf("update bkrange failed, affect rows: %d\n", affect)
	}
	return tx.Commit()
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error when operate database: %s\n", err.Error())
	}
}
