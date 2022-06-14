package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

func main() {
	const pass1 = "thisIsAPassword"
	hashed, _ := hashPassword(pass1)
	fmt.Println(hashed[:4])
	fmt.Println(checkPasswordHash(pass1, hashed))

	const pass2 = "thisIsAnotherPassword"
	hashed, _ = hashPassword(pass2)
	fmt.Println(hashed[:4])
	fmt.Println(checkPasswordHash(pass2, hashed))

	hashed, _ = hashPassword("anotherPass")
	fmt.Println(hashed[:4])
	fmt.Println(checkPasswordHash("shouldntMatch", hashed))
}
