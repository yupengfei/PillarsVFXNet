package utility

import (
	"fmt"
	"testing"
)

func TestReadProperty(t *testing.T) {
	propertyMap := ReadProperty("test.properties")
	fmt.Println(propertyMap["DBDatabase"])
	fmt.Println(propertyMap["DBIP"])
}
