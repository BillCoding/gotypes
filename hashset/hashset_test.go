package hashset

import (
	"fmt"
	"github.com/billcoding/gotypes"
	"testing"
)

//Test simple
func TestSimple(t *testing.T) {
	//Get new HashSet
	hashSet := New()

	//Set element into HashSet
	hashSet.Set("a")
	hashSet.Set("b")
	hashSet.Set("c")

	//Check element exists
	fmt.Println(hashSet.Contains("123"))
	fmt.Println(hashSet.Size())

}

//Test simple
func TestFunction(t *testing.T) {
	//Get new HashSet
	hashSet := New()

	//Set element into HashSet
	hashSet.Set("a1")
	hashSet.Set("b2")
	hashSet.Set("c3")

	//Check element exists
	fmt.Println(hashSet.Contains("123"))

	//Reduce function
	fmt.Println("Reduce : ", hashSet.Reduce("", func(val gotypes.E, e gotypes.E) gotypes.E {
		val = (val).(string) + (e).(string)
		return val
	}))

	//ForEach function
	hashSet.ForEach(func(e gotypes.E) {
		fmt.Println(e)
	})

}
