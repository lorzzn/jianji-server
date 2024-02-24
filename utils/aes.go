package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

type AESKeyIv struct {
	Key string `json:"key"`
	Iv  string `json:"iv"`
}

// 去除填充
func unPad(src []byte) ([]byte, error) {
	length := len(src)
	unPadding := int(src[length-1])
	if unPadding > length {
		return nil, errors.New("填充无效")
	}
	return src[:(length - unPadding)], nil
}

func AESDecryptData(keyIv AESKeyIv, encryptedData string) ([]byte, error) {
	decodeStringBytes, err := base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return nil, err
	}

	// 创建 AES 解密器
	block, err2 := aes.NewCipher([]byte(keyIv.Key))
	if err2 != nil {
		return nil, err
	}

	// 使用 AES CBC 解密模式
	mode := cipher.NewCBCDecrypter(block, []byte(keyIv.Iv))

	// 解密数据
	plaintext := make([]byte, len(decodeStringBytes))
	mode.CryptBlocks(plaintext, decodeStringBytes)

	// 去除填充
	plaintext, err = unPad(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
