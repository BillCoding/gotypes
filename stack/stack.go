package stack

import (
	"github.com/billcoding/gotypes"
	"sync"
)

//Define eachFunc type
type eachFunc func(e gotypes.E)

//Define mapFunc type
type mapFunc func(e gotypes.E) gotypes.E

//Define matchFunc type
type matchFunc func(e gotypes.E) bool

//Define reduceFunc type
type reduceFunc func(val gotypes.E, e gotypes.E) gotypes.E

//Define Stack struct
type Stack struct {
	//A sync locker
	mutex sync.RWMutex
	//elements
	elements []gotypes.E
}

//New empty Stack
func New() *Stack {
	return &Stack{
		elements: make([]gotypes.E, 0),
		mutex:    sync.RWMutex{},
	}
}

//Clear stack
func (s *Stack) Clear() {
	s.elements = make([]gotypes.E, 0)
}

//Return stack size
func (s *Stack) Size() int {
	return len(s.elements)
}

//Push element
func (s *Stack) Push(e gotypes.E) *Stack {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.elements = append(s.elements, e)
	return s
}

//Pop element
func (s *Stack) Pop() gotypes.E {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	var e gotypes.E
	size := s.Size()
	if size > 0 {
		e = s.elements[size-1]
		es := make([]gotypes.E, 0)
		es = append(es, s.elements[:size-1]...)
		s.elements = es
	}
	return e
}

//ForEach function
func (s *Stack) ForEach(eachFunc eachFunc) {
	for i := range s.elements {
		eachFunc(s.elements[s.Size()-i-1])
	}
}

//Map function
func (s *Stack) Map(mapFunc mapFunc) *Stack {
	ns := New()
	s.ForEach(func(e gotypes.E) {
		ns.Push(mapFunc(e))
	})
	return ns
}

//Filter function
func (s *Stack) Filter(matchFunc matchFunc) *Stack {
	ns := New()
	s.ForEach(func(e gotypes.E) {
		if matchFunc(e) {
			ns.Push(e)
		}
	})
	return ns
}

//Reduce function
func (s *Stack) Reduce(val gotypes.E, reduceFunc reduceFunc) gotypes.E {
	s.ForEach(func(e gotypes.E) {
		val = reduceFunc(val, e)
	})
	return val
}

//MatchAny function
func (s *Stack) MatchAny(matchFunc matchFunc) bool {
	for _, e := range s.elements {
		if matchFunc(e) {
			return true
		}
	}
	return false
}

//MatchNone function
func (s *Stack) MatchNone(matchFunc matchFunc) bool {
	for _, e := range s.elements {
		if matchFunc(e) {
			continue
		}
		return true
	}
	return false
}

//MatchAll function
func (s *Stack) MatchAll(matchFunc matchFunc) bool {
	count := 0
	s.ForEach(func(e gotypes.E) {
		if matchFunc(e) {
			count++
		}
	})
	return count == s.Size()
}
