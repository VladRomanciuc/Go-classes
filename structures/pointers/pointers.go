package main


import (
	"fmt"
)


//Pointers hold the memory address of a value.

func main() {
	count := 0
	increment(&count)
	fmt.Println(count)
	increment(&count)
	fmt.Println(count)
	increment(&count)
	fmt.Println(count)
}

func increment(i *int) {
	*i++
}
