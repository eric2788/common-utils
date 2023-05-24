package array


func IndexOf[T comparable](arr []T, item T) int {
	for i, v := range arr {
		if item == v {
			return i
		}
	}
	return -1
}

func Contains[T comparable](arr []T, item T) bool {
	return IndexOf(arr, item) > -1
}

func Remove[T comparable](arr []T, item T) []T {
	index := IndexOf(arr, item)
	if index == -1 {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

func RemoveAt[T comparable](arr []T, index int) []T {
	if index < 0 || index >= len(arr) {
		return arr
	}
	return append(arr[:index], arr[index+1:]...)
}

func RemoveAll[T comparable](arr []T, items ...T) []T {
	for _, item := range items {
		arr = Remove(arr, item)
	}
	return arr
}

func RemoveIf[T comparable](arr []T, predicate func(T) bool) []T {
	for i := 0; i < len(arr); i++ {
		if predicate(arr[i]) {
			arr = RemoveAt(arr, i)
			i--
		}
	}
	return arr
}

func AddDistinct[T comparable](arr []T, item T) []T {
	if Contains(arr, item) {
		return arr
	}
	return append(arr, item)
}

func AddAllDistinct[T comparable](arr []T, items ...T) []T {
	for _, item := range items {
		arr = AddDistinct(arr, item)
	}
	return arr
}

