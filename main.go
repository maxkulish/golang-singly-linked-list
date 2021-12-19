package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	first  *node
	last   *node
}

// Len return linkedList.length field which containes the number of nodes in the list
func (l *linkedList) Len() int {
	return l.length
}

// Display print all node.data fields of the linked list
func (l *linkedList) Display() {
	for l.first != nil {
		fmt.Printf("%v -> ", l.first.data)
		l.first = l.first.next
	}

	fmt.Println()
}

// InsertLast insert node to the end of the list
func (l *linkedList) InsertLast(n *node) {
	// list is empty
	if l.first == nil {
		l.first = n
		l.last = n
		l.length++
	} else {
		l.last.next = n
		l.last = n
		l.length++
	}
}

// Delete removes node from the linked list
func (l *linkedList) Delete(data int) {
	// first element should be removed
	if l.first.data == data {
		l.first = l.first.next
		l.length--
		return
	}

	// previous empty node
	var prev *node
	// current node for iteration
	curr := l.first
	for curr != nil && curr.data != data {
		prev = curr
		curr = curr.next
	}

	if curr == nil {
		fmt.Println("node not found")
		return
	}

	prev.next = curr.next
	l.length--
	fmt.Println("node deleted")
}

func main() {

	list := linkedList{}

	list.InsertLast(&node{data: 5})
	list.InsertLast(&node{data: 4})
	list.InsertLast(&node{data: 3})
	list.InsertLast(&node{data: 2})
	list.InsertLast(&node{data: 1})

	fmt.Println("Length = ", list.Len())
	list.Display()
	list.Delete(3)
	//fmt.Println("Length = ", list.Len())
	//list.Display()
}
