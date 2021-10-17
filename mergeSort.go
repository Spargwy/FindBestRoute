package main

import "fmt"

func MergeSort(array []Path, filter string) []Path {
	lenght := len(array)
	if lenght == 1 {
		return array
	}

	middle := int(lenght / 2)
	var (
		left  = make([]Path, middle)
		right = make([]Path, lenght-middle)
	)

	for i := 0; i < lenght; i++ {
		if i < middle {
			left[i] = array[i]
		} else {
			right[i-middle] = array[i]
		}
	}

	return Merge(MergeSort(left, filter), MergeSort(right, filter), filter)
}

func Merge(left, right []Path, filter string) (result []Path) {
	result = make([]Path, len(left)+len(right))

	i := 0

	switch filter {
	case "time":
		for len(left) > 0 && len(right) > 0 {
			if float64(left[0].ElapsedTime) < float64(right[0].ElapsedTime) {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
			i++

		}
	case "price":
		for len(left) > 0 && len(right) > 0 {
			if left[0].Price < right[0].Price {
				result[i] = left[0]
				left = left[1:]
			} else {
				result[i] = right[0]
				right = right[1:]
			}
			i++

		}
	default:
		fmt.Printf("Incorrect filter: %s", filter)
		return []Path{}
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}
	return
}
