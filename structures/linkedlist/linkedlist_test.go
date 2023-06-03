package linkedlist_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/shaninalex/golang-structures-and-algorithms/structures/linkedlist"
)

func TestInsert(t *testing.T) {
	ll := linkedlist.LinkedList{}
	ll.Insert(5)

	// Check if the Head node has been properly set
	if ll.Head == nil {
		t.Error("Expected Head node to be set, but it is nil")
	}

	// Check if the Head node Value is correct
	if ll.Head.Value != 5 {
		t.Errorf("Expected Head node Value to be 5, got %d", ll.Head.Value)
	}

	ll.Insert(10)
	ll.Insert(15)

	// Check if the second node has been properly set
	if ll.Head.Next == nil {
		t.Error("Expected second node to be set, but it is nil")
	}

	// Check if the second node Value is correct
	if ll.Head.Next.Value != 10 {
		t.Errorf("Expected second node Value to be 10, got %d", ll.Head.Next.Value)
	}

	// Check if the third node has been properly set
	if ll.Head.Next.Next == nil {
		t.Error("Expected third node to be set, but it is nil")
	}

	// Check if the third node Value is correct
	if ll.Head.Next.Next.Value != 15 {
		t.Errorf("Expected third node Value to be 15, got %d", ll.Head.Next.Next.Value)
	}
}

func TestTraverse(t *testing.T) {
	ll := linkedlist.LinkedList{}
	ll.Insert(5)
	ll.Insert(10)
	ll.Insert(15)

	// Redirect stdout to capture the printed output
	output := captureOutput(func() {
		ll.Traverse()
	})

	// Check if the printed output matches the expected output
	expectedOutput := "5\n10\n15\n"
	if output != expectedOutput {
		t.Errorf("Expected output:\n%s\nGot:\n%s", expectedOutput, output)
	}
}

// Helper function to capture the output of a function
func captureOutput(fn func()) string {
	// Redirect stdout to a pipe
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	fn()

	// Reset stdout
	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	return buf.String()
}
