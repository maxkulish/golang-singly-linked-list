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
// will get Panic if pointer to the linkedList is used
func (l linkedList) Display() {
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

// Front return node.data field of the first node
func (l *linkedList) Front() (int, error) {
	// case when list if empty
	if l.first == nil {
		return 0, fmt.Errorf("can't find front value. List might be empty")
	}

	return l.first.data, nil
}

// Back return node.data field of the last node
func (l *linkedList) Back() (int, error) {
	// case when list is empty
	if l.first == nil {
		return 0, fmt.Errorf("can't find back value. List might be empty")
	}

	return l.last.data, nil
}

// Reverse reverses the linked list. The first element becomes last. The last becomes the first.
func (l *linkedList) Reverse() {

	// case when list is empty
	if l.first == nil {
		return
	}

	curr := l.first
	l.last = l.first

	var prev *node
	for curr != nil {
		temp := curr.next
		curr.next = prev
		prev = curr
		curr = temp
	}
	l.first = prev
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
	fmt.Println("Length = ", list.Len())
	list.Display()

	first, _ := list.Front()
	fmt.Println("First: ", first)

	last, _ := list.Back()
	fmt.Println("Last: ", last)

	list.Reverse()
	list.Display()
}
