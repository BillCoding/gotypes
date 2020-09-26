package linkedlist

import (
	"fmt"
	"testing"
)

func TestSimple(t *testing.T) {
	//Get new LinkedList
	linkedList := NewLinkedList()

	//Add element
	linkedList.Add(10)
	linkedList.Add(100)
	linkedList.Add(1000)
	linkedList.Add(10000)

	//Get head element and remove it
	fmt.Println(linkedList.Take().Value)

	//Get tail element and remove it
	fmt.Println(linkedList.Peek().Value)

	//Get LinkedList size
	fmt.Println(linkedList.Size())

	//Foreach function
	linkedList.ForEach(func(e LinkedElement) {
		fmt.Println(e.Value)
	})

}

//Test functions
func TestFunctions(t *testing.T) {
	//Get new LinkedList
	linkedList := NewLinkedList()

	//Add element
	linkedList.Add(10)
	linkedList.Add(100)
	linkedList.Add(1000)
	linkedList.Add(10000)

	//Get head element and remove it
	fmt.Println(linkedList.Take().Value)

	//Get tail element and remove it
	fmt.Println(linkedList.Peek().Value)

	//Get LinkedList size
	fmt.Println(linkedList.Size())

	//Foreach function
	linkedList.ForEach(func(e LinkedElement) {
		fmt.Println(e.Value)
	})

}
