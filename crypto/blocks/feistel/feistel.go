package main

import (
	"crypto/sha256"
	"fmt"
)

func feistel(msg, key []byte, rounds int) []byte {
	// split the msg up into two equal parts
	lhs := msg[:len(msg)/2]
	rhs := msg[len(msg)/2:]

	// pseudocode:
	// for each round:
	//   nextRHS = xor(lhs, hash(rhs+key))
	//   nextLHS = oldRHS
	for i := 0; i < rounds; i++ {
		nextLeft := rhs

		hashedData := hash(rhs, key, len(key))

		rhs = xor(lhs, hashedData)

		lhs = nextLeft

		// lhs, rhs = rhs, xor(lhs, hash(rhs, key, len(key)))
	}
	return append(rhs, lhs...)
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

func xor(lhs, rhs []byte) []byte {
	res := []byte{}
	for i := range lhs {
		res = append(res, lhs[i]^rhs[i])
	}
	return res
}

// outputLength should be equal to the key length
// when used in feistel so that the XOR operates on
// inputs of the same size
func hash(lhs, rhs []byte, outputLength int) []byte {
	h := sha256.New()
	h.Write(append(lhs, rhs...))
	return h.Sum(nil)[:outputLength]
}