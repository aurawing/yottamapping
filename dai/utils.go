package dai

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"time"

	"github.com/eoscanada/eos-go"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

//EthAddrToName convert ethereum address to EOS account name
func EthAddrToName(address string) (string, error) {
	address = strings.ToLower(address)
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(address) {
		return "", errors.New("address is invalid")
	}
	address = string(address[26:42])
	if address[0] == 0x30 {
		address = string(append([]byte{0x61 + byte(rand.Intn(6))}, address[1:]...))
	}
	b, err := hex.DecodeString(address)
	if err != nil {
		return "", fmt.Errorf("decode hex string failed: %s", err.Error())
	}
	v := uint64(binary.LittleEndian.Uint64(b))
	name := eos.NameToString(v)
	if len(name) > 12 {
		name = name[1:13]
	}
	name = strings.Replace(name, ".", "x", -1)
	return name, nil
}
