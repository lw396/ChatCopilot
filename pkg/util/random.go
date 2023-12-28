package util

import (
	"crypto/rand"
	"math/big"
)

const (
	UpperCase           = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerCase           = "abcdefghijklmnopqrstuvwxyz"
	Letter              = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	NumberWithUpperCase = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZ"
	NumberWithLowerCase = "0123456789abcdefghijkmnpqrstuvwxyz"
	NumberWithLetter    = "0123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnpqrstuvwxyz"
	Number              = "0123456789"
	Printable           = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz~@#$%&-_=+"
)

func RandomString(codeBytes string, length uint) string {
	b := make([]byte, length)
	for i := range b {
		index, _ := rand.Int(rand.Reader, big.NewInt(int64(len(codeBytes))))
		b[i] = codeBytes[index.Int64()]
	}
	return string(b)
}
