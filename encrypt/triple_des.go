package encrypt

import (
	"crypto/des"
)

//TripleDesECBEncrypt 3DES-ECB
func TripleDesECBEncrypt(src, key []byte, padding PaddingMode) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	return ECBEncrypt(block, src, padding)
}

//TripleDesECBDecrypt  3DES-ECB
func TripleDesECBDecrypt(src, key []byte, padding PaddingMode) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	return ECBDecrypt(block, src, padding)
}

//TripleDesCBCDecrypt 3DES-CBC
func TripleDesCBCEncrypt(src, key, iv []byte, padding PaddingMode) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	return CBCEncrypt(block, src, iv, padding)
}

//TripleDesCBCDecrypt 3DES-CBC
func TripleDesCBCDecrypt(src, key, iv []byte, padding PaddingMode) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}

	return CBCDecrypt(block, src, iv, padding)
}
