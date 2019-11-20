package dai

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // driver auto register
)

//Dai data access interface
type Dai struct {
	db *sql.DB
}

//New create new Dao instance
func New(ip string, port int, username, password, dbname string) *Dai {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s:%d/%s", username, password, ip, port, dbname))
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
		mappings = append(mappings, NewMapping(transactionHash, blockNumber, ethAddress, balance, param, isVote, nodeAccount, 0))
	}
	return mappings
}

//GetBkRange get range of blocks
func (dai *Dai) GetBkRange() *BkRange {
	var id, from, to int
	err := dai.db.QueryRow("select id,from,to from bkrange where id=?", 1).Scan(&id, &from, &to)
	checkErr(err)
	return NewBkRange(id, from, to)
}

//UpdateLocalData update local database of mapping
func (dai *Dai) UpdateLocalData(mappings []*Mapping, v Verifier, from, to int) error {
	tx, err := dai.db.Begin()
	checkErr(err)
	stmt, err := tx.Prepare("insert into mapping (transactionHash, blockNumber, ethAddress, balance, param, isVote, nodeAccount, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		tx.Rollback()
		checkErr(err)
	}
	for _, m := range mappings {
		if v(m) {
			m.Status = 1
		}
		_, err = stmt.Exec(m.TransactionHash, m.BlockNumber, m.EthAddress, m.Balance, m.Param, m.IsVote, m.NodeAccount, m.Status)
		if err != nil {
			tx.Rollback()
			checkErr(err)
		}
	}
	stmt, err = tx.Prepare("update bkrange set from=?, to=? where id=?")
	if err != nil {
		tx.Rollback()
		checkErr(err)
	}
	res, err := stmt.Exec(from, to, 1)
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