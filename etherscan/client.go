package etherscan

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
)

//Cli client for etherscan
type Cli struct {
	client       *http.Client
	apiURL       string //etherscan API URL
	contractAddr string //ERC20 token contract address
}

//NewCli create a new Cli structure with specified API URL and ERC20 token contract address
func NewCli(apiURL, contractAddr string) *Cli {
	return &Cli{client: &http.Client{}, apiURL: apiURL, contractAddr: contractAddr}
}

//GetTxList fetch all transactions of certain address
func (cli *Cli) GetTxList(addr string, blockNumber int) []*Tx {
	url := fmt.Sprintf("%s?module=account&action=tokentx&contractaddress=%s&address=%s&startblock=6425985&endblock=%d&sort=asc&apikey=X53VEQC87MSYAS9XY436V8QZW7MPM8UEIG", cli.apiURL, cli.contractAddr, addr, blockNumber)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("error when create new request to fetch tx list: %s\n", err.Error())
	}
	response, err := cli.client.Do(request)
	if err != nil {
		log.Fatalf("error when send request for fetching tx list: %s\n", err.Error())
	}
	if response.StatusCode != 200 {
		log.Fatalf("response code is not 200 when send request for fetching tx list: %d\n", response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("error when read response of tx list: %s\n", err.Error())
	}
	var resp = Response{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		log.Fatalf("error when unmarshal response of tx list: %s\n", err.Error())
	}
	if resp.Status != "1" {
		log.Fatalf("response status is not 1: %s %s\n", resp.Status, resp.Message)
	}
	return resp.Result
}

//MappingRuleID transform all rule address of one address to rule ID
func (cli *Cli) MappingRuleID(addr string, blockNumber int) map[int]uint64 {
	addr = strings.ToLower(addr)
	result := make(map[int]uint64)
	transactions := cli.GetTxList(addr, blockNumber)
	for _, tx := range transactions {
		if tx.To == addr && RuleMap[tx.From] > 0 {
			var v string
			if len(tx.Value) <= 14 {
				v = "0"
			} else {
				v = tx.Value[0 : len(tx.Value)-14]
			}
			amount, err := strconv.ParseUint(v, 10, 64)
			if err != nil {
				log.Fatalf("error when convert block amount string to uint64: %s\n", err.Error())
			}
			result[RuleMap[tx.From]] += amount
		}
	}
	return result
}
