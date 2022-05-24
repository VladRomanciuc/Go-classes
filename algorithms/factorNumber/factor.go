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

//GCD Greatest common divisor - Euclidean algorithm
//
func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a % b
	}
	return a
}

func main() {
	fmt.Print("Please select an action:\n1. Factoring a number\n2. Greatest common divisor\n")
	
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	
	selector := input()

	switch {
	case selector == 1:
		fmt.Print("Please enter your number to factor: ")
		number := input()
		fmt.Print("\nThe primes are:", factor(primes, number))
	
	case selector == 2:
		fmt.Print("Please enter your first number to find GCD: ")
		first := input()
		fmt.Print("Please enter your second number to find GCD: ")
		second := input()
		fmt.Print("\nThe GCD is: ", GCD(first, second))

	default:
		fmt.Print("\nUps.... try again.")
	}
	
}