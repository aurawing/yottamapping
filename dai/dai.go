package dai

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // driver auto register
)

//Dai data access interface
type Dai struct {
	db *sql.DB
}

//New create new Dao instance
func New(ip string, port int, username, password, dbname string) *Dai {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", username, password, ip, port, dbname))
	checkErr(err)
	return &Dai{db: db}
}

//Close close driver
func (dai *Dai) Close() error {
	return dai.db.Close()
}

//FetchNewData fetch new data from upstream database
func (dai *Dai) FetchNewData(from, to int) []*Mapping {
	mappings := make([]*Mapping, 0)
	rows, err := dai.db.Query("select transactionHash, blockNumber, ethAddress, balance, param, isVote, nodeAccount from mapping where blockNumber between ? and ?", from, to)
	checkErr(err)
	i := 0
	for rows.Next() {
		i++
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
		mappings = append(mappings, NewMapping(transactionHash, blockNumber, ethAddress, balance, param, isVote, nodeAccount, 0, ytaAccount, "", 0, "", "", time.Now().Unix()))
		log.Printf("fetched mapping record %d from upstream database: %s %s %s\n", i, transactionHash, ethAddress, ytaAccount)
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
	stmt, err := tx.Prepare("insert into mapping (transactionHash, blockNumber, ethAddress, balance, param, isVote, nodeAccount, status, ytaAccount, blockRule, frozenTime, modifyTime) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		checkErr(err)
	}
	i := 0
	for _, m := range mappings {
		if v(m) {
			m.Status = 1
		} else {
			i++
			log.Printf("!!! verify data failed: %s\n", m.TransactionHash)
		}
		_, err = stmt.Exec(m.TransactionHash, m.BlockNumber, m.EthAddress, m.Balance, m.Param, m.IsVote, m.NodeAccount, m.Status, m.YtaAccount, m.BlockRule, m.FrozenTime, time.Now().Unix())
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
		log.Fatalf("!!! update bkrange failed, affect rows: %d\n", affect)
	}
	if i > 0 {
		log.Printf("!!! %d mapping records failed to verify\n", i)
	} else {
		log.Printf("all incremental data fetched successfully!")
	}
	return tx.Commit()
}

//BrowserAndModify browser mapping datas and modify by function Modifier
func (dai *Dai) BrowserAndModify(m Modifier, status int) {
	rows, err := dai.db.Query("select transactionHash, blockNumber, ethAddress, balance, param, status, isVote, nodeAccount, ytaAccount, blockRule, frozenTime, txid1, txid2, modifyTime from mapping where status=?", status)
	checkErr(err)
	i := 0
	for rows.Next() {
		i++
		var transactionHash string
		var blockNumber int
		var ethAddress string
		var balance string
		var param string
		var status int
		var isVote int
		var nodeAccount string
		var ytaAccount string
		var blockRule string
		var frozenTime int64
		var txid1 string
		var txid2 string
		var modifyTime int64
		err := rows.Scan(&transactionHash, &blockNumber, &ethAddress, &balance, &param, &status, &isVote, &nodeAccount, &ytaAccount, &blockRule, &frozenTime, &txid1, &txid2, &modifyTime)
		checkErr(err)
		mapping := NewMapping(transactionHash, blockNumber, ethAddress, balance, param, isVote, nodeAccount, status, ytaAccount, blockRule, frozenTime, txid1, txid2, modifyTime)
		log.Printf("fetched mapping record %d from local database: %s %s %s\n", i, transactionHash, ethAddress, ytaAccount)
		m(mapping)
	}
}

//UpdateTxid1 update txid1 field after create new YTA account
func (dai *Dai) UpdateTxid1(txHash, txid1 string, status int) {
	stmt, err := dai.db.Prepare("update mapping set txid1=?, status=?, modifyTime=? where transactionHash=?")
	if err != nil {
		checkErr(err)
	}
	res, err := stmt.Exec(txid1, status, time.Now().Unix(), txHash)
	if err != nil {
		checkErr(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		checkErr(err)
	}
	if affect != 1 {
		log.Fatalf("update txid1 failed,transactionHash: %s, affect rows: %d\n", txHash, affect)
	}
}

//UpdateTxid2 update txid2 field after voting
func (dai *Dai) UpdateTxid2(txHash, txid2 string, status int) {
	stmt, err := dai.db.Prepare("update mapping set txid2=?, status=?, modifyTime=? where transactionHash=?")
	if err != nil {
		checkErr(err)
	}
	res, err := stmt.Exec(txid2, status, time.Now().Unix(), txHash)
	if err != nil {
		checkErr(err)
	}
	affect, err := res.RowsAffected()
	if err != nil {
		checkErr(err)
	}
	if affect != 1 {
		log.Fatalf("update txid2 failed,transactionHash: %s, affect rows: %d\n", txHash, affect)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatalf("error when operate database: %s\n", err.Error())
	}
}
