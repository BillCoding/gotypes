package arraylist

import (
	"fmt"
	"github.com/billcoding/gotypes"
	"testing"
)

//Test simple
func TestSimple(t *testing.T) {
	//Get new ArrayList
	arrayList := New()

	//Add element
	arrayList.Add(10)

	//Add element
	arrayList.Add(100)

	//Add element
	arrayList.Add(1000)

	//Get size
	fmt.Println(arrayList.Size())

	//Remove index `0`
	fmt.Println(arrayList.Remove(0))

	//Get index `0`
	fmt.Println(arrayList.Get(0))

	//Clear elements
	arrayList.Clear()

	//Get size
	fmt.Println(arrayList.Size())
}

//Test function
func TestFunction(t *testing.T) {
	arrayList := New()
	arrayList.Add("100").Add("200").Add("300")

	//ForEach function
	arrayList.ForEach(func(e gotypes.E) {
		fmt.Println(e)
	})

	//Reduce function
	fmt.Println(arrayList.Reduce("", func(val gotypes.E, e gotypes.E) gotypes.E {
		return val.(string) + e.(string)
	}))
	//output: `100200300`

	//MatchAny function
	fmt.Println(arrayList.MatchAny(func(e gotypes.E) bool {
		return e == "100"
	}))
	//output: `true`

	//MatchNone function
	fmt.Println(arrayList.MatchNone(func(e gotypes.E) bool {
		return e == "1000"
	}))
	//output: `true`

	//MatchAll function
	fmt.Println(arrayList.MatchAll(func(e gotypes.E) bool {
		return len(e.(string)) == 3
	}))
	//output: `true`

	//FindAny function
	fmt.Println(arrayList.FindAny(func(e gotypes.E) bool {
		return len(e.(string)) == 3
	}))
	//output: `100`

	//Count function
	fmt.Println(arrayList.Count(func(e gotypes.E) bool {
		return len(e.(string)) == 3
	}))
	//output: `3`

}
