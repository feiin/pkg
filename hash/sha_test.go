package hash

import "testing"

func TestSha1(t *testing.T) {
	src := "hello world"
	dst := Sha1(src)

	if dst != "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed" {
		t.Errorf("sha1 error")
	}
	t.Logf("result %s",dst)
}


func TestSha256(t *testing.T) {
	src := "hello world"
	dst := Sha256(src)

	if dst != "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9" {
		t.Errorf("sha256 error")
	}
	t.Logf("result %s",dst)
}


func TestSha512(t *testing.T) {
	src := "hello world"
	dst := Sha512(src)

	if dst != "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff499670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f" {
		t.Errorf("sha512 error")
	}
	t.Logf("result %s",dst)
}