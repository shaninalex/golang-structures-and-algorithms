package stack

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestPush(t *testing.T) {
	s := NewStack(3)

	// Push three values onto the stack
	err := s.Push(Node{Value: 1})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = s.Push(Node{Value: 2})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = s.Push(Node{Value: 3})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Try pushing one more value, which should result in an error
	err = s.Push(Node{Value: 4})
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "stack overflow" {
		t.Errorf("Expected error message 'stack overflow', but got '%v'", err.Error())
	}
}

func TestPop(t *testing.T) {
	s := NewStack(3)

	// Push two values onto the stack
	s.Push(Node{Value: 1})
	s.Push(Node{Value: 2})

	// Pop the top value from the stack
	node, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the popped value is correct
	if node.Value != 2 {
		t.Errorf("Expected popped value to be 2, but got %d", node.Value)
	}

	// Pop the top value again
	node, err = s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the popped value is correct
	if node.Value != 1 {
		t.Errorf("Expected popped value to be 1, but got %d", node.Value)
	}

	// Try popping from an empty stack, which should result in an error
	_, err = s.Pop()
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "stack underflow" {
		t.Errorf("Expected error message 'stack underflow', but got '%v'", err.Error())
	}
}

func TestPeek(t *testing.T) {
	s := NewStack(3)

	// Push a value onto the stack
	s.Push(Node{Value: 1})

	// Peek at the top value
	node, err := s.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the peeked value is correct
	if node.Value != 1 {
		t.Errorf("Expected peeked value to be 1, but got %d", node.Value)
	}

	// Pop the value from the stack
	_, err = s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Try peeking from an empty stack, which should result in an error
	_, err = s.Peek()
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != "stack underflow" {
		t.Errorf("Expected error message 'stack underflow', but got '%v'", err.Error())
	}
}

func TestIsFull(t *testing.T) {
	s := NewStack(3)

	// Check if the stack is initially not full
	if s.IsFull() {
		t.Error("Expected stack to be not full, but it is")
	}

	// Push three values onto the stack
	s.Push(Node{Value: 1})
	s.Push(Node{Value: 2})
	s.Push(Node{Value: 3})

	// Check if the stack is full after pushing three values
	if !s.IsFull() {
		t.Error("Expected stack to be full, but it is not")
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack(3)

	// Check if the stack is initially empty
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty, but it is not")
	}

	// Push a value onto the stack
	s.Push(Node{Value: 1})

	// Check if the stack is not empty after pushing a value
	if s.IsEmpty() {
		t.Error("Expected stack to be not empty, but it is")
	}

	// Pop the value from the stack
	_, err := s.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Check if the stack is empty after popping the value
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty, but it is not")
	}
}

func TestSize(t *testing.T) {
	s := NewStack(3)

	// Check if the size of the stack is initially 0
	if s.Size() != 0 {
		t.Errorf("Expected stack size to be 0, but got %d", s.Size())
	}

	// Push three values onto the stack
	s.Push(Node{Value: 1})
	s.Push(Node{Value: 2})
	s.Push(Node{Value: 3})

	// Check if the size of the stack is 3 after pushing three values
	if s.Size() != 3 {
		t.Errorf("Expected stack size to be 3, but got %d", s.Size())
	}

	// Pop two values from the stack
	s.Pop()
	s.Pop()

	// Check if the size of the stack is 1 after popping two values
	if s.Size() != 1 {
		t.Errorf("Expected stack size to be 1, but got %d", s.Size())
	}
}

func TestClear(t *testing.T) {
	s := NewStack(3)

	// Push three values onto the stack
	s.Push(Node{Value: 1})
	s.Push(Node{Value: 2})
	s.Push(Node{Value: 3})

	// Clear the stack
	s.Clear()

	// Check if the stack is empty after clearing it
	if !s.IsEmpty() {
		t.Error("Expected stack to be empty, but it is not")
	}

	// Check if the capacity of the stack is still the same
	if s.Capacity != 3 {
		t.Errorf("Expected stack capacity to be 3, but got %d", s.Capacity)
	}
}

func TestPrint(t *testing.T) {
	s := NewStack(3)

	// Push three values onto the stack
	s.Push(Node{Value: 1})
	s.Push(Node{Value: 2})
	s.Push(Node{Value: 3})

	// Redirect stdout to capture the printed output
	output := captureOutput(func() {
		s.Print()
	})

	// Check if the printed output matches the expected output
	expectedOutput := "[1,2,3]\n"
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
