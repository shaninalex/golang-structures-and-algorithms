package main

import (
	"github.com/shaninalex/golang-structures-and-algorithms/structures/linkedlist"
)

func main() {
	list := &linkedlist.LinkedList{}

	list.Insert(13)
	list.Insert(2)

	list.Traverse()

}
