package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

//AES encryption/decryption

//the encrypt function takes a key, a message and a nonce to encrypt using buildin aes and cipher packages
func encrypt(key, plaintext, nonce []byte) (ciphertext []byte, err error) {
	//generate a new aes cipher using provided key and return an error if it is the case
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//generate the 128-bit block
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	//encrypt the message using a nonce and the Seal function of aes package
	ciphertext = aesgcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nil
}

//the decrypt function it is reversed encrypt
func decrypt(key, ciphertext, nonce []byte) (plaintext []byte, err error) {
	//generate a new aes cipher using the provided key and return an error if it is the case
	cip, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//generate the 128-bit block
	gcm, err := cipher.NewGCM(cip)
	if err != nil {
		return nil, err
	}
	
	// return decrypted message using the provided nonce and Open function
	return gcm.Open(nil, nonce, ciphertext, nil)
}

//main logic
func main() {
	// Select AES-256 by using 32 character key (256 bits)
	key := []byte("AES256Key-32Characters1234567890")
	//the message to be encrypted
	plaintext := []byte("this message is meant for myself, and only myself")

	//declaring nonce for the encryption (normaly should be generated)
	nonce := []byte("twelvebytess")

	//call encrypt
	ciphertext, err := encrypt(key, plaintext, nonce)
	if err != nil {
		panic(err)
	}
	fmt.Println(ciphertext)

	//call decrypt
	decryptedText, err := decrypt(key, ciphertext, nonce)
	if err != nil {
		panic(err)
	}
	//Print result
	fmt.Println(string(decryptedText))
}
