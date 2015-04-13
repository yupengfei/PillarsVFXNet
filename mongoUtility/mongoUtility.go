package mongoUtility

import (
	"PillarsPhenomVFXWeb/utility"

	"gopkg.in/mgo.v2"
)

var Session *mgo.Session
var MetadataCollection *mgo.Collection

func init() {
	Session = ConnectToMgo()
}

func ConnectToMgo() *mgo.Session {
	if Session != nil {
		return Session
	}

	propertyMap := utility.ReadProperty("../mongo.properties")

	var host, database string
	userName := propertyMap["DBUserName"]
	password := propertyMap["DBPassword"]
	host = propertyMap["DBIP"]
	database = propertyMap["DBDatabase"]
	//fmt.Println(userName, ":", password, "@", host, "/", database)
	Session, errMgo := mgo.Dial(userName + ":" + password + "@" + host + "/" + database)
	if errMgo != nil {
		panic(errMgo.Error())
	}
	MetadataCollection = Session.DB("PillarsPhenomVFXWeb").C("MaterialMetadata")
	return Session
}

func CloseMgoConnection() {
	if Session != nil {
		Session.Close()
		Session = nil
	}
}
