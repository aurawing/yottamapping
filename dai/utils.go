package dai

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strings"

	"github.com/eoscanada/eos-go"
)

//EthAddrToName convert ethereum address to EOS account name
func EthAddrToName(address string) (string, error) {
	address = strings.ToLower(address)
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	if !re.MatchString(address) {
		return "", errors.New("address is invalid")
	}
	b, err := hex.DecodeString(string(([]byte(address))[26:42]))
	if err != nil {
		return "", fmt.Errorf("decode hex string failed: %s", err.Error())
	}
	v := uint64(binary.LittleEndian.Uint64(b))
	name := eos.NameToString(v)
	if len(name) > 12 {
		name = name[0:12]
	}
	name = strings.Replace(name, ".", "x", -1)
	return name, nil
}
