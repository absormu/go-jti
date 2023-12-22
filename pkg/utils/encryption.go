package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"

	cm "github.com/absormu/go-jti/pkg/configuration"
)

// GetAESDecrypted decrypts given text in AES 256 CBC
func GetAESDecrypted(encrypted string) ([]byte, error) {
	key := cm.Config.MyKeySecret //32 digit
	iv := cm.Config.MyIvSecret   // 16 digit

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)

	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return nil, err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("block size cant be zero")
	}

	mode := cipher.NewCBCDecrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, ciphertext)
	ciphertext = PKCS5UnPadding(ciphertext)

	return ciphertext, nil
}

// PKCS5UnPadding  pads a certain blob of data with necessary data to be used in AES block cipher
func PKCS5UnPadding(src []byte) []byte {
	length := len(src)
	unpadding := int(src[length-1])

	return src[:(length - unpadding)]
}

// GetAESEncrypted encrypts given text in AES 256 CBC
func GetAESEncrypted(plaintext string) (string, error) {
	key := cm.Config.MyKeySecret //32 digit
	iv := cm.Config.MyIvSecret   // 16 digit

	var plainTextBlock []byte
	length := len(plaintext)

	if length%16 != 0 {
		extendBlock := 16 - (length % 16)
		plainTextBlock = make([]byte, length+extendBlock)
		copy(plainTextBlock[length:], bytes.Repeat([]byte{uint8(extendBlock)}, extendBlock))
	} else {
		plainTextBlock = make([]byte, length)
	}

	copy(plainTextBlock, plaintext)
	block, err := aes.NewCipher([]byte(key))

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, len(plainTextBlock))
	mode := cipher.NewCBCEncrypter(block, []byte(iv))
	mode.CryptBlocks(ciphertext, plainTextBlock)

	str := base64.StdEncoding.EncodeToString(ciphertext)

	return str, nil
}

// https://medium.com/insiderengineering/aes-encryption-and-decryption-in-golang-php-and-both-with-full-codes-ceb598a34f41
