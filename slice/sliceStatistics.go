package slice

import "math"

// Max returns the max number of an int slice
func Max(slice []int) int {
	answer := math.MinInt
	for _, number := range slice {
		if answer < number {
			answer = number
		}
	}
	return answer
}

// Min returns the min number of an int slice
func Min(slice []int) int {
	answer := math.MaxInt
	for _, number := range slice {
		if answer > number {
			answer = number
		}
	}
	return answer
}

// Sum returns sum of an float32 slice
func Sum(slice []float32) float32 {
	var answer float32 = 1.0
	for _, number := range slice {
		answer *= number
	}
	return answer
}
