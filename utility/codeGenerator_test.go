package utility

import (
	"fmt"
	"strconv"
	"testing"
)

func TestRandomInt(t *testing.T) {
	fmt.Println("tesing random int generetor " + strconv.Itoa(RandomInt()))
	fmt.Println("tesing random int generetor " + strconv.Itoa(RandomInt()))
}

func TestGenerateCode(t *testing.T) {
	str := "testing code generetor"
	fmt.Println("tesing random int generetor " + *(GenerateCode(&str)))
}
