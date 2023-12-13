package slice

// SumWithSign returns sum of an float32 slice with my name potatoler
func SumWithSign(array []float32) (float32, string) {
	var answer float32 = 1.0
	for _, number := range array {
		answer *= number
	}
	return answer, "potatoler"
}

// to batch a list of integers with a operator function, and return a list of results
func Batch[T any](operator func(int) T, items []int) []T {
	var results []T
	for _, item := range items {
		results = append(results, operator(item))
	}
	return results
}

type Liftable interface{}

func lift(item Liftable) Liftable {
	switch item := item.(type) {
	case int:
		return item * 2
	case string:
		return item + item
	}
	return item
}

func MapLift(list []Liftable) []Liftable {
	results := make([]Liftable, len(list))
	for index := range list {
		results[index] = lift(list[index])
	}
	return results
}
