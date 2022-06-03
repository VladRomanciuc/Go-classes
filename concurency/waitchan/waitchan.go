package main

import "fmt"

func getDatabasesChannel() chan struct{} {
	//make a channel
	ch := make(chan struct{})
	//go routines send answer to channel one by one
	go func() {
		ch <- struct{}{}
		fmt.Println("first db online")
		ch <- struct{}{}
		fmt.Println("second db online")
		ch <- struct{}{}
		fmt.Println("third db online")
		ch <- struct{}{}
		fmt.Println("fourth db online")
	}()
	return ch
}


func main() {
	ch := getDatabasesChannel()
	//wait for first message
	<-ch
	//wait for second message
	<-ch
	//wait for 3rd message
	<-ch
	//wait for 4rd message
	<-ch
	//and after execute
	fmt.Println("mailio server ready")
}


