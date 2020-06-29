package main

import (
	"fmt"
)

type Employee struct {
	FirstName string
	LastName  string
}

type Stack []*Employee

func NewStack() *Stack {
	var s []*Employee
	return (*Stack)(&s)
}

func (stack *Stack) Pop() *Employee {
	if len(*stack) == 0 {
		return nil
	}
	employee := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return employee
}

func (stack *Stack) Peek() *Employee {
	if len(*stack) == 0 {
		return nil
	}
	employee := (*stack)[len(*stack)-1]
	return employee
}

func (stack *Stack) Push(employeeToAdd *Employee) {
	*stack = append(*stack, employeeToAdd)
}

func main() {
	fooStack := NewStack()

	peekedEmployeeFromEmptyStack := fooStack.Peek()
	if peekedEmployeeFromEmptyStack == nil {
		fmt.Println("This should print as there is nothing in the stack when we peeked")
	} else {
		fmt.Println("Peeked employee", peekedEmployeeFromEmptyStack)
	}

	popedEmployeeFromEmptyStack := fooStack.Pop()
	if popedEmployeeFromEmptyStack == nil {
		fmt.Println("This should print as there is nothing in the stack when we poped\n")
	} else {
		fmt.Println("Poped employee", popedEmployeeFromEmptyStack)
	}

	firstEmployeeToAdd := Employee{
		FirstName: "First",
		LastName:  "Employee",
	}

	fooStack.Push(&firstEmployeeToAdd)
	fmt.Println("Stack after adding employee", *fooStack, "\n")

	secondEmployeeToAdd := Employee{
		FirstName: "Second",
		LastName:  "Employee",
	}

	fooStack.Push(&secondEmployeeToAdd)
	fmt.Println("Stack after adding employee", *fooStack, "\n")

	peekedEmployee := fooStack.Peek()
	if peekedEmployee == nil {
		fmt.Println("Peek error")
	}
	fmt.Println("Peeked employee", peekedEmployee)
	fmt.Println("Stack after peeked employee", *fooStack, "\n")

	retrievedEmployee := fooStack.Pop()
	if retrievedEmployee == nil {
		fmt.Println("Pop error:")
	}
	fmt.Println("Retrieved employee", retrievedEmployee)
	fmt.Println("Stack after retreived employee", *fooStack, "\n")

	fmt.Println("Terminating...")

}
