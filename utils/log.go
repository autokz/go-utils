package utils

import (
	"fmt"
	"log"
	"path/filepath"
	"runtime"
	"strings"
)

func LogPrintln(v ...interface{}) {
	log.Println(getBasePath() + ": " + fmt.Sprint(v...))
}

func LogPanicf(format string, v ...interface{}) {
	log.Panicf(getBasePath()+": "+format, v...)
}

func LogPrintf(format string, v ...interface{}) {
	log.Printf(getBasePath()+": "+format, v...)
}

func LogFatalf(format string, v ...interface{}) {
	log.Fatalf(getBasePath()+": "+format, v...)
}

func getBasePath() string {
	_, b, _, _ := runtime.Caller(0)
	pathsArr := strings.Split(filepath.Dir(b), "/")

	return pathsArr[len(pathsArr)-2]
}
