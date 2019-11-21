package yottamapping

import (
	"log"

	dai "github.com/aurawing/yottamapping/dai"
	ethc "github.com/aurawing/yottamapping/eth"
)

//Mapper structure
type Mapper struct {
	ethcli *ethc.Cli
	from   *dai.Dai
	to     *dai.Dai
}

//NewMapper create a new mapper structure
func NewMapper(apiURL, ethURL, contractAddr, fromIP string, fromPort int, fromUsername, fromPassword, fromDbname, toIP string, toPort int, toUsername, toPassword, toDbname string) *Mapper {
	return &Mapper{
		ethcli: ethc.NewClient(ethURL, contractAddr),
		from:   dai.New(fromIP, fromPort, fromUsername, fromPassword, fromDbname, apiURL, contractAddr),
		to:     dai.New(toIP, toPort, toUsername, toPassword, toDbname, apiURL, contractAddr),
	}
}

//PullData pull data from upstream database and update local database
func (mapper *Mapper) PullData() {
	fromBkRange := mapper.from.GetBkRange()
	toBkRange := mapper.to.GetBkRange()
	if fromBkRange.End <= toBkRange.End {
		log.Fatalf("no new data for fetching, latest block number: %d\n", toBkRange.End)
	}
	log.Printf("fetching data of source database from block %d to %d...\n", toBkRange.End+1, fromBkRange.End)
	dbMappings := mapper.from.FetchNewData(toBkRange.End+1, fromBkRange.End)
	log.Printf("fetched %d records in source database\n", len(dbMappings))
	log.Printf("fetching events on ethereum network...\n")
	bkMappings, err := mapper.ethcli.GetFreezedLogs(toBkRange.End+1, fromBkRange.End)
	if err != nil {
		log.Fatalf("get freezed logs failed: %s\n", err.Error())
	}
	log.Printf("fetched %d events on ethereum network\n", len(bkMappings))
	log.Printf("verifying and updating local database...\n")
	err = mapper.to.UpdateLocalData(dbMappings, mapper.Verify(bkMappings), toBkRange.End+1, fromBkRange.End)
	if err != nil {
		log.Fatalf("please check local database, error happens when commit: %s\n", err.Error())
	}
}
