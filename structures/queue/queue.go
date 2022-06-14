package main

import "fmt"

type queue []int

// Returns true if the queue is empty
func (q queue) empty() bool {
	return len(q) == 0
}

// Enqueue adds v to the queue
func (q *queue) enqueue(v int) {
*q = append(*q, v)
}

// dequeue removes an item from the queue and returns it
func (q *queue) dequeue() int {
if len(*q) > 0 {
		v := (*q)[0]
		*q = (*q)[1:]
		return v
	}
	return 0
}

func main() {
	myQueue := queue{}
	myQueue.enqueue(1)
	myQueue.enqueue(2)
	myQueue.enqueue(100)
	myQueue.enqueue(20)
	fmt.Println(myQueue.empty())
	fmt.Println(myQueue.dequeue())
	fmt.Println(myQueue.dequeue())
	fmt.Println(myQueue.dequeue())
	fmt.Println(myQueue.dequeue())
	fmt.Println(myQueue.empty())
}