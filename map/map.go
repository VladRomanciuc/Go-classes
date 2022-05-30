package main

import (
	"fmt"
	"errors"
)
//Declare a struct for a user
type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

//Function to check if user is in the map...
func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	//return entry and a bool value (ok = true : exist, ok = false : nonex)
	user, ok := users[name]
	//the first check only with bool
	if !ok {
		return false, errors.New("user doesn't exist")
	}
	//the second check is with bool but inside the user struct
	if !user.scheduledForDeletion {
		return false, nil
	}
	//delete the user from map
	delete(users, name)
	return true, nil
}




func test(users map[string]user, name string) {
	fmt.Println("checking", name)
	deleted, err := deleteIfNecessary(users, name)
	if err != nil {
		fmt.Println(err)
		return
	}
	if deleted {
		fmt.Println("deleted:", name)
		return
	}
	fmt.Println("not deleted:", name)
}

func main() {
	users := map[string]user{
		"john": {
			name:                 "john",
			number:               18965554631,
			scheduledForDeletion: true,
			},
		"elon": {
			name:                 "elon",
			number:               19875556452,
			scheduledForDeletion: true,
			},
		"breanna": {
			name:                 "breanna",
			number:               98575554231,
			scheduledForDeletion: false,
			},
		"kade": {
			name:                 "kade",
			number:               10765557221,
			scheduledForDeletion: false,
			},
		}
	test(users, "john")
	test(users, "musk")
	test(users, "santa")
	test(users, "kade")
	fmt.Println(len(users))

}