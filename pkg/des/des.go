package des

import (
	"crypto/des"
	"github.com/andreburgaud/crypt2go/ecb"
)

func Encrypt(data, byteKey []byte) ([]byte, error) {
	data = validateTextLength(data)
	cipher, err := des.NewCipher(byteKey)
	if err != nil {
		return nil, err
	}
	encrypted := make([]byte, len(data))
	mode := ecb.NewECBEncrypter(cipher)
	mode.CryptBlocks(encrypted, data)
	return encrypted, nil
}

func Decrypt(data, byteKey []byte) ([]byte, error) {
	cipher, err := des.NewCipher(byteKey)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	mode := ecb.NewECBDecrypter(cipher)
	mode.CryptBlocks(decrypted, data)
	return decrypted, nil
}

func validateTextLength(data []byte) []byte {
	if rem := len(data) % 8; rem != 0 {
		data = fillLackingBytes(data, 8-rem)
	}
	return data
}

func fillLackingBytes(data []byte, rem int) []byte {
	for i := 0; i < rem; i++ {
		data = append(data, byte(32))
	}

	return data
}
