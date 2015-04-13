package mongoUtility

import (
	"testing"
)

func Test_ConnectToMgo(t *testing.T) {
	session := ConnectToMgo()
	if session == nil {
		t.Error("connect to mongo 1 failed")
	}
}

func Test_CloseMgoConnection(t *testing.T) {
	CloseMgoConnection()
}
