package encrypt

import (
	"encoding/base64"
	"testing"
	// "fmt"
)

func Test3DesEncrypt(t *testing.T) {
	data := []byte("{\"name\":\"test\"}")
	key := []byte("123123123123123123123123")
	result, err := TripleDesECBEncrypt(data, key, PADDING_PKCS5)

	if err != nil {
		t.Error(err)
	}

	t.Logf("Test3DesEncrypt result:%s", base64.StdEncoding.EncodeToString(result))
}

func Test3DesDecrypt(t *testing.T) {

	data, _ := base64.StdEncoding.DecodeString("zIlxdewS0LO8voXxZJs1bg==")
	key := []byte("123123123123123123123123")
	result, err := TripleDesECBDecrypt(data, key, PADDING_PKCS5)

	if err != nil {
		t.Error(err)
	}

	t.Logf("Test3DesDecrypt result:%s", string(result))
}

func Test3DesCBCEncrypt(t *testing.T) {

	iv := []byte("abcdefgh")

	data := []byte("{\"name\":\"test\"}")
	key := []byte("123123123123123123123123")
	result, err := TripleDesCBCEncrypt(data, key, iv, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	str := base64.StdEncoding.EncodeToString(result)

	if str != "8C7+1KZWr4qKiXslaNk5Fg==" {
		t.Error("Test3DesCBCEncrypt error")

	}

	t.Logf("Test3DesCBCEncrypt result:%s", str)
}

func Test3DesCBCDecrypt(t *testing.T) {
	iv := []byte("abcdefgh")

	data, _ := base64.StdEncoding.DecodeString("8C7+1KZWr4qKiXslaNk5Fg==")
	key := []byte("123123123123123123123123")
	result, err := TripleDesCBCDecrypt(data, key, iv, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	if string(result) != "{\"name\":\"test\"}" {
		t.Error("Test3DesCBCDecrypt error")
	}
	t.Logf("Test3DesCBCDecrypt result:%s", string(result))
}
