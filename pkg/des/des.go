package des

import (
	"crypto/des"
)

func Encrypt(data, byteKey []byte) ([]byte, error) {
	key, err := des.NewCipher(byteKey)
	if err != nil {
		return nil, err
	}
	encrypted := make([]byte, len(data))
	for bs, be := 0, des.BlockSize; bs < len(data); bs, be = bs+des.BlockSize, be+des.BlockSize {
		key.Encrypt(encrypted[bs:be], data[bs:be])
	}
	return encrypted, nil
}

func Decrypt(data, byteKey []byte) ([]byte, error) {
	key, err := des.NewCipher(byteKey)
	if err != nil {
		return nil, err
	}
	decrypted := make([]byte, len(data))
	for bs, be := 0, des.BlockSize; bs < len(data); bs, be = bs+des.BlockSize, be+des.BlockSize {
		key.Decrypt(decrypted[bs:be], data[bs:be])
	}
	return decrypted, nil
}
