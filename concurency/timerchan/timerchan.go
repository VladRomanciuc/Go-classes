package main

import (
	"fmt"
	"time"
)

//func to stream message on channel
func streamTweets(ch chan struct{}) {
	t := time.Millisecond * 670
	
	for {
		ch <- struct{}{}
		time.Sleep(t)
		t *= 2
	}

}


func main() {
	//make channel
	ch := make(chan struct{})
	//start streaming
	go streamTweets(ch)

	//loop using select
	for {
		select {
			//print when recieving on channel
		case <-ch:
			fmt.Println("got a tweet")
			//if channel sleep for 5 seconds return to initial
		case <-time.After(5 * time.Second):
			fmt.Println("too long since last tweet, disconnecting")
			return
		}
	}
}




