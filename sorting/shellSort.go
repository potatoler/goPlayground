package sorting

func ShellSort(array []int) []int {
	length := len(array)
	step := length / 2
	for step > 0 {
		for i := step; i < length; i++ {
			j := i
			for j >= step && array[j] < array[j-step] {
				array[j], array[j-step] = array[j-step], array[j]
				j = j - step
			}
		}
		step /= 2
	}
	return array
}
