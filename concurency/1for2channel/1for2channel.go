package main

import (
	"time"
	"fmt"
)

//func sending messages to the channel
func smsReady(ch chan string) {
	time.Sleep(time.Millisecond)
	ch <- "hi friend"
	time.Sleep(time.Millisecond * 100)
	ch <- "What's going on?"
	time.Sleep(time.Second)
	ch <- "Will you make your appointment?"
	time.Sleep(time.Millisecond * 350)
	ch <- "Let's be friends"
	close(ch)
}

//func sending messages to the channel 2
func emailReady(ch chan string) {
	time.Sleep(time.Millisecond * 503)
	ch <- "Welcome to the business"
	time.Sleep(time.Millisecond * 43)
	ch <- "I'll pay you to be my friend"
	time.Sleep(time.Second)
	ch <- "How's the family?"
	time.Sleep(time.Millisecond * 440)
	ch <- "Want to go out tonight?"
	close(ch)
}

func main() {

	//make 2 channel for 2 func
	chEmails := make(chan string)
	chSms := make(chan string)

	//go routine execution
	go smsReady(chSms)
	go emailReady(chEmails)
	
	//Loop to pick the first channel that recieve a message
	for {
		//using select
		select {
			//sms case
	  		case i, ok := <- chSms:
				//if ok false return
				if !ok {
					return
			}
				fmt.Println("sending sms:", i)
	  		//email case
			case s, ok := <- chEmails:
				
				if !ok {
					return
			}	
				fmt.Println("sending email:", s)	
	
		}
	}
}


