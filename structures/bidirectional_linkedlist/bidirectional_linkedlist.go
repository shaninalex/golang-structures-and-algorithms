package bidirectionallinkedlist

type Node struct {
	Value int64
	Next  *Node
	Prev  *Node
}

type BiLinkedList struct {
	Head *Node
	Tail *Node
}
