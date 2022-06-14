package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func encrypt(key, plaintext, nonce []byte) (ciphertext []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {
	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(cip)
	if err != nil {
		return nil, err
	}
	
	
	return gcm.Open(nil, nonce, ciphertext, nil)
}

func main() {
	// Select AES-256 by using 32 character key (256 bits)
	key := []byte("AES256Key-32Characters1234567890")
	plaintext := []byte("this message is meant for myself, and only myself")

	// here we are hard-coding a nonce... this is a BAD IDEA, but are
	// doing it because our answer key is deterministic ;)
	// normally this should be randomly generated for EACH encryption
	nonce := []byte("twelvebytess")

	ciphertext, err := encrypt(key, plaintext, nonce)
	if err != nil {
		panic(err)
	}
	fmt.Println(ciphertext)
	decryptedText, err := decrypt(key, ciphertext, nonce)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decryptedText))
}
