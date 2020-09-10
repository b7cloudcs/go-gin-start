package util

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// SHA1 hashes using sha1 algorithm
func SHA1(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}

// Md5 32位小写
func Md5(text string) string {
	data := []byte(text)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
