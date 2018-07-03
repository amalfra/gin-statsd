package middleware

import (
	"fmt"
	"io"
	"log"
	"os"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
	logPrefix     = "[gin-statsd]"
)

type logLevel int

const (
	infoLevel logLevel = iota
	warningLevel
	errorLevel
)

// setupLogger will configure the standard log package to print in custom format
func setupLogger(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {
	infoLogger = log.New(infoHandle,
		fmt.Sprintf("%s INFO: ", logPrefix),
		log.Ldate|log.Ltime)

	warningLogger = log.New(warningHandle,
		fmt.Sprintf("%s WARNING: ", logPrefix),
		log.Ldate|log.Ltime)

	errorLogger = log.New(errorHandle,
		fmt.Sprintf("%s ERROR: ", logPrefix),
		log.Ldate|log.Ltime)
}

func init() {
	setupLogger(os.Stdout, os.Stdout, os.Stderr)
}

// printLog outputs log to correct stream depending on level
func printLog(msg string, level logLevel) {
	var l *log.Logger

	switch level {
	case infoLevel:
		l = infoLogger
	case warningLevel:
		l = warningLogger
	case errorLevel:
		l = errorLogger
	default:
		l = infoLogger
	}

	l.Println(msg)
}
