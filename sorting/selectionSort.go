package sorting

func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		ith := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[ith] {
				ith = j
			}
		}
		array[i], array[ith] = array[ith], array[i]
	}
	return array
}
