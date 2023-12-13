package pointerface

import (
	"math/rand"
	"strconv"
	"time"
)

type Sortable interface {
	Length() int
	Less(i, j int) bool
	Swap(i, j int)
}

func Sort(foo Sortable) {
	for index := 1; index < foo.Length(); index++ {
		for i := 0; i < foo.Length()-index; i++ {
			if foo.Less(i+1, i) {
				foo.Swap(i, i+1)
			}
		}
	}
}

func IsSorted(foo Sortable) bool {
	for i := foo.Length() - 1; i > 0; i-- {
		if foo.Less(i, i-1) {
			return false
		}
	}
	return true
}

type IntArray []int

func (arr IntArray) Length() int {
	return len(arr)
}
func (arr IntArray) Less(i, j int) bool {
	return arr[i] < arr[j]
}
func (arr IntArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type StringArray []string

func (arr StringArray) Length() int {
	return len(arr)
}
func (arr StringArray) Less(i, j int) bool {
	return arr[i] < arr[j]
}
func (arr StringArray) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

type Float64Array []float64

func (arr Float64Array) Length() int {
	return len(arr)
}
func (arr Float64Array) Less(i, j int) bool {
	return arr[i] < arr[j]
}
func (arr Float64Array) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
func (arr Float64Array) List() string {
	var str string
	str += "["
	for index := 0; index < arr.Length(); index++ {
		str += strconv.FormatFloat(arr[index], 'f', -1, 64)
		if index != arr.Length()-1 {
			str += ", "
		}
	}
	str += "]"
	return str
}
func (arr Float64Array) String() string {
	return arr.List()
}
func (arr Float64Array) Fill(howMany int) {
	rand.Seed(int64(time.Now().Nanosecond()))
	for i := 0; i < howMany; i++ {
		arr[i] = rand.Float64()
	}
}

type Person struct {
	firstname string
	lastname  string
}

func MakePerson(first, last string) Person {
	return Person{first, last}
}

type Persons []Person

func (arr Persons) Length() int {
	return len(arr)
}
func (arr Persons) Less(i, j int) bool {
	if arr[i].firstname != arr[j].firstname {
		return arr[i].firstname < arr[j].firstname
	} else {
		return arr[i].lastname < arr[j].lastname
	}
}
func (arr Persons) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
