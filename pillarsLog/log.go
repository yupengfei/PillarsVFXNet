package pillarsLog

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var PillarsLogger *log.Logger
var outFile *os.File

func init() {
	if PillarsLogger == nil {
		propertyMap := readProperty("./log.properties")
		logFileName := propertyMap["LogFile"]
		fmt.Println(logFileName)
		var err error
		outFile, err = os.OpenFile(logFileName, os.O_RDWR|os.O_TRUNC, 0666)
		if err != nil {
			panic(err.Error())
		}
		PillarsLogger = log.New(outFile, "", log.Ldate|log.Ltime|log.Llongfile)
	}
}

func CloseLogFile() {
	outFile.Close()
}

func readProperty(fileName string) map[string]string {
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println(fileName, err)
		return nil
	}
	buff := bufio.NewReader(file)
	propertyMap := make(map[string]string)
	for {
		line, err := buff.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		line = strings.Trim(line, "\n")
		propertyPair := strings.Split(line, "=")
		propertyMap[propertyPair[0]] = propertyPair[1]
	}
	return propertyMap
}
