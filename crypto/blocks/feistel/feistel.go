package main

import (
	"crypto/sha256"
	"fmt"
)
//feistel encryption using buildin sha256 package

//the function takes message, a key and the number of encryption rounds
func feistel(msg, key []byte, rounds int) []byte {
	// split the msg up into two equal parts
	lefths := msg[:len(msg)/2]
	righths := msg[len(msg)/2:]

	//for loop for each round
	for i := 0; i < rounds; i++ {
		//starting from with the right part
		next := righths
		//encrypt message using hash function and the key
		hashedData := hash(righths, key, len(key))
		//xor on nearby message
		righths = xor(lefths, hashedData)
		//swaping messages for next round
		lefths = next
	}
	return append(righths, lefths...)
}

// outputLength should be equal to the key length
// when used in feistel so that the XOR operates on
// inputs of the same size
func hash(lhs, rhs []byte, outputLength int) []byte {
	h := sha256.New()
	h.Write(append(lhs, rhs...))
	return h.Sum(nil)[:outputLength]
}

//XOR function works on every byte in the  messages
func xor(lhs, rhs []byte) []byte {
	res := []byte{}
	for i := range lhs {
		res = append(res, lhs[i]^rhs[i])
	}
	return res
}

func main() {
	plaintext := []byte("Plain texts friend")
	// normally we would use a key schedule to
	// break this key up, but this is a simple
	// toy network so we will just use the master
	// key as each round key
	key := []byte("thesecret")

	encrypted := feistel(plaintext, key, 8)
	fmt.Println(encrypted)

	decrypted := feistel(encrypted, key, 8)
	fmt.Println(string(decrypted))
}