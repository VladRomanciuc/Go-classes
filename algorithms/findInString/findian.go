/*
Write a program which prompts the user to enter a string. 
The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’. 
The program should print “Found!” if the entered string starts with the character ‘i’, 
ends with the character ‘n’, and contains the character ‘a’. The program should print “Not Found!” otherwise. 
The program should not be case-sensitive, so it does not matter if the characters are upper-case or lower-case.

Examples: The program should print “Found!” for the following example entered strings, 
“ian”, “Ian”, “iuiygaygn”, “I d skd a efju N”. 
The program should print “Not Found!” for the following strings, “ihhhhhn”, “ina”, “xian”. 
*/

package main

import (
	"fmt"
	"os"
	"bufio"
	
)

import "strings"

func main() {
	
	//print the message for entering the text
	fmt.Println("Enter a string:") 
	
	//Scan the all entered text
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	//transform to text
	textInput := scanner.Text()
	
	// Confirm that the text was read correctly
	fmt.Printf("\tYour entered string is: \t\"%v\" \n", textInput) 
	
	//transform the input to lower case
	textInputLower := strings.ToLower(textInput)
	fmt.Printf("\tTransformed into lower case: \t%v \n", textInputLower)

	//filter the empty spaces
	textInputLowerNoSpace := strings.ReplaceAll(textInputLower, " ", "")
	fmt.Printf("\tThe input with fitered spaces: \t%v \n", textInputLowerNoSpace)
	
	//the logic for the task conditions. Scan the input to find a, i at the start and n at the end, 
	if strings.Contains(textInputLowerNoSpace, "a") { 
		if strings.HasPrefix(textInputLowerNoSpace, "i") { 
			if strings.HasSuffix(textInputLowerNoSpace, "n") { 
	
	//print Found message
				fmt.Println("\tFound!!! The input starts with the character `i`, ends with the character `n`, and contains the character `a`.\t")
	
	//end of execution
				os.Exit(0) 
			}
		}
	}
    //print Not found message
	fmt.Println("\tSorry... Not Found!") 
}
