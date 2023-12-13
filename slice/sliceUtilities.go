package slice

import "fmt"

// TimeExpand create slice with some times space
func TimesExpand(slice []int, factor int) []int {
	newSlice := make([]int, len(slice)*factor)
	copy(newSlice, slice)
	return newSlice
}

// Filter return a new slice holding only the elements that satisfly filter()
func Filter(slice []int, filter func(int) bool) []int {
	var answer []int
	for _, item := range slice {
		if filter(item) {
			answer = append(answer, item)
		}
	}
	return answer
}

// Inserts a slice to some position of another slice
func Insert(distance, source []int, index int) []int {
	var answer []int
	answer = append(answer, distance[:index]...)
	answer = append(answer, source...)
	answer = append(answer, distance[index:]...)
	return answer
}

// Remove a range of consecutive elements
func Excavate(slice []int, start, end int) []int {
	var answer []int
	answer = append(answer, slice[:start]...)
	if end != len(slice) {
		answer = append(answer, slice[end+1:]...)
	}
	return answer
}

// A simple tool that prints all elements of a slice with a quote going before
func Print(slice []int, description string) {
	fmt.Printf("%s\n", description)
	for index := range slice {
		fmt.Printf("%d ", slice[index])
	}
	fmt.Printf("\n")
}
