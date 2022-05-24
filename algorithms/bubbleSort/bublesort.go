/*
Write a Bubble Sort program in Go.
The program should prompt the user to type in a sequence of up to 10 integers.
The program should print the integers out on one line, in sorted order, from least to greatest.
Use your favorite search tool to find a description of how the bubble sort algorithm works.

As part of this program, you should write a function called BubbleSort() which takes a slice of integers as an argument and returns nothing.
The BubbleSort() function should modify the slice so that the elements are in sorted order.

A recurring operation in the bubble sort algorithm is the Swap operation which swaps the position of two adjacent elements in the slice.
You should write a Swap() function which performs this operation.
Your Swap() function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice.
The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

	//The function for the bubble sort algorithm
func BubbleSort(rawSlice []int) []int {

	//Initiate the loop in reverse order until the last element in the slice
	for i := len(rawSlice); i > 0; i-- {
	
	//The second loop goess over the slice and swap the integers
	   for j := 1; j < i; j++ {

	//Swap algorithm that call the Swap function if previous number in the slice is bigger than actual
		  if rawSlice[j-1] > rawSlice[j] {
			 	Swap(rawSlice, j)
				}
			}
	   }

	//Return the ordered slice
	return rawSlice

  	}
  	//The function to swap the order of numbers
func Swap(rawSlice []int, j int) {
	//Declare a temporary variable to swap the order
	var swaper = rawSlice[j]
	rawSlice[j] = rawSlice[j-1]
	rawSlice[j-1] = swaper

	//Go alternative
	//rawSlice[j], rawSlice[j-1] = rawSlice[j-1], rawSlice[j]
	}

	//BubbleSort faster for go
/*
func BubbleSortOpt(rawSlice []int) {
	for sweepN := 0; sweepN < len(rawSlice); sweepN++ {
		swapped := false

		for i :=0; i < len(rawSlice)-1-sweepN; i++ {
			if rawSlice[i] > rawSlice[i+1] {
				rawSlice[i], rawSlice[i+1] = rawSlice[i+1], rawSlice[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
*/

	//Main function of the bubble sort algorithm program
func main() {

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
    
	//Print the results using the BubbleSort function
    fmt.Println("The sorted sequence looks like this:", BubbleSort(newSlice))

}