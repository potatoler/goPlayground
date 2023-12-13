package sorting

func partition(array []int, low, high int) int {
	pivot := low
	pointer := pivot + 1
	for i := pointer; i <= high; i++ {
		if array[i] < array[pivot] {
			array[i], array[pointer] = array[pointer], array[i]
			pointer += 1
		}
	}
	pointer -= 1
	array[pivot], array[pointer] = array[pointer], array[pivot]
	return pointer
}

func quickSort(array []int, low, high int) []int {
	if low < high {
		partitionIndex := partition(array, low, high)
		quickSort(array, low, partitionIndex-1)
		quickSort(array, partitionIndex+1, high)
	}
	return array
}

func quickSortGoStyle(array []int) {
	length := len(array)
	if length < 2 {
		return
	}
	head, pointer := 0, length-1
	pivot := array[head]
	for head < pointer {
		if array[head+1] > pivot {
			array[head+1], array[pointer] = array[pointer], array[head+1]
			pointer--
		} else if array[head+1] < pivot {
			array[head], array[head+1] = array[head+1], array[head]
			head++
		} else {
			head++
		}
	}
	quickSortGoStyle(array[:head])
	quickSortGoStyle(array[head+1:])
}

func QuickSort(array []int) []int {
	quickSortGoStyle(array)
	return array
}
