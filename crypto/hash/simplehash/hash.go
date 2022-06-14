package main

import (
	"fmt"
	"math/bits"
)

/*
For each byte in the input:
    Rotate bits left 3
    Shift bits left 2

Create a final array of length 4
Each index in final array is the XOR of all indexes modulo i in the input
For example: [0^4^8..., 1^5^9..., 2^6^10..., 3^7^11...]
*/

func hash(input []byte) [4]byte {
	for i, b := range input {
		rotated := bits.RotateLeft(uint(b), 3)
		input[i] = byte(rotated)
	}

	for i, b := range input {
		input[i] = b << 2
		// shift by 2 and save result back into "input"
	}

	final := [4]byte{}
	for i, b := range input {
		// final[i%4] = final[i%4] XOR b
		final[i%4] = final[i%4] ^ b
	}

	return final
}

func main() {
	fmt.Printf("%X\n", hash([]byte("Lane Wagner")))
	fmt.Printf("%X\n", hash([]byte("Lane Wagner")))
	fmt.Printf("%X\n", hash([]byte("Lane Swagner")))
	fmt.Printf("%X\n", hash([]byte("Bootdev")))
	fmt.Printf("%X\n", hash([]byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")))
}
