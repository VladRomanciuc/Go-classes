package main

import "fmt"

func otp(msg, key []byte) []byte {
	res := []byte{}
	for i := range key {
		res = append(res, msg[i]^key[i])
	}
	return res
}

func main() {
	message1 := "hello world"
	message2 := "other hello"
	key := "thiskeyhere"
	cipherText1 := otp([]byte(message1), []byte(key))
	cipherText2 := otp([]byte(message2), []byte(key))
	// create cipherText 1 and cipherText2

	fmt.Println(cipherText1)
	fmt.Println(cipherText2)

	c1XORc2 := otp(cipherText1, cipherText2)
	// XOR the ciphertexts and print the result

	fmt.Println(c1XORc2)

	// print the result of XORing the original messages
	fmt.Println(otp([]byte(message1), []byte(message2)))
}
