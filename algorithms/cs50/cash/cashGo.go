package main

import (
	"fmt"
)

func cash(input int) {
	coins := []int {25, 10, 5, 1}
	res := []int {}

	for i:=0; i<len(coins); i++ {
		n := input/coins[i]
		res = append(res, n)
		fmt.Printf("Coins %v: %v\n", coins[i], n)
		input = input - coins[i]*n
	} 

	sum := 0
	for n:=0; n <len(res); n++{
		sum = sum + res[n]
	}
	fmt.Printf("Total coins owed: %v\n", sum)
}


func main() {

	var answer int
	for { 
		fmt.Println("Please, enter the amount owed: ")
		fmt.Scanln(&answer)
		if answer >= 0 {
			cash(answer)
		} else {
			fmt.Println("Sorry, out of range...")
		}
		fmt.Println("-----------------------------")
	}
}