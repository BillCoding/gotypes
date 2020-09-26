package queue

import (
	"fmt"
	"github.com/billcoding/gotypes"
	"testing"
)

//Test simple
func TestSimple(t *testing.T) {
	//Get new Queue
	queue := New()

	//Push element into Queue
	queue.Push(777)
	queue.Push(111)
	queue.Push(222)

	//Get first element and remove it
	fmt.Println(queue.Pull())

	//Get Queue size
	fmt.Println(queue.Size())
}

//Test function
func TestFunction(t *testing.T) {
	//Get new Queue
	queue := New()

	//Push element into Queue
	queue.Push(777)
	queue.Push(111)
	queue.Push(222)

	//Foreach function
	queue.ForEach(func(e gotypes.E) {
		fmt.Println(e)
	})
}
