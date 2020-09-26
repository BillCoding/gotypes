package linkedlist

import (
	"github.com/billcoding/gotypes"
	"sync"
)

//Define eachFunc type
type eachFunc func(e LinkedElement)

//Define matchFunc type
type matchFunc func(e LinkedElement) bool

//Define reduceFunc type
type reduceFunc func(val gotypes.E, e LinkedElement) gotypes.E

//Define mapFunc type
type mapFunc func(e LinkedElement) gotypes.E

//Define LinkedElement struct
type LinkedElement struct {
	Prev  gotypes.E
	Value gotypes.E
	Next  gotypes.E
}

//Define LinkedList struct
type LinkedList struct {
	//A sync locker
	mutex sync.RWMutex
	//LinkedElements
	elements []LinkedElement
}

//Create new LinkedList
func New() *LinkedList {
	elements := make([]LinkedElement, 0)
	return &LinkedList{
		elements: elements,
		mutex:    sync.RWMutex{},
	}
}

//Clear List
func (l *LinkedList) Clear() {
	l.elements = make([]LinkedElement, 0)
}

//Return List size
func (l *LinkedList) Size() int {
	return len(l.elements)
}

//Add element into tail
func (l *LinkedList) Add(e gotypes.E) *LinkedList {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	var prevLinkedElement LinkedElement
	if l.Size() > 0 {
		prevLinkedElement = l.elements[l.Size()-1]
	}
	prev := prevLinkedElement.Value
	current := LinkedElement{
		Prev:  prev,
		Value: e,
		Next:  nil,
	}
	prevLinkedElement.Next = current.Value
	l.elements = append(l.elements, current)
	return l
}

//Return first element
func (l *LinkedList) Head() LinkedElement {
	var e LinkedElement
	l.mutex.Lock()
	defer l.mutex.Unlock()
	if l.Size() > 0 {
		e = l.elements[0]
	}
	return e
}

//Return last element
func (l *LinkedList) Tail() LinkedElement {
	var e LinkedElement
	l.mutex.Lock()
	defer l.mutex.Unlock()
	size := l.Size()
	if size > 0 {
		e = l.elements[size-1]
	}
	return e
}

//Return head and remove it
func (l *LinkedList) Take() LinkedElement {
	var e LinkedElement
	l.mutex.Lock()
	defer l.mutex.Unlock()
	size := l.Size()
	if size > 0 {
		e = l.elements[0]
		es := make([]LinkedElement, 0)
		es = append(es, l.elements[1:]...)
		l.elements = es
	}
	return e
}

//Return tail and remove it
func (l *LinkedList) Peek() LinkedElement {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	var e LinkedElement
	size := l.Size()
	if size > 0 {
		e = l.elements[size-1]
		es := make([]LinkedElement, 0)
		es = append(es, l.elements[:size-1]...)
		l.elements = es
	}
	return e
}

//ForEach function
func (l *LinkedList) ForEach(eachFunc eachFunc) {
	for _, e := range l.elements {
		eachFunc(e)
	}
}

//Map function
func (l *LinkedList) Map(mapFunc mapFunc) *LinkedList {
	ll := New()
	l.ForEach(func(e LinkedElement) {
		ll.Add(mapFunc(e))
	})
	return ll
}

//Filter function
func (l *LinkedList) Filter(matchFunc matchFunc) *LinkedList {
	ll := New()
	l.ForEach(func(e LinkedElement) {
		if matchFunc(e) {
			ll.Add(e.Value)
		}
	})
	return ll
}

//Reduce function
func (l *LinkedList) Reduce(val gotypes.E, reduceFunc reduceFunc) gotypes.E {
	firstPtr := val
	l.ForEach(func(e LinkedElement) {
		firstPtr = reduceFunc(firstPtr, e)
	})
	return firstPtr
}

//Count function
func (l *LinkedList) Count(matchFunc matchFunc) int {
	count := 0
	l.ForEach(func(e LinkedElement) {
		if matchFunc(e) {
			count++
		}
	})
	return count
}

//MatchAny function
func (l *LinkedList) MatchAny(matchFunc matchFunc) bool {
	for _, e := range l.elements {
		if matchFunc(e) {
			return true
		}
	}
	return false
}

//MatchNone function
func (l *LinkedList) MatchNone(matchFunc matchFunc) bool {
	for _, e := range l.elements {
		if matchFunc(e) {
			continue
		}
		return true
	}
	return false
}

//MatchAll function
func (l *LinkedList) MatchAll(matchFunc matchFunc) bool {
	count := 0
	for _, e := range l.elements {
		if matchFunc(e) {
			count++
		}
	}
	return count == l.Size()
}
