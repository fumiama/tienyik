package tienyik

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

const (
	ETPYE_AES_CBC = "2"
)

type TYAES [32]byte

func NewTYAES(rawKey []byte) (tya TYAES) {
	if len(rawKey) != 32 {
		panic("len(key) must == 32")
	}
	copy(tya[:], rawKey)
	return
}

func pkcs7Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padtext := make([]byte, padding)
	for i := range padtext {
		padtext[i] = byte(padding)
	}
	return append(data, padtext...)
}

func pkcs7Unpadding(data []byte) []byte {
	if len(data) == 0 {
		return data
	}
	padding := int(data[len(data)-1])
	if padding > len(data) || padding == 0 {
		return data
	}
	return data[:len(data)-padding]
}

func (tya TYAES) Encrypt(b []byte) []byte {
	blk, err := aes.NewCipher(tya[:])
	if err != nil {
		panic(err)
	}

	paddedData := pkcs7Padding(b, aes.BlockSize)

	mode := cipher.NewCBCEncrypter(blk, make([]byte, aes.BlockSize))
	ciphertext := make([]byte, len(paddedData))
	mode.CryptBlocks(ciphertext, paddedData)

	return ciphertext
}

func (tya TYAES) Decrypt(b []byte) ([]byte, error) {
	blk, err := aes.NewCipher(tya[:])
	if err != nil {
		return nil, err
	}

	if len(b) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	if len(b)%aes.BlockSize != 0 {
		return nil, errors.New("invalid ciphertext length")
	}

	mode := cipher.NewCBCDecrypter(blk, make([]byte, aes.BlockSize))
	plaintext := make([]byte, len(b))
	mode.CryptBlocks(plaintext, b)

	return pkcs7Unpadding(plaintext), nil
}
