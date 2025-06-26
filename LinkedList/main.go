package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

/*
The concept of a linked list is a data structure that consists of nodes,
where each node contains data and a pointer to the next node in the sequence.

1. Push new node to list.
Common algorithm:
	a. If the list is empty, set the head to the new node.
	b. If the list is not empty, traverse to the end and append the new node.

2. Pop a node from the list.
Common algorithm:
	a. If the list is empty, return.
	b. If the head is the node to be removed, set the head to the next node.
	c. If the node to be removed is not the head, traverse the list to find it,
	   and change the next pointer of the previous node to skip the current node.
*/

// L1->L2->L3->L4->Nil
type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Push(data interface{}) {
	node := &Node{data: data}

	// If the list is empty, set the head to the new node
	if l.head == nil {
		l.head = node
		l.length++
		return
	}

	// If the list is not empty, traverse to the end and append the new node
	current := l.head
	for current.next != nil {
		current = current.next
	}

	// Push the new node to the end of the list
	current.next = node
	l.length++
}

func (l *LinkedList) Pop(data interface{}) {
	// If the list is empty, return
	if l.head == nil {
		return
	}

	// If the head is the node to be removed
	if l.head.data == data {
		l.head = l.head.next
		l.length--
		return
	}

	// If the node to be removed is not the head, in the middle or at the end
	prev := l.head
	curr := l.head.next

	for curr != nil {
		if curr.data == data {
			// Change the next pointer of the previous node to skip the current node
			prev.next = curr.next
			l.length--
			return
		}
		// Move to the next node
		prev = curr
		curr = curr.next
	}
}

func (l *LinkedList) CheckNode(data interface{}) *Node {
	curr := l.head
	for curr != nil {
		// If current node's data matches, return the current node
		if curr.data == data {
			return curr
		}
		// If current node's data matches, change the current node to the next one
		curr = curr.next
	}

	return nil
}

func (l *LinkedList) Update(oldData, newData interface{}) {
	// Check if the old data exists in the list
	oldNode := l.CheckNode(oldData)
	if oldNode != nil {
		// If it exists, update the data of the old node
		oldNode.data = newData
	}
}

func (l *LinkedList) Print() {
	if l.head == nil {
		fmt.Println("List is empty")
		return
	}

	curr := l.head
	fmt.Print("Head -> ")
	for curr != nil {
		fmt.Print(curr.data, " -> ")
		curr = curr.next
	}
	fmt.Println("Nil")
	fmt.Println("Length of the list: ", l.length)
}

func linkedListImplementation() {
	myList := &LinkedList{}

	fmt.Println("====START====")

	fmt.Println("====PUSH====")
	myList.Push("first")
	myList.Push(1)
	myList.Push("second")
	myList.Push(2)
	myList.Print()

	fmt.Println("====POP====")
	myList.Pop("first")
	myList.Pop("second")
	myList.Print()

	fmt.Println("====UPDATE====")
	myList.Update(1, "first")
	myList.Update(2, "second")
	myList.Print()

	fmt.Println("====END====")
}

func main() {
	linkedListImplementation()
}
