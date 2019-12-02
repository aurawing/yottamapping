package yottamapping

import (
	"log"

	dai "github.com/aurawing/yottamapping/dai"
	ethc "github.com/aurawing/yottamapping/eth"
	"github.com/aurawing/yottamapping/etherscan"
)

//NewMapper1 create a new mapper structure
func NewMapper1(apiURL, ethURL, ytaContractAddr, mapContractAddr, fromIP string, fromPort int, fromUsername, fromPassword, fromDbname, toIP string, toPort int, toUsername, toPassword, toDbname string, txTimeGap uint64, proxyURL string) *Mapper {
	return &Mapper{
		from:         dai.New(fromIP, fromPort, fromUsername, fromPassword, fromDbname),
		to:           dai.New(toIP, toPort, toUsername, toPassword, toDbname),
		ethcli:       ethc.NewClient(ethURL, ytaContractAddr, mapContractAddr),
		etherscanCli: etherscan.NewCli(apiURL, ytaContractAddr, proxyURL),
		txTimeGap:    txTimeGap,
	}
}

//PullData pull data from upstream database and update local database
func (mapper *Mapper) PullData() {
	fromBkRange := mapper.from.GetBkRange()
	toBkRange := mapper.to.GetBkRange()
	if fromBkRange.End <= toBkRange.End {
		log.Fatalf("no new data for fetching, latest block number: %d\n", toBkRange.End)
	}
	log.Printf("fetching data from source database between block %d to %d...\n", toBkRange.End+1, fromBkRange.End)
	dbMappings := mapper.from.FetchNewData(toBkRange.End+1, fromBkRange.End)
	log.Printf("fetched %d records from source database\n", len(dbMappings))
	log.Printf("fetching frozen events from ethereum network...\n")
	bkMappings, err := mapper.ethcli.GetFreezedLogs(toBkRange.End+1, fromBkRange.End)
	if err != nil {
		log.Fatalf("fetching frozen logs failed: %s\n", err.Error())
	}
	log.Printf("fetched %d frozen events from ethereum network\n", len(bkMappings))
	log.Printf("verifying and updating local database...\n")
	err = mapper.to.UpdateLocalData(dbMappings, mapper.Verify(bkMappings), toBkRange.End+1, fromBkRange.End)
	if err != nil {
		log.Fatalf("please check local database, error happens when committing: %s\n", err.Error())
	}
}
