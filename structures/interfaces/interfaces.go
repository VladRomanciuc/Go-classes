package main

import (
	"fmt"
	"errors"
)

//function that mirror email and sms structs using the expense interface
func getExpenseReport(e expense) (string, error) {
	//assign email as expenses
	email, ok := e.(email)
	
	//check if it is true then access the address and generate cost
	if ok {
		address := email.toAddress
		cost := email.cost()
		return fmt.Sprintf("the email going to %v will cost: %v", address, cost), nil
	}
	//assign sms as expenses
	sms, ok := e.(sms)
	//check if it is true then access the phone number and generate cost
	if ok {
		phone := sms.toPhoneNumber
		cost := sms.cost()
		return fmt.Sprintf("the sms going to %v will cost: %v", phone, cost), nil
	}
	//if it is other type return unknow
		i := "expense has an unknown type"
		return fmt.Sprintln("error:"), errors.New(string(i))
	
}


//function for the cost of email
func (e email) cost() float64 {
	//using the email struct check the sub tier and return the cost repectively
	if e.authorSubscriptionTier == "free" {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

//function for the cost of sms
func (s sms) cost() float64 {
	//using the sms struct check the sub tier and return the cost repectively
	if s.authorSubscriptionTier == "free" {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

//Zero value function
func (i invalid) cost() float64 {
	return 0.0
}

//an interface for cost
type expense interface {
	cost() float64
}

//The structures for email, sms and invalid
type email struct {
	authorSubscriptionTier string
	body                   string
	toAddress              string
}

type sms struct {
	authorSubscriptionTier string
	body                   string
	toPhoneNumber          string
}

type invalid struct{}


//
func main() {

	//Slice of expenses
	expenses := []expense{
		email{
			authorSubscriptionTier: "pro",
			body:                   "hello there",
			toAddress:              "john@doe.com",
		},
		email{
			authorSubscriptionTier: "free",
			body:                   "This meeting could have been an email",
			toAddress:              "jane@doe.com",
		},
		email{
			authorSubscriptionTier: "free",
			body:                   "This meeting could have been an email",
			toAddress:              "elon@doe.com",
		},
		sms{
			authorSubscriptionTier: "free",
			body:                   "This meeting could have been an email",
			toPhoneNumber:          "+155555509832",
		},
		sms{
			authorSubscriptionTier: "free",
			body:                   "This meeting could have been an email",
			toPhoneNumber:          "+155555504444",
		},
		invalid{},
	}

	//Run the loop for expenses report
	for _, e := range expenses {
		rep, err := getExpenseReport(e)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		fmt.Println("report:", rep)
	}
}
