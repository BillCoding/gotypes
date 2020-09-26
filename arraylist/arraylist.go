package arraylist

import (
	"github.com/billcoding/gotypes"
	"math/rand"
	"sync"
)

//Define eachFunc
type eachFunc func(e gotypes.E)

//Define matchFunc
type matchFunc func(e gotypes.E) bool

//Define mapFunc
type mapFunc func(e gotypes.E) gotypes.E

//Define reduceFunc
type reduceFunc func(val gotypes.E, e gotypes.E) gotypes.E

//Define ArrayList struct
type ArrayList struct {
	//A RW sync locker
	mutex sync.RWMutex
	//elements
	elements []gotypes.E
}

//Create new empty ArrayList
func New() *ArrayList {
	elements := make([]gotypes.E, 0)
	return &ArrayList{
		elements: elements,
		mutex:    sync.RWMutex{},
	}
}

//Clear ArrayList
func (l *ArrayList) Clear() {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.elements = make([]gotypes.E, 0)
}

//Return ArrayList size
func (l *ArrayList) Size() int {
	return len(l.elements)
}

//Add element
func (l *ArrayList) Add(e gotypes.E) *ArrayList {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.elements = append(l.elements, e)
	return l
}

//Get index element
func (l *ArrayList) Get(index int) gotypes.E {
	l.mutex.RLock()
	defer l.mutex.RUnlock()
	var e gotypes.E
	size := l.Size()
	if size > 0 && index >= 0 && index <= size-1 {
		e = l.elements[index]
	}
	return e
}

//Remove index element
func (l *ArrayList) Remove(index int) gotypes.E {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	var e gotypes.E
	size := l.Size()
	if size > 0 && index >= 0 && index <= size-1 {
		e = l.elements[index]
		es := make([]gotypes.E, 0)
		if index > 0 {
			//copy left
			es = append(es, l.elements[:index]...)
		}
		if index < size-1 {
			//copy right
			es = append(es, l.elements[index+1:]...)
		}
		l.elements = es
	}
	return e
}

//ForEach function
func (l *ArrayList) ForEach(eachFunc eachFunc) {
	for _, e := range l.elements {
		eachFunc(e)
	}
}

//Map function
func (l *ArrayList) Map(mapFunc mapFunc) *ArrayList {
	ln := New()
	l.ForEach(func(e gotypes.E) {
		ln.Add(mapFunc(e))
	})
	return ln
}

//Filter function
func (l *ArrayList) Filter(matchFuns ...matchFunc) *ArrayList {
	if len(matchFuns) <= 0 {
		return l
	}
	nl := New()
	l.ForEach(func(e gotypes.E) {
		for _, mf := range matchFuns {
			if mf(e) {
				nl.Add(e)
			}
		}
	})
	return nl
}

//MatchAny function
func (l *ArrayList) MatchAny(matchFun matchFunc) bool {
	for _, e := range l.elements {
		if matchFun(e) {
			return true
		}
	}
	return false
}

//MatchNone function
func (l *ArrayList) MatchNone(matchFun matchFunc) bool {
	for _, e := range l.elements {
		if matchFun(e) {
			continue
		}
		return true
	}
	return false
}

//matchAll function
func (l *ArrayList) MatchAll(matchFun matchFunc) bool {
	count := 0
	for _, e := range l.elements {
		if matchFun(e) {
			count++
		}
	}
	return count == l.Size()
}

//FindAny function
func (l *ArrayList) FindAny(matchFuncs ...matchFunc) gotypes.E {
	if len(matchFuncs) <= 0 {
		return l.Get(rand.Intn(l.Size() - 1))
	}
	for _, e := range l.elements {
		count := 0
		for _, mf := range matchFuncs {
			if mf(e) {
				count++
			}
		}
		if count == len(matchFuncs) {
			return e
		}
	}
	return nil
}

//Count function
func (l *ArrayList) Count(matchFunc matchFunc) int {
	count := 0
	l.ForEach(func(e gotypes.E) {
		if matchFunc(e) {
			count++
		}
	})
	return count
}

//Reduce function
func (l *ArrayList) Reduce(val gotypes.E, reduceFunc reduceFunc) gotypes.E {
	l.ForEach(func(e gotypes.E) {
		val = reduceFunc(val, e)
	})
	return val
}
