package utils

import (
	"strings"
)

func GetCravatarURL(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	return "http://cravatar.cn/avatar/" + MD5Hex(email)
}
