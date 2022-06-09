package main

/*
American Express uses 15-digit numbers, 
MasterCard uses 16-digit numbers, and 
Visa uses 13- and 16-digit numbers.

All American Express numbers start with 34 or 37; 
most MasterCard numbers start with 51, 52, 53, 54, or 55
all Visa numbers start with 4
*/

/*
Luhn’s Algorithm
So what’s the secret formula? Well, most cards use an algorithm invented by Hans Peter Luhn of IBM.
According to Luhn’s algorithm, you can determine if a credit card number is (syntactically) valid as follows:

Multiply every other digit by 2, starting with the number’s second-to-last digit, and then add those products’ digits together.
Add the sum to the sum of the digits that weren’t multiplied by 2.
If the total’s last digit is 0 (or, put more formally, if the total modulo 10 is congruent to 0), the number is valid!
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func input() []int {
	fmt.Println("Waiting for the entry of a card number: ")
	//Read the line of entered numbers in the terminal
    console := bufio.NewReader(os.Stdin)
    line, _, _ := console.ReadLine()
	
	var newSlice []int
		for i:=0; i<len(line); i++ {
			if (string(line[i]) != " " && string(line[i]) != "-"){
			number, _ := strconv.Atoi(string(line[i]))
			if number >= 0 {
				newSlice = append(newSlice, number)
			}
		}		
	}
	return newSlice
}

func checksum(slice []int) string {
//step 1
		var step1 []int
	for i := len(slice)-2; i >= 0; i -= 2 {	
		number := slice[i]*2
		if number > 9 {
			n := strconv.Itoa(number)
			for _, digit := range n {
				c, _ := strconv.Atoi(string(digit))
				step1 = append(step1, c)
			}
		} else {
			step1 = append(step1, number)	
		}
		
	}

//step 2
	var step2 []int
	for i := len(slice)-1; i >= 0; i -= 2 {
		step2 = append(step2, slice[i])
	}

	check := strconv.Itoa(sum(step1)+sum(step2))
	res := ""
	if strings.HasSuffix(check, "0") && check != "0" {
		res = "VALID"
		return res
	} else {
		res = "INVALID"
		return res
	}
}


func sum(slice []int) int {
	total :=0
	for _, val := range slice {
		total = total + val
	}
	return total
}

func clean(slice []int) string {
	card := strings.Builder{}
	for _, v := range slice {
		i := strconv.Itoa(v)
		card.WriteString(i)
	}
	return card.String()
}

func checkA(clean string) string {
	res := ""
	if len(clean) == 15 {
		if strings.HasPrefix(clean, "34") {
			res = "American Express"
		} else if strings.HasPrefix(clean, "37") {
			res = "American Express"
		} else {
			res = "INVALID"
		}
	}
	return res
}

func checkM(clean string) string {
	prefix := []string{"51", "52", "53", "54", "55"}
	res := "INVALID"
	for i := 0; i < len(prefix); i++ {
			if len(clean) == 16 && strings.HasPrefix(clean, prefix[i]) {
				res = "MasterCard"
			} 
		} 
		return res
	}

func checkV(clean string) string {
	res := ""
	if len(clean) == 13 {
		if strings.HasPrefix(clean, "4"){
			res = "Visa"
			}
		} else if len(clean) == 16 {
			if strings.HasPrefix(clean, "4"){
		 	res = "Visa"
			}
		} else {
			res = "INVALID"
	}
	return res
}

func check(clean string) string {
	res := ""
	if checkA(clean) == "American Express" {
		res = "American Express"
	} else if checkM(clean) == "MasterCard" {
		res = "MasterCard"
	} else if checkV(clean) == "Visa" {
		res = "Visa"
	} else {
		res = "INVALID"
	}
	return res
}



func main() {
	
	i := input()
	c := clean(i)
	ch := check(clean(i))
	a := checksum(i)
	
	fmt.Printf("%s card: %s, %s", ch, c, a)
}