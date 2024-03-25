package sorting

func merge(left, right []int) []int {
	leftLength := len(left)
	rightLength := len(right)
	answer := make([]int, 0)
	leftIndex, rightIndex := 0, 0
	for leftIndex < leftLength && rightIndex < rightLength {
		if left[leftIndex] > right[rightIndex] {
			answer = append(answer, right[rightIndex])
			rightIndex++
		} else {
			answer = append(answer, left[leftIndex])
			leftIndex++
		}
	}
	if leftIndex < leftLength {
		answer = append(answer, left[leftIndex:]...)
	}
	if rightIndex < rightLength {
		answer = append(answer, right[rightIndex:]...)
	}
	return answer
}

func MergeSort(array []int) []int {
	length := len(array)
	if length == 1 {
		return array
	}
	middle := length / 2
	leftPart := MergeSort(array[:middle])
	rightPart := MergeSort(array[middle:])
	return merge(leftPart, rightPart)
}

func MergeSortNonRecursive(array []int) []int {
	step := 1
	for step < len(array) {
		left, mid := 0, step
		for left < len(array) {
			var temp []int
			i, j := left, mid
			for i < left+step && j < mid+step && j < len(array) {
				if array[i] < array[j] {
					temp = append(temp, array[i])
					i++
				} else {
					temp = append(temp, array[j])
					j++
				}
			}
			for i < left+step && i < len(array) {
				temp = append(temp, array[i])
				i++
			}
			for j < mid+step && j < len(array) {
				temp = append(temp, array[j])
				j++
			}
			if i == len(array) {
				copy(array[left:i], temp)
			} else {
				copy(array[left:j], temp)
			}
			left, mid = left+step*2, mid+step*2
		}
		if step > len(array)>>1 {
			break
		}
		step <<= 1
	}
	return array
}
