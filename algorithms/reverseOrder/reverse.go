package main

import (
	"fmt"
	"bufio"
	"os"
)



//The most efficient way (can handle large Chinise words)
func reverse(entry string) string {
	var result string
	for _, w := range entry {
		result = string(w) + result
	}
	return result
}

/*
//The function start reading from the end of the entry
func reverse(entry string) string {
	var result string
	for i := len(entry)-1; i >= 0; i-- {
		result = result + string(entry[i])
	}
	return result
}


//The same using string builder
func reverse(entry string) string {
	var result strings.Builder
	for i := len(entry)-1; i >= 0; i-- {
		result.WriteByte(entry[i])
	}
	return result.String()
}
//Simple for loop
func reverse(entry string) string {
	var result string
	for i := 0; i< len(entry); i++ {
		result = string(entry[i]) + result
	}
	return result
}
*/

//Create a reusable function to read the entries in the console
func scanTerminal() string {
	
	fmt.Println("Please enter a word or a phrase to reverse:")

	//create the variable to be returned with the inserted text
	var entryText string
	
	//create the standart scanner for the entries in the console
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		entryText = scanner.Text()
	}

	//return the entry
	return entryText
}

func main(){

	fmt.Print("\nYour reverse entry is:\n", reverse(scanTerminal()))

}