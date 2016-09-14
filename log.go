package log

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var logger = log.New(os.Stderr, "", log.Lshortfile|log.LstdFlags)

func LogInit(logpath string) error {

	os.Mkdir(logpath, 0666)

	s := fmt.Sprintf("%s/%s.log", logpath, filepath.Base(os.Args[0]))

	w, err := os.OpenFile(s, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0x666)
	if err != nil {

		return err
	}

	logger.SetFlags(log.Lshortfile | log.LstdFlags)
	logger.SetOutput(w)

	return nil
}

func Debug(v ...interface{}) { logger.Output(2, fmt.Sprint(v...)) }

func Debugf(format string, v ...interface{}) {
	logger.Output(2, fmt.Sprintf(format, v...))
}
func Error(v ...interface{}) {
	logger.Output(2, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	logger.Output(2, fmt.Sprintf(format, v...))
}
func Info(v ...interface{}) {
	logger.Output(2, fmt.Sprint(v...))
}
func Infof(format string, v ...interface{}) {
	logger.Output(2, fmt.Sprintf(format, v...))
}
func Print(v ...interface{}) {
	logger.Output(2, fmt.Sprint(v...))
}
func Printf(format string, v ...interface{}) {
	logger.Output(2, fmt.Sprintf(format, v...))
}
func Fatal(v ...interface{}) {
	logger.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func Fatalf(format string, v ...interface{}) {
	logger.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func SetFlags(flag int) {
	logger.SetFlags(flag)
}
