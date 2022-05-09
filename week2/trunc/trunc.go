/*
Task
Write a program which prompts the user to enter a floating point number and prints the integer which is a truncated version of the floating point number that was entered. Truncation is the process of removing the digits to the right of the decimal place.
Submit your source code for the program, “trunc.go”.
*/

package main

import (
	"fmt"
	"log"
)

// declare 2 variable
var inputFloatPointNumber float64

var truncatedFloatPointerNumber int64

func main() {

	//the input for the float number
	fmt.Println("Enter a Floating number:")
	_, err := fmt.Scan(&inputFloatPointNumber)

	if err != nil {
		log.Printf("[Error] Read more about float numbers!") //error message
	}

	fmt.Printf("The number I have read is '%v'.\n", inputFloatPointNumber) //prints in terminal the first variable

	//transform the number from float to integer
	truncatedFloatPointerNumber = int64(inputFloatPointNumber)

	fmt.Printf("The truncated number is '%v'.\n", truncatedFloatPointerNumber) //prints in terminal the second variable

}
