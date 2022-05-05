package arrays

import ()

func Contains[T comparable](arr []T, val T) (index int) {
	index = -1
	for i, value := range arr {
		if value == val {
			index = i
			break
		}
	}
	return index
}
