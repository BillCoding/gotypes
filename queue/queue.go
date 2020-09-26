package queue

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

//Define Queue struct
type Queue struct {
	//A sync locker
	mutex sync.RWMutex
	//elements
	elements []gotypes.E
}

//New empty Queue
func New() *Queue {
	elements := make([]gotypes.E, 0)
	return &Queue{
		elements: elements,
		mutex:    sync.RWMutex{},
	}
}

//Clear queue
func (q *Queue) Clear() {
	q.elements = make([]gotypes.E, 0)
}

//Return queue size
func (q *Queue) Size() int {
	return len(q.elements)
}

//Push element
func (q *Queue) Push(e gotypes.E) *Queue {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.elements = append(q.elements, e)
	return q
}

//Pull first element
func (q *Queue) Pull() gotypes.E {
	q.mutex.RLock()
	defer q.mutex.RUnlock()
	var e gotypes.E
	size := q.Size()
	if size > 0 {
		e = q.elements[0]
		es := make([]gotypes.E, 0)
		es = append(es, q.elements[1:]...)
		q.elements = es
	}
	return e
}

//ForEach function
func (q *Queue) ForEach(eachFunc eachFunc) {
	for _, e := range q.elements {
		eachFunc(e)
	}
}

//Map function
func (q *Queue) Map(mapFunc mapFunc) *Queue {
	nq := New()
	q.ForEach(func(e gotypes.E) {
		nq.Push(mapFunc(e))
	})
	return nq
}

//filter function
func (q *Queue) Filter(matchFunc matchFunc) *Queue {
	nq := New()
	q.ForEach(func(e gotypes.E) {
		if matchFunc(e) {
			nq.Push(e)
		}
	})
	return nq
}

//MatchAny function
func (q *Queue) MatchAny(matchFunc matchFunc) bool {
	for _, e := range q.elements {
		if matchFunc(e) {
			return true
		}
	}
	return false
}

//MatchNone function
func (q *Queue) MatchNone(matchFunc matchFunc) bool {
	for _, e := range q.elements {
		if matchFunc(e) {
			continue
		}
		return true
	}
	return false
}

//MatchAll function
func (q *Queue) MatchAll(matchFunc matchFunc) bool {
	count := 0
	for _, e := range q.elements {
		if matchFunc(e) {
			count++
		}
	}
	return count == q.Size()
}

//Count function
func (q *Queue) Count(matchFunc matchFunc) int {
	count := 0
	q.ForEach(func(e gotypes.E) {
		if matchFunc(e) {
			count++
		}
	})
	return count
}

//Reduce function
func (q *Queue) Reduce(val gotypes.E, reduceFunc reduceFunc) gotypes.E {
	q.ForEach(func(e gotypes.E) {
		val = reduceFunc(val, e)
	})
	return val
}
