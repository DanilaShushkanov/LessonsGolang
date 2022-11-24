package main

import (
	"fmt"
)

// Добавить второй параметр ок для некоторых функций и попробовать ддженерики
func main() {
	lister := NewLinkedList()

	fmt.Printf("IsEmpty: %v\n", lister.IsEmpty())
	fmt.Printf("Length: %v\n", lister.GetLength())
	lister.Add("1")
	fmt.Printf("IsEmpty: %v\n", lister.IsEmpty())
	lister.Add("2")
	fmt.Printf("Length: %v\n", lister.GetLength())
	lister.AddAt("3", 2)
	fmt.Printf("Length: %v\n", lister.GetLength())
	v1, ok := lister.IndexOf("2")
	fmt.Printf("IndexOf('2'): %v, %v\n", v1, ok)
	v2, ok := lister.ElementAt(2)
	fmt.Printf("ElementAt('2'): %v, %v\n", v2, ok)
	v3, ok := lister.ElementAt(200)
	fmt.Printf("ElementAt('200'): %v, %v\n", v3, ok)
	v4, ok := lister.RemoveAt(2)
	fmt.Printf("RemoveAt(2): %v, %v\n", v4, ok)
	fmt.Println("Remove('1')")
	lister.Remove("1")
	fmt.Printf("Length: %v\n", lister.GetLength())
}

type LinkedLister interface {
	GetLength() int
	GetHead() *Node
	IsEmpty() bool
	IndexOf(string) (int, bool)
	Add(string)
	Remove(string)
	ElementAt(int) (string, bool)
	RemoveAt(int) (string, bool)
	AddAt(string, int) bool
}

func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

type LinkedList struct {
	length int
	head   *Node
}

type Node struct {
	element string
	next    *Node
}

func NewNode(element string) *Node {
	return &Node{
		element: element,
		next:    &Node{},
	}
}

func (ll *LinkedList) GetLength() int {
	return ll.length
}

func (ll *LinkedList) GetHead() *Node {
	return ll.head
}

func (ll *LinkedList) IsEmpty() bool {
	return ll.length == 0
}

func (ll *LinkedList) IndexOf(element string) (int, bool) {
	currentNode := ll.head
	counter := -1
	var ok bool
	for currentNode != nil {
		counter++
		if currentNode.element == element {
			ok = true
			break
		}
		currentNode = currentNode.next
	}

	return counter, ok
}

func (ll *LinkedList) Add(element string) {
	node := NewNode(element)
	if ll.head == nil {
		ll.head = node
	} else {
		currentNode := ll.head

		for currentNode.next != nil {
			currentNode = currentNode.next
		}

		currentNode.next = node
	}
	ll.length++
}

func (ll *LinkedList) Remove(element string) {
	currentNode := ll.head
	var previousNode = NewNode("")
	if currentNode.element == element {
		ll.head = currentNode.next
	} else {
		for currentNode.element != element {
			previousNode = currentNode
			currentNode = currentNode.next
		}
		previousNode.next = currentNode.next
	}

	ll.length--
}

func (ll *LinkedList) ElementAt(index int) (string, bool) {
	if index > ll.length || index < 0 {
		return "", false
	}
	currentNode := ll.head
	counter := 0

	for counter < index {
		counter++
		currentNode = currentNode.next
	}
	return currentNode.element, true
}

func (ll *LinkedList) RemoveAt(index int) (string, bool) {
	if index < 0 || index > ll.length {
		return "", false
	}

	currentNode := ll.head
	var previousNode = NewNode("")
	counter := 0
	if index == 0 {
		ll.head = currentNode.next
	} else {
		for counter < index {
			counter++
			previousNode = currentNode
			currentNode = currentNode.next
		}
		previousNode.next = currentNode.next
	}
	ll.length--

	return currentNode.element, true
}

func (ll *LinkedList) AddAt(element string, index int) bool {
	if index < 0 || index > ll.length {
		return false
	}

	node := NewNode(element)
	currentNode := ll.head
	var previousNode = NewNode("")
	counter := 0

	if index == 0 {
		node.next = currentNode
		ll.head = node
	} else {
		for counter < index {
			counter++
			previousNode = currentNode
			currentNode = currentNode.next
		}
		node.next = currentNode
		previousNode.next = node
	}
	ll.length++

	return true
}
