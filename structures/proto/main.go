package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"
)

func main() {
	person := &Person{
		Firstname: "John",
		Lastname:  "Doe",
	}

	data, err := proto.Marshal(person)

	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	fmt.Println(data)

	err = ioutil.WriteFile("person.data", data, 0644)
	if err != nil {
		log.Fatal("there was an error writing the file: ", err)
	}

	person2 := &Person{}
	err = proto.Unmarshal(data, person2)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	}
	fmt.Println(person2.Firstname)
	fmt.Println(person2.Lastname)
}

