package mysqlUtility

import (
	"PillarsPhenomVFXWeb/pillarsLog"
	"PillarsPhenomVFXWeb/utility"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// begin with capitial wordd so it can be accessed by outer
var DBConn *sql.DB

func init() {
	DBConn = ConnectToDB()
}

func ConnectToDB() *sql.DB {
	// connection already exist
	if DBConn != nil {
		return DBConn
	}

	// connection not exist
	propertyMap := utility.ReadProperty("./mysql.properties")
	var userName, password, host, port, database string
	userName = propertyMap["DBUserName"]
	password = propertyMap["DBPassword"]
	host = propertyMap["DBIP"]
	port = propertyMap["DBPort"]
	database = propertyMap["DBDatabase"]

	sqlstring := userName + ":" + password + "@tcp(" + host + ":" + port + ")/" + database
	fmt.Println("connection", sqlstring)
	DBConn, err := sql.Open("mysql", sqlstring)
	if err != nil {
		pillarsLog.PillarsLogger.Print("can not connect to mysql server")
		pillarsLog.PillarsLogger.Print(err.Error())
	}
	return DBConn
}

func CloseDBConnection() {
	if DBConn != nil {
		DBConn.Close()
		DBConn = nil
	}
}
