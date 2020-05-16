package encrypt

import (
	"testing"
	"crypto/rand"
	"fmt"

)


func TestBase58Encode(t *testing.T)  {
	
	str := []byte("hello world")
	result := Encode58(str)

	if result != "StV1DL6CwTryKyV" {
		t.Errorf("base58 encode error")
	}
	t.Logf("result %s", result)

}

func TestBase58Decode(t *testing.T)  {
	
	result := Decode58("StV1DL6CwTryKyV")

	// if err != nil {
	// 	t.Errorf("base58 decode error")
	// }
	t.Logf("result %s", string(result))

}

func getRandomBytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
	}
	return b
}

func BenchmarkEncode58(b *testing.B) {


	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Encode58([]byte("hello world"))
	}
}