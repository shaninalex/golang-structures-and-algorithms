package main

import (
	"fmt"

	"github.com/shaninalex/golang-structures-and-algorithms/structures/stack"
)

func main() {
	// list := &linkedlist.LinkedList{}
	//
	// list.Insert(13)
	// list.Insert(2)
	//
	// list.Traverse()
	//
	st := stack.NewStack(10)
	node1 := stack.Node{Value: 1}
	st.Push(node1)
	node22 := stack.Node{Value: 21}
	st.Push(node22)
	node3 := stack.Node{Value: 333}
	st.Push(node3)

	st.Print()

	st.Pop()
	st.Print()

	fmt.Println("Size := ", st.Size())
}
