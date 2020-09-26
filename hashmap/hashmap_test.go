package hashmap

import (
	"fmt"
	"github.com/billcoding/gotypes"
	"testing"
)

//Test simple
func TestSimple(t *testing.T) {
	//Get new HashMap
	hashMap := New()

	//Put an element
	hashMap.Put("abc", "a")
	hashMap.Put("111", "1")
	hashMap.Put("222", "2")

	//Check key exists
	fmt.Println(hashMap.Key("abc"))

	//Check Value exists
	fmt.Println(hashMap.Value("abc"))

	hashMap.Delete("111")

	//Get HashMap size
	fmt.Println(hashMap.Size())
}

//Test function
func TestFunction(t *testing.T) {
	//Get new HashMap
	hashMap := New()

	hashMap.Put("1111", 1111)
	hashMap.Put("33", 33)
	hashMap.Put("22", 22)
	hashMap.Put("22df", "22")
	hashMap.Put("55", 55)

	//Count
	fmt.Println(hashMap.Count(func(k gotypes.K, v gotypes.V) bool {
		return v == "22"
	}))

	//ForEach
	hashMap.ForEach(func(k gotypes.K, v gotypes.V) {
		fmt.Println(k, " = ", v)
	})

}
