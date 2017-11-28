package logging

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	FulcrumLog = NewLog()
)

type Logger struct {
	LogFile *log.Logger
}

func NewLog() *Logger {

	timestamp := strconv.FormatInt(time.Now().UTC().UnixNano(), 10)

	logpath := "log/fulcrum-" + timestamp + ".log"

	file, err := os.Create(logpath)
	if err != nil {
		fmt.Printf("Testing\n")
		panic(err)
	}

	lf := log.New(file, "", log.LstdFlags|log.Lshortfile)
	l := Logger{LogFile: lf}

	return &l
}

func Log(sys string, msg string) bool {
	FulcrumLog.LogFile.Printf("(%s) - %s\n", sys, msg)
	return true
}
