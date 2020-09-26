package hashmap

import (
	"github.com/billcoding/gotypes"
	"github.com/billcoding/gotypes/arraylist"
	"github.com/billcoding/gotypes/hashset"
	"sync"
)

//Define matchFunc type
type matchFunc func(k gotypes.K, v gotypes.V) bool

//Define mapFunc type
type mapFunc func(k gotypes.K, v gotypes.V) (gotypes.K, gotypes.V)

//Define eachFunc type
type eachFunc func(k gotypes.K, v gotypes.V)

//Define reduceFunc
type reduceFunc func(val gotypes.E, k gotypes.K, v gotypes.V) gotypes.E

//Define HashMap struct
type HashMap struct {
	entries map[gotypes.K]gotypes.V
	mutex   sync.RWMutex
}

//Create new empty HashMap
func New() *HashMap {
	return &HashMap{
		entries: make(map[gotypes.K]gotypes.V),
		mutex:   sync.RWMutex{},
	}
}

//Clear HashMap
func (m *HashMap) Clear() {
	m.entries = make(map[gotypes.K]gotypes.V)
}

//return HashMap size
func (m *HashMap) Size() int {
	return len(m.entries)
}

//push entry
func (m *HashMap) Put(k gotypes.K, v gotypes.V) *HashMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.entries[k] = v
	return m
}

//delete key
func (m *HashMap) Delete(k gotypes.K) *HashMap {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	delete(m.entries, k)
	return m
}

//Get entry
func (m *HashMap) Get(k gotypes.K) gotypes.V {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	var v gotypes.V
	size := m.Size()
	if size > 0 {
		v = m.entries[k]
	}
	return v
}

//Check contains key
func (m *HashMap) Key(k gotypes.K) bool {
	return m.Get(k) != nil
}

//Check contains value
func (m *HashMap) Value(v gotypes.V) bool {
	m.mutex.RLock()
	defer m.mutex.RUnlock()
	size := m.Size()
	if size > 0 {
		for _, _v := range m.entries {
			if _v == v {
				return true
			}
		}
	}
	return false
}

//Get Keys HashSet
func (m *HashMap) Keys() *hashset.HashSet {
	hashSet := hashset.New()
	m.ForEach(func(k gotypes.K, v gotypes.V) {
		hashSet.Set(k)
	})
	return hashSet
}

//Get Values ArrayList
func (m *HashMap) Values() *arraylist.ArrayList {
	arrayList := arraylist.New()
	m.ForEach(func(k gotypes.K, v gotypes.V) {
		arrayList.Add(v)
	})
	return arrayList
}

//ForEach function
func (m *HashMap) ForEach(eachFunc eachFunc) {
	for k, v := range m.entries {
		eachFunc(k, v)
	}
}

//Map function
func (m *HashMap) Map(mapFunc mapFunc) *HashMap {
	nm := New()
	m.ForEach(func(k gotypes.K, v gotypes.V) {
		nm.Put(mapFunc(k, v))
	})
	return nm
}

//Filter function
func (m *HashMap) Filter(matchFunc matchFunc) *HashMap {
	nm := New()
	m.ForEach(func(k gotypes.K, v gotypes.V) {
		if matchFunc(k, v) {
			nm.Put(k, v)
		}
	})
	return nm
}

//matchAny function
func (m *HashMap) MatchAny(matchFunc matchFunc) bool {
	for k, v := range m.entries {
		if matchFunc(k, v) {
			return true
		}
	}
	return false
}

//matchNone function
func (m *HashMap) MatchNone(matchFunc matchFunc) bool {
	for k, v := range m.entries {
		if matchFunc(k, v) {
			continue
		}
		return true
	}
	return false
}

//MatchAll function
func (m *HashMap) MatchAll(matchFunc matchFunc) bool {
	count := 0
	for k, v := range m.entries {
		if matchFunc(k, v) {
			count++
		}
	}
	return count == m.Size()
}

//Count function
func (m *HashMap) Count(matchFunc matchFunc) int {
	count := 0
	m.ForEach(func(k gotypes.K, v gotypes.V) {
		if matchFunc(k, v) {
			count++
		}
	})
	return count
}

//Count function
func (m *HashMap) Reduce(val gotypes.E, reduceFunc reduceFunc) gotypes.E {
	m.ForEach(func(k gotypes.K, v gotypes.V) {
		val = reduceFunc(val, k, v)
	})
	return val
}
