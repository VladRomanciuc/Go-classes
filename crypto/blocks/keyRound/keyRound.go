package main

import "fmt"

/*
key schedule's algorithm will accept a master key ([4]byte) and a round number (int) as inputs. The key schedule will then produce a round key where each byte in the round key is the original byte from the master key XORed with the binary representation of the round number.
*/

func deriveRoundKey(masterKey [4]byte, roundNumber int) [4]byte {
		for i, v := range masterKey {
		masterKey[i] = v ^ byte(roundNumber)
	}
	return masterKey
}

func main() {
	fmt.Println(deriveRoundKey([4]byte{0xAA, 0xFF, 0x11, 0xBC}, 1))
	fmt.Println(deriveRoundKey([4]byte{0xAA, 0xFF, 0x11, 0xBC}, 2))
	fmt.Println(deriveRoundKey([4]byte{0xAA, 0xFF, 0x11, 0xBC}, 3))
	fmt.Println(deriveRoundKey([4]byte{0xAA, 0xFF, 0x11, 0xBC}, 4))

	fmt.Println(deriveRoundKey([4]byte{0xAA, 0xBC, 0xFF, 0xBC}, 1))
	fmt.Println(deriveRoundKey([4]byte{0xFF, 0xFF, 0x11, 0xBC}, 1))
	fmt.Println(deriveRoundKey([4]byte{0xFF, 0xBC, 0xFF, 0x11}, 1))
	fmt.Println(deriveRoundKey([4]byte{0xBC, 0xFF, 0x11, 0xBC}, 1))
	fmt.Println(deriveRoundKey([4]byte{0x11, 0x11, 0x11, 0xBC}, 1))
	fmt.Println(deriveRoundKey([4]byte{0xAA, 0xAA, 0xFF, 0xBC}, 1))
	fmt.Println(deriveRoundKey([4]byte{0xAA, 0xFF, 0x11, 0xAA}, 1))
}
