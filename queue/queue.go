package main

import (
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
}

type Queue []*Employee

func NewQueue() *Queue {
	var s []*Employee
	return (*Queue)(&s)
}

func (queue *Queue) Dequeue() *Employee {
	if len(*queue) == 0 {
		return nil
	}
	employee := (*queue)[0]
	*queue = (*queue)[:len(*queue)-1]
	return employee
}

func (queue *Queue) Enqueue(employeeToAdd *Employee) {
	*queue = append(*queue, employeeToAdd)
}

func main() {
	fooQueue := NewQueue()

	dequeuedEmployeeFromEmptyQueue := fooQueue.Dequeue()
	if dequeuedEmployeeFromEmptyQueue == nil {
		fmt.Println("This should print as there is nothing in the queue when we dequeued")
	} else {
		fmt.Println("dequeued employee", dequeuedEmployeeFromEmptyQueue)
	}

	firstEmployeeToAdd := Employee{
		FirstName: "First",
		LastName:  "Employee",
	}

	fooQueue.Enqueue(&firstEmployeeToAdd)
	fmt.Println("Queue after adding employee", *fooQueue, "\n")

	secondEmployeeToAdd := Employee{
		FirstName: "Second",
		LastName:  "Employee",
	}

	fooQueue.Enqueue(&secondEmployeeToAdd)
	fmt.Println("Queue after adding employee", *fooQueue, "\n")

	dequeuedEmployee := fooQueue.Dequeue()
	if dequeuedEmployee == nil {
		fmt.Println("Peek error")
	}
	fmt.Println("Dequeued employee", dequeuedEmployee)
	fmt.Println("Queue after dequeued employee", *fooQueue, "\n")

	fmt.Println("Terminating...")

}
