package cmd

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAESEncAndDec(t *testing.T) {
	plaintext := "1234567890"
	key := "Aa111111"
	ciphertext := AES256GCMEncrypt([]byte(plaintext), key)
	fmt.Println(hex.EncodeToString(ciphertext))
	plaintext2 := AES256GCMDecrypt(ciphertext, key)
	fmt.Println(string(plaintext2))
}
