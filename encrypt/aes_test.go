package encrypt

import (
	"encoding/base64"
	"testing"
	// "fmt"
)

func TestAesEncrypt(t *testing.T) {
	data := []byte("{\"name\":\"test\"}")
	key := []byte("abcdefghabcdefgh")
	restult, err := AesECBEncrypt(data, key, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	t.Logf("TestAesEncrypt result:%s", base64.StdEncoding.EncodeToString(restult))
}

func TestAesDecrypt(t *testing.T) {

	data, _ := base64.StdEncoding.DecodeString("4MlR9b93s7/0UM77rW6S1Q==")
	key := []byte("abcdefghabcdefgh")
	restult, err := AesECBDecrypt(data, key, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	t.Logf("TestAesDecrypt result:%s", string(restult))
}

func TestAesCBCEncrypt(t *testing.T) {

	iv := []byte("1234512345123451")

	data := []byte("{\"name\":\"test\"}")
	key := []byte("abcdefghabcdefgh")
	result, err := AesCBCEncrypt(data, key, iv, PADDING_ZEROS)

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
	result, err := AesCBCDecrypt(data, key, iv, PADDING_ZEROS)

	if err != nil {
		t.Error(err)
	}

	if string(result) != "{\"name\":\"test\"}" {
		t.Error("AesCBCDecrypt error")
	}
	t.Logf("AesCBCDecrypt result:%s", string(result))
}

func TestAesCBCDecrypt2(t *testing.T) {

	iv, _ := base64.StdEncoding.DecodeString("b70j+p7+471uH6Cdc2l7qw==")
	key, _ := base64.StdEncoding.DecodeString("nsSz5ITIwcCt0HmXTDOyhA==")

	// iv := []byte("1234512345123451")

	data, _ := base64.StdEncoding.DecodeString("Cc/2cwVCIKKq1HwqN5HPOvl/vlbtiCaDABMjF3Rodqu/UP9lqkE21x0ujn2IW4YEMEtc3U55dSdHDsTyRUE5VeFcgEMYxyr2Xy/F0iXD4bk0eFoIU5w3xHzSHfyuQngQwc03rVqNfI++plY2A0CUXVl9FxdVfSgs+ryQ3DvBeIam2WGcDHCQJ4e2+tV1hhEjjcAA7j4WWuahPoJtg/wm2CJdPs0APxXkLG502x/zBP2tej0UeAbOLKyXhEpXC6XulztLfLF4fK7nkyFKn/rBe3/kYz1nGiJhNDQ/wwpxJAT/4Kml6zsoxhj3/G5H6BKDEZeZYTRs1Etj8HBVxd/QYd0PYlUpJ4noHgyKLF3tvD3GWJhPmjYct4EIjEtEBpLw5SXJ4NWR6XjL/bbMbS1EaJsx/QV0V7Y3LFKMB0YEXUrnRfUQ/HneAVcsKqfUIKaL00cpu8EEx2G29JKQDGYA5xKWGNr5385FWpJpYH7hSKgiWBnaBho0QjEFXXslABY7e41sLY85hITSpi44Mj8eoA==")
	// key := []byte("abcdefghabcdefgh")
	result, err := AesCBCDecrypt(data, key, iv, PADDING_PKCS7)

	if err != nil {
		t.Error(err)
	}

	// if string(result) != "{\"name\":\"test\"}" {
	// 	t.Error("AesCBCDecrypt error")
	// }
	t.Logf("AesCBCDecrypt result:%s", string(result))
}
