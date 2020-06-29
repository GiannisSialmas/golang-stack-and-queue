# General
This was an interesting task. Even though i was aware of stack and queue in theory, it was the first time i implemented something like that in a programming language.

In both the stack and queue implementation, we define a custom struct type Employee which are the items that are going to be added to both data structures.

The backbone of both data structures implementation is a slice which holds Employee struck pointers.

Both data structures have a main function in which 2 employess are added, and then one is retrieved to showcase the LIFO or FIFO.

# Stack Data Structure
Stack is a data structure which follows a particular order in which the operations are performed. The order is LIFO(Last In First Out).

As i described in the general section of the readme, the backbone of the imlementation is a slice which holds Employee struck pointers.

We wrap the slice in a custom Stack type for semantic reasons. We can get a new Stack using the NewStack function which returns a pointer to a new Stack.

We implement the following methods for the Stack type:
- Push which appends an employee to the end of the slice
- Peek which returns the last item of the slice(i.e the last item added chronologically) without removing it fror the stack, or nil if the stack is empty
- Pop which returns the last item of the slice(i.e the last item added chronologically) and also removing it fror the stack, or nil if the stack is empty

# Queue Data Structure
Queue is a data structure which follows a particular order in which the operations are performed. The order is FIFO(First In First Out).

As i described in the general section of the readme, the backbone of the imlementation is a slice which holds Employee struck pointers.

We wrap the slice in a custom Queue type for semantic reasons. We can get a new Queue using the NewQueue function which returns a pointer to a new Queue.

We implement the following methods for the Queue type:
- Enqueue which appends an Employee to the end of the slice
- Dequeue which returns the first item in the slice(i.e the first item added chronologically) and also removeing it from the queue, or nil if the queue is empty

# Bonus 
For this particular use case, the implementation footprint was pretty small so i would not include a package. Although if i did include one, it would be `container/list`