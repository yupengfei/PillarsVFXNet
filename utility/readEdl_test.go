package utility

import (
	"fmt"
	"testing"
)

func Test_ReadEdl(t *testing.T) {
	s, err := ReadEdl("../doc/Untitled Sequence.01.edl")
	if err != nil && err.Error() != "EOF" {
		fmt.Println(err.Error())
	}
	fmt.Println("len-------->", len(s))
	Dump(s)
}
