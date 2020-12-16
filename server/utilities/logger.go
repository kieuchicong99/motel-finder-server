package utilities

import (
	"log"
	"os"
)

var(
	InfoLog, ErrLog  *log.Logger
)

func init()  {
	InfoLog = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrLog = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}