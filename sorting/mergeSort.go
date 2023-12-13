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
