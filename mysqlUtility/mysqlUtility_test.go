package mysqlUtility

import (
	"testing"
)

func TestConnectToDB(t *testing.T) {
	ConnectToDB()
}

func TestDBClose(t *testing.T) {
	CloseDBConnection()
}
