package hash

import (
	"crypto/md5"
	"fmt"
)

//Md5String get md5 value of input string
func Md5String(input string) string {

	data := []byte(input)
	// h := md5.New()
	// h.Write(data)
    h := md5.Sum(data)
    md5str := fmt.Sprintf("%x", h)
    return md5str
}