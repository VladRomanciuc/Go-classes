package main

import (
	"fmt"
	"crypto/sha256"
)

/*
1. Split key in two halves. Second half is larger if key is odd length
2. innerHex = hash(first + message)
3. resultHex = hash(second + innerHex)
4. return resultHex
*/

func hmac(message, key string) string {
	firstHalf := key[:len(key)/2]
	secondHalf := key[len(key)/2:]

	innerDat := sha256.Sum256([]byte(firstHalf + message))
	inner := fmt.Sprintf("%x", innerDat)

	secondDat := sha256.Sum256([]byte(secondHalf + inner))
	result := fmt.Sprintf("%x", secondDat)
	return result
}

func main() {
	fmt.Println(hmac("hey this is my message", "super_secret_password"))
	fmt.Println(hmac("that's the kind of password an idiot has on their luggage", "12345"))
	fmt.Println(hmac("thats the password I have on my luggage", "12345"))
	fmt.Println(hmac("{'email':'superswagz@hotmail.io}", "correct horse battery staple"))
}
