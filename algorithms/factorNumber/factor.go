package main

import (
	"fmt"
	"log"
)

//Function that run 2 loops over and devide the number appending the divider as a result
func factor(primes []int, number int) []int {
	var result []int
	for _, prime := range primes {
		for number%prime == 0 {
			result = append(result, prime)
			number = number/prime
		}
	}
	if number > 1 {
		result = append(result, number)
	}
	return result
}

//Function to read the entry from console
func input() int {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func main() {
	primes := []int{2,3,4,5,6,7,8,9}
	fmt.Print("Please enter your number to factor:\n ")
	number := input()
	fmt.Print("\nThe primes are:", factor(primes, number))
}