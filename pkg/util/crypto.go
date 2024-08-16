package util

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
)

func Sha256(raw []byte) []byte {
	bs := sha256.Sum256(raw)
	return bs[:]
}

func HmacSha256(raw []byte, secret []byte) []byte {
	h := hmac.New(sha256.New, secret)
	_, _ = h.Write(raw)
	return h.Sum(nil)
}

func HmacRipeMD160(message, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	_, _ = h.Write(message)
	return h.Sum(nil)
}

func Md5(row []byte) []byte {
	h := md5.New()
	h.Write(row)
	return h.Sum(nil)
}
