package main

import "fmt"

type Element interface{}

type Node struct {
	data     Element
	nextNode *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.nextNode = second
	l.length++
}

func (l LinkedList) printList() {
	toPrint := l.head

	for i := 1; i <= l.length; i++ {
		fmt.Println(toPrint.data)
		toPrint = toPrint.nextNode
	}
}

func (dl LinkedList) lastElement() Element {
	currentNode := dl.head

	for currentNode.nextNode != nil {
		currentNode = currentNode.nextNode
	}

	return currentNode.data
}

func main() {
	myList := LinkedList{}
	node1 := &Node{data: 48}
	node2 := &Node{data: 18}
	node3 := &Node{data: 16}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	fmt.Println(myList)
	myList.printList()
	fmt.Println(myList.lastElement())
}
