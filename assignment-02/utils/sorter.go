package utils

type Sorter interface {
	SortIntArray([]int) interface{}
	SortFloatArray([]float64) interface{}
	SortStringArray([]string) interface{}
}

type sorter struct{}

func NewSorter() Sorter {
	return &sorter{}
}

func (s *sorter) SortIntArray(arr []int) interface{} {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func (s *sorter) SortFloatArray(arr []float64) interface{} {
	// Bubble sort for floats
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func (s *sorter) SortStringArray(arr []string) interface{} {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}
