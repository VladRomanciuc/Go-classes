/*
Write two goroutines which have a race condition when executed concurrently. 
Explain what the race condition is and how it can occur.
*/

package main

import (
	"fmt"
	"time"
)

// Declare our lovely variable (the value on initialisation is 0)
var i int

//The function that will add +1 to initial state and print it
func race() {
	i++
	fmt.Println("The number is: ", i)
}

//The main function of the program
func main() {
	//Call 2 time the race function (expected to be executed at the same time??? with the same output)
	go race()
	go race()
	//give 1 spare second to display the result
	time.Sleep(1 * time.Second)

}

/*----------------------------------------------------------------

The result of firt execution:
The number is:  2
The number is:  1

The result of the second execution
The number is:  1
The number is:  2

Different behaviour on each time when program is executed
*/

