package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"os/exec"
	)

//Function to return the ping command and params + error handling
func command (params string)string{
	if c, err := exec.Command("cmd", "/c", "ping", "-n", "3", params).CombinedOutput()
		err != nil {
			s := "Error! Command execution failed"
				return (s)
			} else {
				return string(c)
			}
	}

//The main logic of the program
func main() {
	//Print the info to start the program
	fmt.Println("Ping your VM or any IP address")
	fmt.Println("------------------------------")
	fmt.Println("Enter the command (like `connect 172.168.0.1` or `exit` to end the program).")
	fmt.Println("Waiting for your entry: ")

	//Declare a variable for the console scanner
 	scanner := bufio.NewScanner(os.Stdin)

	//Initialise the loop to read the commands from console
 	for scanner.Scan() {
		//Declare the variable for the raw input
	 	userInput := scanner.Text()

		//Condition to end the program
 		if strings.Compare("exit", userInput) == 0 {
			os.Exit(0)
		}

		fmt.Print("Getting data...")

		//Variable to store the command
		var commandToEx[] string = strings.Split(userInput, " ")

		//Condition for the connect command
		if strings.Compare("connect", commandToEx[0]) == 0	{

			output := command(commandToEx[1])
				fmt.Println(output)
				} else {
					fmt.Println("The CLI is too young to handle this command.")
			}
				fmt.Println("Waiting for your entry:")
	}
}
