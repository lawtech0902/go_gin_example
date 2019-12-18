package util

import (
	"crypto/md5"
	"encoding/hex"
)

// 对图片名进行md5
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	
	return hex.EncodeToString(m.Sum(nil))
}
