package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"

)

//Declare a simple loop function to go over a slice and sum numbers returning total
func sum (slice []int) int {
	total :=0
	for _, val := range slice {
		total = total + val
	}
	return total
}
/* Version 2 of sum (more calls to memory) 'Recursion'
func sum (slice []int) int {
	if len(slice) == 0 {
		return 0
	}
	return slice[0] + sum(slice[1:])
}
*/

//Function to read the entries in the console
func readInput() []int {

	fmt.Println("Waiting for the entry of a series of integers separated by space (up to 10 preferable):")

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

//Main function to print the result by calling sum and readInput functions
func main(){

	fmt.Print("The sum of entered integers is: ", sum(readInput())) 
	
}