package stack

import (
	"fmt"
	"testing"
)

//Test simple
func TestSimple(t *testing.T) {
	//Get new Stack
	stack := New()

	//Push element
	stack.Push(123)
	stack.Push("222")
	stack.Push("3333")
	stack.Push("4444")
	stack.Push("5555")

	//Pop element
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())

	//Get Stack  size
	fmt.Println(stack.Size())

}
