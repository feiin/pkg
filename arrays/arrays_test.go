package arrays

import (
	"testing"
)

func TestIntArrayEqual(t *testing.T) {
	var a = []int{1, 2, 3, 4}
	var b = []int{1, 2, 3, 4}
	if !IntEqual(a, b) {
		t.Errorf("not equal to b")
	}
}


func TestIntArrayEqual2(t *testing.T) {
	var a []int
	var b = []int{1, 2, 3, 4}
	if IntEqual(a, b) {
		t.Errorf("not equal to b")
	}
}



func TestStringArrayEqual(t *testing.T) {
	var a = []string { "hello", "world"}
	var b = []string { "hello", "world"}
	if !StringEqual(a, b) {
		t.Errorf("not equal to b")
	}
}


func TestStringArrayEqual2(t *testing.T) {
	var a = []string { "hello", "world"}
	var b  []string 
	if StringEqual(a, b) {
		t.Errorf("not equal to b")
	}
}