package encrypt

import (
	"crypto/cipher"
)

//ECBEncrypt
func ECBEncrypt(block cipher.Block, src []byte, padding PaddingMode) ([]byte, error) {
	blockSize := block.BlockSize()
	src = Padding(padding, src, blockSize)

	dst := make([]byte, len(src))

	for index := 0; index < len(src); index += blockSize {
		block.Encrypt(dst, src[index:index+blockSize])
		copy(dst[index:], dst)
	}
	 
	return dst, nil
}

//ECBDecrypt
func ECBDecrypt(block cipher.Block, src []byte, padding PaddingMode) ([]byte, error) {
	dst := make([]byte, len(src))
	blockSize := block.BlockSize()

	for index := 0; index < len(src); index += blockSize {
		block.Decrypt(dst, src[index:index+blockSize])
		copy(dst[index:], dst)
	}

	dst = UnPadding(padding, dst)

	return dst, nil
}