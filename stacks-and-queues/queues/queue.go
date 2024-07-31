package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) enqueue(n int) {
	q.items = append(q.items, n)
}

func (q *Queue) dequeue() int {
	dequeuedItem := q.items[0]
	q.items = q.items[1:]

	return dequeuedItem
}

func main(){
	myQueue := Queue{}

	myQueue.enqueue(44)
	myQueue.enqueue(55)
	myQueue.enqueue(66)
	
	fmt.Println(myQueue)

	myQueue.dequeue()

	fmt.Println(myQueue)
}