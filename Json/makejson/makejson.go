/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
*/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

//Create a reusable function to read the entries in the console
func scanTerminal(entry string) string {
	
	fmt.Println(entry)

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

func main() {
	//Create the map
   	m := make(map[string] string)

	//Adding the name and address to the map using the keys “name” and “address”, respectively.
	m["name"] = scanTerminal("Please, enter the name of the person: ")
	m["address"] = scanTerminal("And now its address: ")

	// Transform map to json using Marshal()
	jsonTransform, err := json.Marshal(m)
	
	//Error catcher if any occurs
	if err != nil {
		fmt.Printf("Ups... there is a error related to: %s", err)
		os.Exit(0)
	}

	//Print the transformed JSON object
	fmt.Printf("The json looks like this: %v", string(jsonTransform))

}
