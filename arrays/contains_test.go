package arrays

import (
	"testing"
)

func TestContains(t *testing.T) {
	var arr = []string{"hello", "world"}
	item := "world"

	i := Contains(arr, item)

	if i != 1 {
		t.Errorf("not found word")
	}
}

func TestContainsInt(t *testing.T) {
	var arr = []int{1, 2, 3, 4}
	i := Contains(arr, 30)
	if i != -1 {
		t.Errorf("not found val")
	}
}

func TestContainsInt64(t *testing.T) {
	var arr = []int64{1, 2, 3, 4}
	i := Contains(arr, 4)
	if i != 3 {
		t.Errorf("not found val")
	}
}

func TestContainsUInt64(t *testing.T) {
	var arr = []uint64{1, 2, 3, 4}
	i := Contains(arr, 4)
	if i != 3 {
		t.Errorf("not found val")
	}
}
