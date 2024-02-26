package utils

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func hex(array []byte) string {
	var sb strings.Builder
	for _, b := range array {
		sb.WriteString(fmt.Sprintf("%02x", b))
	}
	return sb.String()
}

func MD5Hex(message string) string {
	hash := md5.New()
	_, err := hash.Write([]byte(message))
	if err != nil {
		return ""
	}
	hashBytes := hash.Sum(nil)
	return hex(hashBytes)
}
