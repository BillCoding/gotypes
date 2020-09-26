package hashset

import (
	"github.com/billcoding/gotypes"
	"sync"
)

//Define HashSet struct
type HashSet struct {
	//elements
	elements map[gotypes.E]gotypes.E
	//a sync locker
	mutex sync.RWMutex
}

//New HashSet
func New() *HashSet {
	return &HashSet{
		elements: make(map[gotypes.E]gotypes.E),
		mutex:    sync.RWMutex{},
	}
}

//Clear HashSet
func (s *HashSet) Clear() {
	s.elements = make(map[gotypes.E]gotypes.E)
}

//Return HashSet size
func (s *HashSet) Size() int {
	return len(s.elements)
}

//Delete element
func (s *HashSet) Delete(e gotypes.E) *HashSet {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	delete(s.elements, e)
	return s
}

//Set element
func (s *HashSet) Set(e gotypes.E) *HashSet {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.elements[e] = 0
	return s
}

//Check contains
func (s *HashSet) Contains(e gotypes.E) bool {
	return s.elements[e] != nil
}

//Define eachFunc type
type eachFunc func(e gotypes.E)

//Define mapFunc type
type mapFunc func(e gotypes.E) gotypes.E

//Define matchFunc type
type matchFunc func(e gotypes.E) bool

//Define reduceFunc type
type reduceFunc func(val gotypes.E, e gotypes.E) gotypes.E

//ForEach function
func (s *HashSet) ForEach(eachFunc eachFunc) {
	for k := range s.elements {
		eachFunc(k)
	}
}

//Map function
func (s *HashSet) Map(mapFunc mapFunc) *HashSet {
	ns := New()
	s.ForEach(func(e gotypes.E) {
		ns.Set(mapFunc(e))
	})
	return ns
}

//Filter function
func (s *HashSet) Filter(matchFunc matchFunc) *HashSet {
	ns := New()
	s.ForEach(func(e gotypes.E) {
		if matchFunc(e) {
			ns.Set(e)
		}
	})
	return ns
}

//MatchAny function
func (s *HashSet) MatchAny(matchFunc matchFunc) bool {
	for k := range s.elements {
		if matchFunc(k) {
			return true
		}
	}
	return false
}

//MatchNone function
func (s *HashSet) MatchNone(matchFunc matchFunc) bool {
	for k := range s.elements {
		if matchFunc(k) {
			continue
		}
		return true
	}
	return false
}

//MatchAll function
func (s *HashSet) MatchAll(matchFunc matchFunc) bool {
	count := 0
	for k := range s.elements {
		if matchFunc(k) {
			count++
		}
	}
	return count == s.Size()
}

//Count function
func (s *HashSet) Count(matchFunc matchFunc) int {
	count := 0
	s.ForEach(func(e gotypes.E) {
		if matchFunc(e) {
			count++
		}
	})
	return count
}

//Count function
func (s *HashSet) Reduce(val gotypes.E, reduceFunc reduceFunc) gotypes.E {
	s.ForEach(func(e gotypes.E) {
		val = reduceFunc(val, e)
	})
	return val
}
