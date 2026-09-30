package util

import (
	"crypto/md5"
	"encoding/hex"
)

func CheckMd5PasswordHash(password, hash string) bool {
	hasher := md5.New()
	hasher.Write([]byte(password))
	computedHash := hex.EncodeToString(hasher.Sum(nil))
	return computedHash == hash
}
