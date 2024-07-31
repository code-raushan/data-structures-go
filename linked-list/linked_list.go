package main

import "fmt"

type node struct {
	val int
	next *node
}

type linkedList struct {
	head *node
	length int
}

func (l *linkedList) prepend (n *node) {
	second := l.head
	l.head = n
	l.head.next = second

	l.length ++
}

func (l linkedList) print() {
	head := l.head
	for l.length != 0 {
		fmt.Print(head.val,  " ")
		head = head.next
		l.length --
	}

	fmt.Println()
}

func (l *linkedList) deleteNode(val int) {
	if l.length ==0 {
		return 
	}

	if (l.head.val == val ) {
		l.head = l.head.next
		l.length --

		return 
	}

	secondToTargetNode := l.head
	for secondToTargetNode.next.val != val {
		secondToTargetNode = secondToTargetNode.next
	}

	secondToTargetNode.next = secondToTargetNode.next.next
	l.length--
}

func main(){
	llist := &linkedList{}

	val1 := &node{val: 5}
	val2 := &node{val: 7}
	val3 := &node{val: 9}

	llist.prepend(val1)
	llist.prepend(val2)
	llist.prepend(val3)

	llist.print()

	llist.deleteNode(7)

	llist.print()
}