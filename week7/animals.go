/*
Write a program which allows the user to get information about a predefined set of animals. 
Three animals are predefined, cow, bird, and snake. Each animal can eat, move, and speak. 
The user can issue a request to find out one of three things about an animal: 
1) the food that it eats, 
2) its method of locomotion, and 
3) the sound it makes when it speaks. 
The following table contains the three animals and their associated data which should be hard-coded into your program.

Animal - cow, bird, snake

Food eaten - grass, worms, mice

Locomotion method - walk, fly, slither

Spoken sound - moo, peep, hsss 

Your program should present the user with a prompt, “>”, to indicate that the user can type a request. 
Your program accepts one request at a time from the user, prints out the answer to the request, and prints out a new prompt.
Your program should continue in this loop forever. 
Every request from the user must be a single line containing 2 strings. 
The first string is the name of an animal, either “cow”, “bird”, or “snake”. 
The second string is the name of the information requested about the animal, either “eat”, “move”, or “speak”. 
Your program should process each request by printing out the requested data.

You will need a data structure to hold the information about each animal. 
Make a type called Animal which is a struct containing three fields:food, locomotion, and noise, all of which are strings. 
Make three methods called Eat(), Move(), and Speak(). 
The receiver type of all of your methods should be your Animal type. 
The Eat() method should print the animal’s food, the Move() method should print the animal’s locomotion, 
and the Speak() method should print the animal’s spoken sound. 
Your program should call the appropriate method when the user makes a request.
*/

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

	//Create a struct for the animal type
type Animal struct {
	food, locomotion, spoken string
}

	//Create a function of type Animal for food category
func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

	//Create a function of type Animal for locomotion category
func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

	//Create a function of type Animal for spoken category
func (animal Animal) Speak() {
	fmt.Println(animal.spoken)
}

	//Create the main function of the program
func main() {

	//Declare a variable of type map to hold info about animals
	data := map[string] Animal{
				"bird":  {"worms", "fly", "peep"},
				"cow":   {"grass", "walk", "moo"},
				"snake": {"mice", "slither", "hsss"},
			}

	//Initiate the forever loop
	for {
	//Initiate the request of info by printing the message
		fmt.Printf("To access the information about the interested animal, please enter its Name and Info requested separated by space:> ")
	
	//Read the line of entered numbers in the terminal
		console := bufio.NewScanner(os.Stdin)
		console.Scan()
		request := console.Text()
	
	//Declare two variable for name and info and assign them the results of the request
		name := strings.Split(request, " ")[0]
		info := strings.Split(request, " ")[1]
	
	//Case switcher for the name entry
		switch name {
			case "cow": data[info].Eat()
			case "bird": data[info].Eat()
			case "snake": data[info].Eat()
			default: fmt.Printf("Invalid input for `%s` (Try cow, bird or snake).", name)
			return
		}

	//Case switcher for the info entry
		switch info {
			case "eat": data[name].Eat()
			case "move": data[name].Move()
			case "speak": data[name].Speak()
			default: fmt.Printf("Invalid input for `%s` (Try eat, move or speak).", info)
			return
		}

	}
}