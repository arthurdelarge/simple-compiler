package data

import "testing"

func TestStack(t *testing.T) {
	stack := CreateStack[int]()

	stack.Push(1)
	stack.Push(2)

	element, ok := stack.Pop()
	if !ok {
		t.Errorf("Expected to pop an element")
	}

	if element != 2 {
		t.Errorf("Expected to pop 2, got %d", element)
	}

	element, ok = stack.Pop()
	if !ok {
		t.Errorf("Expected to pop an element")
	}

	if element != 1 {
		t.Errorf("Expected to pop 1, got %d", element)
	}

	_, ok = stack.Pop()
	if ok {
		t.Errorf("Expected not to pop an element")
	}
}
