package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(key []byte) string {
	// 使用 md5.Sum 计算 MD5 哈希
	hash := md5.Sum([]byte(key))

	// 将 MD5 哈希转换为十六进制字符串
	return fmt.Sprintf("%x", hash)
}
