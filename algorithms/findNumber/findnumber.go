package main

import (
	"fmt"
	"log"
)
//Function to read the entry from console
func input() int {
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

//The number finder function that return true or false
func finder(data []int, number int) bool {
	for _, val := range data {
		if val == number {
			fmt.Println("Your number match our records!")
			return true
		}
	}
	fmt.Println("Ups... there are no matches!")
	return false
}

//main logic
func main () {
	data := []int {1, 2, 3, 4, 5}
	
	fmt.Println("Your number: ")

	finder(data, input())
	
}

