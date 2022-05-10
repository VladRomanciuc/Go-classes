/*
Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.

Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
*/

package main

import (
	"fmt"
	"io"
	"bufio"
	"strings"
	"os"
)

//Create the struct of the entry
type Person struct {
	fname string
	lname string 
}

//Function to restrict the length to 20 characters
func maxFieldLength(field string) string {
    if len(field)>20 { return string(field[0:20]) 
		} else { return field }
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}

//Main logic of the program
func main() {

	fmt.Println("Place the file in the folder with source code and change the full name of the file to names.txt.")

	//Open the file name with an error catcher
	file, err := os.Open("c:/Users/alina/Desktop/Go classes/week4/read/names.txt")
    check(err)
	defer file.Close()

	//Declare an array of Person structure that will record the scaned names
	var people[] Person
	
	//Read the file
	scanner := bufio.NewScanner(file)

	//Iterate each line and find the name and surname. 
    for scanner.Scan() {
        line := scanner.Text()
            scannedName :=strings.Split(string(line)," ")
			//Map results in the people array
            fullName := Person {
                maxFieldLength(scannedName[0]),
                maxFieldLength(scannedName[1]),
            } 
            people = append(people, fullName)
 	
	//Error catcher	
			if err != nil || io.EOF == err {
            	break
        	}  
    	}  
	//Print the list of names
	for _, person := range people {
			fmt.Printf("\tFirst name: %v \tLast Name: %v \n", person.fname, person.lname)
	}
}
