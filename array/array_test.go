package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestRemove(t *testing.T) {

	arr := []string{"a", "b", "c", "d", "e"}

	arr = Remove(arr, "b")

	assert.ElementsMatch(t, arr, []string{"a", "c", "d", "e"})


	arr2 := []int{1, 2, 3, 4, 5}

	arr2 = Remove(arr2, 2)

	assert.ElementsMatch(t, arr2, []int{1, 3, 4, 5})
}

func TestIndexOf(t *testing.T){

	arr := []string{"a", "b", "c", "d", "e"}

	assert.Equal(t, IndexOf(arr, "b"), 1)
	assert.Equal(t, IndexOf(arr, "f"), -1)
}

func TestContains(t *testing.T){
	
	arr := []string{"a", "b", "c", "d", "e"}

	assert.True(t, Contains(arr, "b"))
	assert.False(t, Contains(arr, "f"))
}

func TestRemoveAt(t *testing.T){
	
	arr := []string{"a", "b", "c", "d", "e"}

	arr = RemoveAt(arr, 2)

	assert.ElementsMatch(t, arr, []string{"a", "b", "d", "e"})
}

func TestRemoveAll(t *testing.T){
	
	arr := []string{"a", "b", "c", "d", "e"}

	arr = RemoveAll(arr, "b", "d")

	assert.ElementsMatch(t, arr, []string{"a", "c", "e"})
}


func TestRemoveIf(t *testing.T){
	
	arr := []string{"a", "b", "c", "d", "e"}

	arr = RemoveIf(arr, func(item string) bool {
		return item == "b" || item == "d"
	})

	assert.ElementsMatch(t, arr, []string{"a", "c", "e"})
}

func TestAddDistinct(t *testing.T){
	
	arr := []string{"a", "b", "c", "d", "e"}

	arr = AddDistinct(arr, "b")

	assert.ElementsMatch(t, arr, []string{"a", "b", "c", "d", "e"})

	arr = AddDistinct(arr, "f")

	assert.ElementsMatch(t, arr, []string{"a", "b", "c", "d", "e", "f"})
}

func TestAddAllDistinct(t *testing.T){
	
	arr := []string{"a", "b", "c", "d", "e"}

	arr = AddAllDistinct(arr, "b", "f")

	assert.ElementsMatch(t, arr, []string{"a", "b", "c", "d", "e", "f"})
}