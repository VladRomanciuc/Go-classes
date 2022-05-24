package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
	"log"
)

func insertion(slice []int, number int) []int {
	var sorted []int
	for _, number := range slice {
		sorted = insert(sorted, number)
	}	
	for index, val := range sorted {
		slice[index] = val
	}
	return sorted
}

func insert(sorted []int, number int) []int {
	for index, sortedNumber := range sorted {
		if number < sortedNumber {
			return append(sorted[:index], append([]int{number}, sorted[index:]...)...)
		}
	}
	return append(sorted, number)
}


func inputSlice() []int {

	//Read the line of entered numbers in the terminal
    console := bufio.NewReader(os.Stdin)
    line, _, _ := console.ReadLine()

	//Separate the numbers, looking for a space between
    separatedLine := strings.Split(string(line), " ")

	//Declare a variable type slice to store the numbers
    var newSlice []int

	//Transform the numbers from string to integers and add them to the slice
    for _, number := range(separatedLine) {
      	transformedNumber, _ := strconv.Atoi(number)
      	newSlice = append(newSlice, transformedNumber)
    	}
	return newSlice
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
	fmt.Println("Waiting for the entry of a series of integers separated by space (up to 10 preferable): ")
	slice := inputSlice()

	fmt.Println("Please enter the number to insert: ")
	number := input()

	result := insertion(slice, number)
	fmt.Println("Your new series with inserted number is: ", result)
}