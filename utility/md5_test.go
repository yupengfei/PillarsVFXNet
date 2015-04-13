package utility

import (
	"bytes"
	"testing"
)

func TestMd5sum(t *testing.T) {
	message := "123456"
	test := Md5sum(&message)
	result := []byte("e10adc3949ba59abbe56e057f20f883e")
	if len(test) != len(result) {
		t.Error("Md5sum length mismatch")
	} else {
		if !bytes.Equal(test, result) {
			t.Error("Md5sum wrong")
		}
	}
}
