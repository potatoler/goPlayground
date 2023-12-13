package sorting

func heapify(array []int, low, high int) []int {
	parent := low
	child := parent*2 + 1
	for child <= high {
		if child+1 <= high && array[child] < array[child+1] {
			child += 1
		}
		if array[parent] >= array[child] {
			return array
		} else {
			array[parent], array[child] = array[child], array[parent]
			parent = child
			child = parent*2 + 1
		}
	}
	return array
}

func HeapSort(array []int) []int {
	length := len(array)
	for i := (length - 2) / 2; i >= 0; i-- {
		array = heapify(array, i, length-1)
	}
	for i := length - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]
		array = heapify(array, 0, i-1)
	}
	return array
}
