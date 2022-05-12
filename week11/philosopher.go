/*

Implement the dining philosopher’s problem with the following constraints/modifications.

There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

The host allows no more than 2 philosophers to eat concurrently.

Each philosopher is numbered, 1 through 5.

When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself,
where <number> is the number of the philosopher.

When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, 
where <number> is the number of the philosopher.

*/

package main

import (
	"fmt"
	"sync"
	"time"
	)


	//Declare the chostick structure
type Chopstick struct {
	sync.Mutex

	}

	//Declare the philosopher structure
type Philosopher struct {
	
	number, count      int

	leftChopstick, rightChopstick *Chopstick

	}

	//Turn on the sync
var on sync.Once

	//Function to print the start of the program
func start() {
	fmt.Println("\n-----------------------------")
	fmt.Println("Wake up one of the philosophers")
	fmt.Println("\n-----------------------------")

	}

	//Standart function for channels
func host(ch chan int) {
	ch <- 1
	ch <- 2
	<-ch
	}

	//Function for philosophers to eat chopsticks blocking between each adjacent pair of philosophers
func (philosopher Philosopher) eat(ch chan int) {

	on.Do(start)
	
	for i := 0; i < 3; i++ {
	
		<-ch
	
			fmt.Printf("Philosopher %v is starving and starts eating \n", philosopher.number)
	
			philosopher.leftChopstick.Lock()
			philosopher.rightChopstick.Lock()
	
			fmt.Printf("\nFinally, philosopher %v is full and finish eating.\n", philosopher.number)
			fmt.Println("\n--------------------------------------------------")
	
			philosopher.rightChopstick.Unlock()
			philosopher.leftChopstick.Unlock()
	
		ch <- i

		}

	}

//**************************************************************************************************

func main() {

	//Declare channels variable
	ch := make(chan int, 2)

	//Create the chopstick slice (15 items that allow philosophers to eat 3 times each)
	chopsticks := make([]*Chopstick, 15)
	for i := 0; i < 15; i++ {
		chopsticks[i] = new(Chopstick)
		}

	//Create the philosophers slice (5 of them)
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i + 1, 0, chopsticks[i], chopsticks[(i+1)%5]}
	}

	//Start communication
	go host(ch)

	//Start the loop for eating chopsticks with 5 philosophers
	for i := 0; i < 5; i++ {
		go philosophers[i].eat(ch)
	}

	//Give some spare time
	time.Sleep(1 * time.Second)
}

