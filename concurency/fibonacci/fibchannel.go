package main

import (
	"fmt"
)

//fibonacci generator sending results to channel and close the channel at n numbers
func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {

	//make a channel
	ch := make(chan int, 10)
	//go routine execute the generator for 10 entries
	go fibonacci(cap(ch), ch)
	
	//iterate the channel and print each result
	for i := range ch{
		fmt.Println(i)
	}

}



