package main

import (
	"fmt"
)
//Declare a structure
type Person struct {
	name string
	lastName string
	age int
	passport int
}

//Basic work with structures
func (person Person) Print() {
	fmt.Println(person)
}

func (person Person) Age() {
	fmt.Println(person.age)
}

func (person Person) GetPassport() int {
	return person.passport
}

func main() {
	p := Person {"John","Smith",34,2345}

	p.Print()
	p.Age()
	fmt.Println(p.GetPassport())
}