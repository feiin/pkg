package arrays

import (
	"reflect"
)

//Contains returns the index of item in arr
func Contains(arr interface{}, item interface{}) (index int) {
	index =  -1
	switch reflect.TypeOf(arr).Kind() {
	case reflect.Slice, reflect.Array:
		items := reflect.ValueOf(arr)

		for  i := 0; i < items.Len(); i++ {
			if reflect.DeepEqual(item, items.Index(i).Interface()) {
				index = i
				break 
			}
		}
	}
	return index
}

//ContainsInt returns the index of val in arr
func ContainsInt(arr []int, val int) (index int) {
	index = -1

	for i,value := range arr {
		if value == val {
			index = i
			break
		}
	}
	return index
}

//ContainsInt64 returns the index of val in arr
func ContainsInt64(arr []int64, val int64) (index int) {
	index = -1

	for i,value := range arr {
		if value == val {
			index = i
			break
		}
	}
	return index
}

//ContainsUInt64 returns the index of val in arr
func ContainsUInt64(arr []uint64, val uint64) (index int) {
	index = -1

	for i,value := range arr {
		if value == val {
			index = i
			break
		}
	}
	return index
}

//ContainsString returns the index of val in arr
func ContainsString(arr []string, val string) (index int) {
	index = -1

	for i,value := range arr {
		if value == val {
			index = i
			break
		}
	}
	return index
}