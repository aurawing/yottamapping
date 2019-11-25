package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"io"
)

// //AesEncryptSimple AES加密
// func AesEncryptSimple(origData []byte, key string, iv string) ([]byte, error) {
// 	return AesDecryptPkcs5(origData, []byte(key), []byte(iv))
// }

// func AesEncryptPkcs5(origData []byte, key []byte, iv []byte) ([]byte, error) {
// 	return AesEncrypt(origData, key, iv, PKCS5Padding)
// }

// func AesEncrypt(origData []byte, key []byte, iv []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	blockSize := block.BlockSize()
// 	origData = paddingFunc(origData, blockSize)

// 	blockMode := cipher.NewCBCEncrypter(block, iv)
// 	crypted := make([]byte, len(origData))
// 	blockMode.CryptBlocks(crypted, origData)
// 	return crypted, nil
// }

// //AesDecryptSimple AES解密
// func AesDecryptSimple(crypted []byte, key string, iv string) ([]byte, error) {
// 	return AesDecryptPkcs5(crypted, []byte(key), []byte(iv))
// }

// func AesDecryptPkcs5(crypted []byte, key []byte, iv []byte) ([]byte, error) {
// 	return AesDecrypt(crypted, key, iv, PKCS5UnPadding)
// }

// func AesDecrypt(crypted, key []byte, iv []byte, unPaddingFunc func([]byte) []byte) ([]byte, error) {
// 	block, err := aes.NewCipher(key)
// 	if err != nil {
// 		return nil, err
// 	}
// 	blockMode := cipher.NewCBCDecrypter(block, iv)
// 	origData := make([]byte, len(crypted))
// 	blockMode.CryptBlocks(origData, crypted)
// 	origData = unPaddingFunc(origData)
// 	return origData, nil
// }

// func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
// 	padding := blockSize - len(ciphertext)%blockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(ciphertext, padtext...)
// }

// func PKCS5UnPadding(origData []byte) []byte {
// 	length := len(origData)
// 	unpadding := int(origData[length-1])
// 	if length < unpadding {
// 		return []byte("unpadding error")
// 	}
// 	return origData[:(length - unpadding)]
// }

func AES256GCMDecrypt(ciphertext []byte, key string) []byte {
	keyBytes := Sha256Sum([]byte(key))
	c, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		panic(err)
	}
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	return plaintext
}

func AES256GCMEncrypt(plaintext []byte, key string) []byte {
	keyBytes := Sha256Sum([]byte(key))
	c, err := aes.NewCipher(keyBytes)
	if err != nil {
		panic(err)
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err)
	}
	return gcm.Seal(nonce, nonce, plaintext, nil)
}

func Sha256Sum(bytes []byte) []byte {
	h := sha256.New()
	h.Write(bytes)
	return h.Sum(nil)
}
