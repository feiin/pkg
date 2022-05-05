package arrays

import (
	"testing"
)

func TestArrayTEqual(t *testing.T) {
	var a = []int{1, 2, 3, 4}
	var b = []int{1, 2, 3, 4}
	if !Equal(a, b) {
		t.Errorf("not equal to b")
	}
}

func TestArrayTEqual2(t *testing.T) {
	var a = []int{1, 2, 3, 4}
	var b = []int{1, 2, 3}
	if Equal(a, b) {
		t.Errorf("not equal to b")
	}
}

func TestArrayTEqual3(t *testing.T) {
	var a []int
	var b = []int{1, 2, 3, 4}
	if Equal(a, b) {
		t.Errorf("not equal to b")
	}
}
