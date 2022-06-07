package main

import (
	"fmt"
)

func marion(a int) {
	s := " "
	d := "#"

	for i:=1; i <= a; i++ {
		
		space := (a-i)
		
		for n := space; n>0; n--{
			fmt.Print(s)
		}
		for m := a; m>space; m--{
			fmt.Print(d)
		}

		fmt.Print("\n")
	}
	
}

func marion2(input int ) {
	
	s := " "
	h := "#"

	for line:=1; line <= input; line++ { //i = line
		space := (input-line)

		for n := space; n > 0; n--{
			fmt.Print(s)
		}
		for m := input; m > space; m--{
			fmt.Print(h)
		}
		fmt.Print(s)
		for m := input; m > space; m--{
			fmt.Print(h)
		}
		for n := space; n > 0; n--{
			fmt.Print(s)
		}

		fmt.Print("\n")

	}
}


func main(){

	var answer int
	
	for { 
		fmt.Println("Enter a number between 1 and 8:")
		fmt.Scanln(&answer)
		if answer >= 1 && answer <=8 {
			marion(answer)
			fmt.Print("\n-----------\n")
			marion2(answer)
		}
	}

}