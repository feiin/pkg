package math


import (
	"testing"
)
func TestRandomInt(t *testing.T) {
	

	var result map[int]int = make(map[int]int)

	for i := 0; i < 1000000; i++ {

		 x := RandomInt(1,11)

		 if v, ok := result[x]; ok {
			result[x] = v + 1
		 } else {
			result[x] =  1
		 }
	}

	t.Logf("results random int %+v",result)


	var result2 map[int32]int = make(map[int32]int)

	for i := 0; i < 1000000; i++ {

		 x := RandomInt32(1,11)

		 if v, ok := result2[x]; ok {
			result2[x] = v + 1
		 } else {
			result2[x] =  1
		 }
	}

	t.Logf("results random int32 %+v",result2)


	var result3 map[int64]int = make(map[int64]int)

	for i := 0; i < 1000000; i++ {

		 x := RandomInt64(1,11)

		 if v, ok := result3[x]; ok {
			result3[x] = v + 1
		 } else {
			result3[x] =  1
		 }
	}

	t.Logf("results random int64 %+v",result3)
}