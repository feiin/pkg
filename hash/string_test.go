package hash


import (
	"testing"
)
func TestMd5String(t *testing.T) {
	val := Md5String("123456")
	if val != "e10adc3949ba59abbe56e057f20f883e" {
		t.Errorf("string md5 error :%s", val)
	}
}