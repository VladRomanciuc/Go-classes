package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func main() {
	const encryptionKey = "kjhgfdsaqwertyuioplkjhgfdsaqwert"
	const iv = "1234567812345678"

	const messageToSend = "Hey there, let's start a coup or something"

	cipherText := encrypt(messageToSend, encryptionKey, iv)

	decryptedText := decrypt(cipherText, encryptionKey, iv)
	fmt.Println(decryptedText)
}

func encrypt(plainText, key, iv string) string {
	bytes := []byte(plainText)
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(blockCipher, []byte(iv))
	stream.XORKeyStream(bytes, bytes)
	return string(bytes)
}

func decrypt(cipherText, key, iv string) string {
	blockCipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	stream := cipher.NewCTR(blockCipher, []byte(iv))
	bytes := []byte(cipherText)
	stream.XORKeyStream(bytes, bytes)
	return string(bytes)
}