package main

import "fmt"


//XOR operation (exclusive or) on each byte
func encrypt(plaintext, key []byte) []byte {
	final := []byte{}
	for i := range plaintext {
		final = append(final, plaintext[i]^key[i])
	}
	return final
}

func decrypt(ciphertext, key []byte) []byte {
	final := []byte{}
	for i := range ciphertext {
		final = append(final, ciphertext[i]^key[i])
	}
	return final
}

func main() {
	const key = "mysecurepass"
	const message = "I'm lovin it"
	cipher := encrypt([]byte(message), []byte(key))
	fmt.Println(cipher)
	original := decrypt([]byte(cipher), []byte(key))
	fmt.Println(original)
	fmt.Println(string(original))
}
