package encrypt

import (
	"testing"
	"encoding/base64"
	// "fmt"
)


func TestDesEncrypt(t *testing.T) {
	data := []byte("{\"name\":\"test\"}")
	key := []byte("abcdefgh")
	restult, err := DesECBEncrypt(data,key, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	t.Logf("DesECBEncrypt result:%s", base64.StdEncoding.EncodeToString(restult))
}

func TestDesDecrypt(t *testing.T) {
	
	data, _ := base64.StdEncoding.DecodeString("nm69gUH//5iMDJkkwdQEYA==")
	key := []byte("abcdefgh")
	restult, err := DesECBDecrypt(data,key,PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	t.Logf("TestAesDecrypt result:%s", string(restult))
}


func TestDesCBCEncrypt(t *testing.T) {

	iv := []byte("abcdefgh")

	data := []byte("{\"name\":\"test\"}")
	key := []byte("12345678")
	result, err := DesCBCEncrypt(data,key,iv, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	str := base64.StdEncoding.EncodeToString(result)

	if str != "kl7U88C9OAPY7p8XQSFugA==" {
		t.Error("TestDesCBCDecrypt error")

	}

	t.Logf("TestDesCBCDecrypt result:%s", str)
}


func TestDesCBCDecrypt(t *testing.T) {
	iv := []byte("abcdefgh")
	
	data, _ := base64.StdEncoding.DecodeString("kl7U88C9OAPY7p8XQSFugA==")
	key := []byte("12345678")
	result, err := DesCBCDecrypt(data,key,iv, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	if string(result) != "{\"name\":\"test\"}" {
		t.Error("TestDesCBCDecrypt error")
	}
	t.Logf("TestDesCBCDecrypt result:%s", string(result))
}