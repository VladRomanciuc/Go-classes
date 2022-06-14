package main

import "fmt"

func encrypt(plaintext string, key int) string {
	final := ""
	for _, c := range plaintext {
		offset := findOffset(c, key)
		final += offset
	}
	return final
}

func decrypt(ciphertext string, key int) string {
	final := ""
	for _, c := range ciphertext {
		offset := findOffset(c, -key)
		final += offset
	}
	return final
}

// findOffset takes a character from the alphabet and returns the
// character at the positive or negative offset.
// wrapping is supported
func findOffset(c rune, offset int) string {
	const alphabet = "abcdefghijklmnopqrstuvwxyz"
	for i, curr := range alphabet {
		if curr == c {
			modI := (i + offset) % len(alphabet)
			if modI < 0 {
				modI = len(alphabet) + modI
			}
			return string(alphabet[modI])
		}
	}
	return " "
}

func main() {
	raw := "hello world"
	encrypted := encrypt(raw, 5)
	fmt.Println(encrypted)
	decrypted := decrypt(encrypted, 5)
	fmt.Println(decrypted)

	raw = "bootdev knows the crypto"
	encrypted = encrypt(raw, 19)
	fmt.Println(encrypted)
	decrypted = decrypt(encrypted, 19)
	fmt.Println(decrypted)

	raw = "wagslane"
	encrypted = encrypt(raw, 57)
	fmt.Println(encrypted)
	decrypted = decrypt(encrypted, 57)
	fmt.Println(decrypted)

	raw = "this one would be really hard to do by hand for a number of reasons"
	encrypted = encrypt(raw, 43255435)
	fmt.Println(encrypted)
	decrypted = decrypt(encrypted, 43255435)
	fmt.Println(decrypted)

	raw = "lorem ipsum is simply dummy text of the printing and typesetting industry lorem ipsum has been the industrys standard dummy text ever since the when an unknown printer took a galley of type and scrambled it to make a type specimen book"
	encrypted = encrypt(raw, 4325543588768787)
	fmt.Println(encrypted)
	decrypted = decrypt(encrypted, 4325543588768787)
	fmt.Println(decrypted)
}
