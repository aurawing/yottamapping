package yottamapping

import (
	"log"

	dai "github.com/aurawing/yottamapping/dai"
	eth "github.com/aurawing/yottamapping/eth"
)

//Mapper structure
type Mapper struct {
	ethcli *eth.Cli
	from   *dai.Dai
	to     *dai.Dai
}

//NewMapper create a new mapper structure
func NewMapper(ethURL, contractAddr, fromIP string, fromPort int, fromUsername, fromPassword, fromDbname, toIP string, toPort int, toUsername, toPassword, toDbname string) *Mapper {
	return &Mapper{
		ethcli: eth.NewClient(ethURL, contractAddr),
		from:   dai.New(fromIP, fromPort, fromUsername, fromPassword, fromDbname),
		to:     dai.New(toIP, toPort, toUsername, toPassword, toDbname),
	}
}

//PullData pull data from upstream database and update local database
func (mapper *Mapper) PullData() {
	fromBkRange := mapper.from.GetBkRange()
	toBkRange := mapper.to.GetBkRange()
	if fromBkRange.To <= toBkRange.To {
		log.Fatalf("no new data for fetching, latest block number: %d\n", toBkRange.To)
	}
	log.Printf("1. fetching events from block %d to %d...\n", toBkRange.To+1, fromBkRange.To)
	dbMappings := mapper.from.FetchNewData(toBkRange.To+1, fromBkRange.To)
	log.Printf("fetched %d events\n", len(mappings))
	log.Printf("fetching events on ethereum mainnet...\n")
	bkMappings := eth.GetFreezedLogs(toBkRange.To+1, fromBkRange.To)
	err := mapper.to.UpdateLocalData(mdbMappings, dai.Verifier(mapper.Verify(bkMappings)), toBkRange.To+1, fromBkRange.To)
	if err != nil {
		log.Fatalf("please check localdatabase, error happens when commit: %s\n", err.Error())
	}
}
