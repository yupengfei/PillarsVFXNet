package utility

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5sum(input *string) []byte {
	h := md5.New()
	h.Write([]byte(*input))
	var x []byte = h.Sum(nil)
	y := make([]byte, 32)
	hex.Encode(y, x)
	return y
}
