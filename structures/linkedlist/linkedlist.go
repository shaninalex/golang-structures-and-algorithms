// Single directional linked list
package linkedlist

import "fmt"

type Node struct {
	Value int64
	Next  *Node
}

type LinkedList struct {
	Head *Node
	// TODO: length
}

func (l *LinkedList) Insert(val int64) {
	node := &Node{Value: val}
	if l.Head == nil {
		l.Head = node
		return
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = node
		return
	}
}

func (l *LinkedList) Traverse() {
	p := l.Head
	for p != nil {
		fmt.Println(p.Value)
		p = p.Next
	}

}
