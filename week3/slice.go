/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice. 
The program should be written as a loop. 
Before entering the loop, the program should create an empty integer slice of size (length) 3. 
During each pass through the loop, the program prompts the user to enter an integer to be added to the slice. 
The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
The slice must grow in size to accommodate any number of integers which the user decides to enter. 
The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"bufio"
)

func main() {
	
	//create an empty integer slice of size (length) 3
	var mainSlice  []int = make([]int, 3)
	fmt.Printf("The slice is with lenght of %v and capacity of %v.\n", len(mainSlice), cap(mainSlice))

	//declare the variable for input in the console
	var entry string

	//entering the loop
	for {

	//print the input bar
		fmt.Printf("Add a number to the slice or `x` to finish: ")
	
	//scan entries
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan() 
		entry = scanner.Text()

	//quit the loop when the user enters the character ‘X’
		if strings.ToUpper(entry)[0] == 'X' {
			fmt.Println("Almost finishing...")
			os.Exit(0)
		}

	//filter for numbers
		newEntry, err := strconv.Atoi(entry)
		if err != nil {
			fmt.Println("A number is expected...")
            continue
		}
	
	//add a number to the slice
		mainSlice = append(mainSlice, newEntry)
		sort.Ints(mainSlice[:])

		fmt.Printf("This is the sorted slice %v\n", mainSlice)

	}
}

