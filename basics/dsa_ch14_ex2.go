package main

import "fmt"

type Element interface{}

type Node struct {
	data     Element
	next     *Node
	previous *Node
}

type DoublyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

func (dl *DoublyLinkedList) prepend(n *Node) {
	after := dl.head
	dl.head = n

	if after == nil {
		dl.tail = n
	} else {
		after.previous = n
	}

	dl.head.next = after

	dl.length++
}

func (l DoublyLinkedList) printList() {
	toPrint := l.head

	for i := 1; i <= l.length; i++ {
		fmt.Println(toPrint.data)
		toPrint = toPrint.next
	}
}

func (dl DoublyLinkedList) printListReverse() {
	toPrint := dl.tail

	for i := 1; i <= dl.length; i++ {
		fmt.Println(toPrint.data)
		toPrint = toPrint.previous
	}
}

func main() {
	myList := DoublyLinkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}
	node4 := &Node{data: 27}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	fmt.Println(myList)
	myList.printList()
	myList.printListReverse()
}
