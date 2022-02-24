package math

import (
	"math/rand"
)

//RandomInt64  返回随机数 [min,max)
func RandomInt64(min, max int64) int64 {
	if min >= max || max == 0 {
		return max
	}
	return rand.Int63n(max-min) + min
}

//RandomInt 返回随机数 [min,max)
func RandomInt(min, max int) int {
	if min >= max || max == 0 {
		return max
	}
	return rand.Intn(max-min) + min
}

//RandomInt32 返回随机数 [min,max)
func RandomInt32(min, max int32) int32 {
	if min >= max || max == 0 {
		return max
	}
	return rand.Int31n(max-min) + min
}
