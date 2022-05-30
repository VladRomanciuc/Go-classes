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

//int generator
func makeRange(min, max int) []int {
    a := make([]int, max-min+1)
    for i := range a {
        a[i] = min + i
    }
    return a
}

func main () {
	data := makeRange(1, 100)
	FizzBuzz(data)
}

/* simple loop in main function
func main() {
	for i := 1; i < 101; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else {
			fmt.Println(i)
		}
	}
}

*/