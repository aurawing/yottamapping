package cmd

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAESEncAndDec(t *testing.T) {
	plaintext := "1234567890"
	key := "Aa111111"
	ciphertext, err := AES256GCMEncrypt([]byte(plaintext), key)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(hex.EncodeToString(ciphertext))
	plaintext2, err := AES256GCMDecrypt(ciphertext, key)
	if err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(plaintext2))
}
