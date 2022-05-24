/*
Let us assume the following formula for
displacement s as a function of time t, acceleration a, initial velocity vo,
and initial displacement so.

s = ½ * a * t x 2 + vo x t + so

Write a program which first prompts the user
to enter values for acceleration, initial velocity, and initial displacement.
Then the program should prompt the user to enter a value for time and the
program should compute the displacement after the entered time.

You will need to define and use a function
called GenDisplaceFn() which takes three float64
arguments, acceleration a, initial velocity vo, and initial
displacement so. GenDisplaceFn()
should return a function which computes displacement as a function of time,
assuming the given values acceleration, initial velocity, and initial
displacement. The function returned by GenDisplaceFn() should take one float64 argument t, representing time, and return one
float64 argument which is the displacement travelled after time t.

For example, let’s say that I want to assume
the following values for acceleration, initial velocity, and initial
displacement: a = 10, vo = 2, so = 1. I can use the
following statement to call GenDisplaceFn() to
generate a function fn which will compute displacement as a function of time.

fn := GenDisplaceFn(10, 2, 1)

Then I can use the following statement to
print the displacement after 3 seconds.

fmt.Println(fn(3))

And I can use the following statement to print
the displacement after 5 seconds.

fmt.Println(fn(5))
*/

package main

import (
	"fmt"
)

	//Create the function for the formula with 3 arguments and a function for the time
func GenDisplaceFn(acceleration, initialVelocity, initialDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return 0.5*acceleration*time*time + initialVelocity*time + initialDisplacement
	}
}
	//Create the main function that read the initial values and prints final result
func main () {

	//Declare 3 variable that holds the initial values
	var acceleration, initialVelocity, initialDisplacement float64
	
	//Print message to screen and scan the entries in the terminal
  	fmt.Print("Please enter the initial value for the acceleration:")
  	fmt.Scanln(&acceleration)
  
  	fmt.Print("Please enter the initial value for the velocity:")
  	fmt.Scanln(&initialVelocity)
  
  	fmt.Print("Please enter the initial value for the displacement:")
  	fmt.Scanln(&initialDisplacement)

	//Declare a variable fn that holds the result of the GeDisplaceFn
	fn := GenDisplaceFn(acceleration, initialVelocity, initialDisplacement)

	//Print the results
	fmt.Println("The result after 3sec is:", fn(3))
	fmt.Println("The result after 5sec is:", fn(5))
}
