package encrypt

import (
	"testing"
	"encoding/base64"
	// "fmt"
)


func TestAesEncrypt(t *testing.T) {
	data := []byte("{\"name\":\"test\"}")
	key := []byte("abcdefghabcdefgh")
	restult, err := AesECBEncrypt(data,key, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	t.Logf("TestAesEncrypt result:%s", base64.StdEncoding.EncodeToString(restult))
}

func TestAesDecrypt(t *testing.T) {
	
	data, _ := base64.StdEncoding.DecodeString("4MlR9b93s7/0UM77rW6S1Q==")
	key := []byte("abcdefghabcdefgh")
	restult, err := AesECBDecrypt(data,key,PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	t.Logf("TestAesDecrypt result:%s", string(restult))
}


func TestAesCBCEncrypt(t *testing.T) {

	iv := []byte("1234512345123451")

	data := []byte("{\"name\":\"test\"}")
	key := []byte("abcdefghabcdefgh")
	result, err := AesCBCEncrypt(data,key,iv, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	str := base64.StdEncoding.EncodeToString(result)

	if str != "h6aEbBD7wh4ek7EUuSWB+A==" {
		t.Error("AesCBCEncrypt error")

	}

	t.Logf("AesCBCEncrypt result:%s", str)
}


func TestAesCBCDecrypt(t *testing.T) {
	iv := []byte("1234512345123451")
	
	data, _ := base64.StdEncoding.DecodeString("h6aEbBD7wh4ek7EUuSWB+A==")
	key := []byte("abcdefghabcdefgh")
	result, err := AesCBCDecrypt(data,key,iv, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	if string(result) != "{\"name\":\"test\"}" {
		t.Error("AesCBCDecrypt error")
	}
	t.Logf("AesCBCDecrypt result:%s", string(result))
}