package main

import (
	"fmt"
	"log"
)

//Fuction stores the 3 workable values that save the operational time of the loop
func fibLoop(n int) int {
	//For n=0 or 1 returns 0 1 1
	if n <= 1 {
		return n
	}
	//initial values of the 2 previous numbers
	n2 := 0
	n1 := 1
	var actual int
	//the loop return the value of n in the fibonacci sequence
	for i := 2; i <= n; i++ {
		actual = n2 + n1
		n2 = n1
		n1 = actual
	}
	return actual
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
	fmt.Print("Please enter the number for which you want to know the fibonacci value: ")
	n := input()
	fmt.Printf("\nThe fibonacci value is: %v.", fibLoop(n))
}