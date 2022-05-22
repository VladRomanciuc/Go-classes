package main

import (
	"fmt"
)

//Function for fizzbuzz printing when an number is divided by 3 and 5
func printFizzBuzz(number int) {
	switch {
	case number%3 == 0 && number%5 == 0:
		fmt.Print("Fizz Buzz")
	case number%3 == 0:
		fmt.Print("Fizz")
	case number%5 ==0:
		fmt.Print("Buzz")
	default:
		fmt.Print(number)
	}
}

// FizzBuzz loop to print the numbers
func FizzBuzz(number []int) {
	for i:=1; i<=len(number); i++ {
		if i<len(number) {
		printFizzBuzz(number[i-1])
		fmt.Print(", ")
		} else if i == len(number) {
			printFizzBuzz(number[i-1])
			fmt.Print(" ")
		}
	}
}

func main () {
	data := []int{8,9,10,11,23,45,67,123,4456,45,3,4,5,23}
	FizzBuzz(data)
}