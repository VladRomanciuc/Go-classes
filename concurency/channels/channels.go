package main

import (
	"fmt"
	"time"
)


type email struct {
	body string
	date time.Time
}

//the function take the email..... and send the result to a channel
func isOld(e email, c chan bool) {
	if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
		c <- true
	}
	c <- false
}

//Func to send the emails
func sendEmails(c chan bool) {
	//slice of emails
	emails := []email{
		{
			body: "Are you going to make it?",
			date: time.Date(2019, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "I need a break",
			date: time.Date(2021, 0, 0, 0, 0, 0, 0, time.UTC),
		},
		{
			body: "What we're you thinking?",
			date: time.Date(2022, 0, 0, 0, 0, 0, 0, time.UTC),
		},
	}
	//Loop the parse the slice, check and recieve the answer to channel c, after close channel
	for _, email := range emails {
		isOld(email, c)
	}
	close(c)
}


func main() {
	//make the channel
	c := make(chan bool)

	//Send emails
	go sendEmails(c)

	//Recieving result to channel
	isOld := <-c
	fmt.Println("email 1 is old:", isOld)
	isOld = <-c
	fmt.Println("email 2 is old:", isOld)
	isOld = <-c
	fmt.Println("email 3 is old:", isOld)

}
