package sorting

func LSDRadixSort(array []int) []int {
	pivot := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > pivot {
			pivot = array[i]
		}
	}
	base := 1
	for pivot/base > 0 {
		count := make([]int, 10)
		for i := 0; i < len(array); i++ {
			count[(array[i]/base)%10]++
		}
		for i := 1; i < 10; i++ {
			count[i] += count[i-1]
		}
		prototype := make([]int, len(array))
		for i := len(array) - 1; i >= 0; i-- {
			prototype[count[(array[i]/base)%10]-1] = array[i]
			count[(array[i]/base)%10]--
		}
		array = prototype
		base *= 10
	}
	return array
}

func msdPartlySort(array []int, base int) []int {
	if base == 0 || len(array) == 0 {
		return array
	}
	// fmt.Printf("base %d: ", base)
	count := make([]int, 10)
	for i := 0; i < len(array); i++ {
		count[(array[i]/base)%10]++
	}
	var digitBin [10]int
	for i := 0; i < 10; i++ {
		digitBin[i] = count[i]
	}
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	prototype := make([]int, len(array))
	answer := make([]int, 0)
	for i := len(array) - 1; i >= 0; i-- {
		prototype[count[(array[i]/base)%10]-1] = array[i]
		count[(array[i]/base)%10]--
	}
	// for i := 0; i < len(digitBin); i++ {
	// 	fmt.Printf("%d ", digitBin[i])
	// }
	// fmt.Printf("\n")
	po := 0
	for i := 0; i < 10; i++ {
		answer = append(answer, msdPartlySort(prototype[po:po+digitBin[i]], base/10)...)
		po += digitBin[i]
	}
	return answer
}

func MSDRadixSort(array []int) []int {
	pivot := array[0]
	for i := 1; i < len(array); i++ {
		if array[i] > pivot {
			pivot = array[i]
		}
	}
	base := 1
	for pivot/base > 0 {
		base *= 10
	}
	base /= 10
	array = msdPartlySort(array, base)
	return array
}
