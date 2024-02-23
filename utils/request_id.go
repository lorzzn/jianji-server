package utils

import (
	"github.com/google/uuid"
)

// GenerateRequestId 会生成一个 uuid 格式的 request id
func GenerateRequestId() string {
	// 生成一个随机的 UUID
	randomUUID := uuid.New()

	return randomUUID.String()
}
