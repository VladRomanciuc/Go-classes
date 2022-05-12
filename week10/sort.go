/*
Write a program to sort an array of integers. 
The program should partition the array into 4 parts, each of which is sorted by a different goroutine. 
Each partition should be of approximately equal size. 
Then the main goroutine should merge the 4 sorted subarrays into one large sorted array. 

The program should prompt the user to input a series of integers. 
Each goroutine which sorts ¼ of the array should print the subarray that it will sort. 
When sorting is complete, the main goroutine should print the entire sorted list.

*/

package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
	"sync"
)

//The function to take the inputs in the console and convert them intro a slice
func consoleEntry() []int {
	
	fmt.Println("Waiting for a series of integers separated by space to sort: ")
  	lineReader := bufio.NewReader(os.Stdin)
  	rawInput, _, _ := lineReader.ReadLine()
  	separatedInput := strings.Split(string(rawInput), " ")
  	
	var entrySlice []int
  	for _, i := range(separatedInput) {
    	numbers, _ := strconv.Atoi(i)
    	entrySlice = append(entrySlice, numbers)
  		}

		return entrySlice
	}

//The function for the bubble sort algorithm
func BubbleSort(rawSlice []int, wg *sync.WaitGroup) []int {

	defer wg.Done()

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
	}

	
//****************************************************************************************************************

func main () {
	
	//Declare a variable slice to store the numbers entered
		entrySlice := consoleEntry()
		fmt.Println("---------------------")
		fmt.Println("\nThe array to sort is: ", entrySlice)
		fmt.Println("\n---------------------")
	//Initialise the wait group needed to work with goroutines
		var wg sync.WaitGroup

	//Divide the slice into 4 sequences for the 1 phase of goroutines
		size := len(entrySlice)/4
		a := entrySlice[: 1*size]
		b := entrySlice[1*size : 2*size]
		c := entrySlice[2*size : 3*size]
		d := entrySlice[3*size : 4*size]
	
	//Add 4 goroutines jobs to sort the sequences
		wg.Add(4)

		go BubbleSort(a, &wg)
  		go BubbleSort(b, &wg)
  		go BubbleSort(c, &wg)
  		go BubbleSort(d, &wg)

	//Waiting for execution
		wg.Wait()

	//Initialise a lice for results and merge the sequences
		resultSlice := append(append(append(a, b...), c...), d...)
	
	//Add a new sorting job for the goroutines
		wg.Add(1)
	
		go BubbleSort(resultSlice, &wg)
	
	//Waiting for execution	
		wg.Wait()

	//Print the sorted slice
		fmt.Println("\nSorted by Go routine and looks like this:", resultSlice)
		fmt.Println("\n---------------------")
}