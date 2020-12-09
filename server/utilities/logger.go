package utilities

import (
	"log"
	"os"
)

var(
	infoLog, errLog  *log.Logger
)

func init()  {
	infoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}