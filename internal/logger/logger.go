package logger 

import (
	"os"
	"log"
	"io"
	"strings"
)

var (
	Debug *log.Logger
	Info  *log.Logger
	Warn  *log.Logger 
	Error *log.Logger 
)

func Init(level string) {
	devNull := io.Discard

	debugOut := devNull
	infoOut := devNull
	warnOut := devNull 
	errOut := os.Stderr

	level = strings.ToUpper(level)

	switch level {
	case "DEBUG":
		debugOut = os.Stdout
	case "INFO":
		infoOut = os.Stdout
	case "WARN":
		warnOut = os.Stdout
	}

	Debug = log.New(debugOut, "DEBUG: ", log.Ldate | log.Ltime | log.Lshortfile)
	Info = log.New(infoOut, "INFO: ", log.Ldate | log.Ltime)
	Warn = log.New(warnOut, "WARN: ", log.Ldate | log.Ltime)
	Error = log.New(errOut, "Error: ", log.Ltime | log.Ldate | log.Lshortfile)
}