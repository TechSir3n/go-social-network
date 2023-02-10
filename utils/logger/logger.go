package logger

import (
	"context"
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
	formatter "github.com/t-tomalak/logrus-easy-formatter"
)

var Logger = logrus.New()

func GetFile() *os.File {
	filePath, err := os.OpenFile("logger.log", os.O_APPEND | os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err, "[ERROR] open file")
		return nil
	}
	return filePath
}

func init() {
	file := GetFile()
	Logger = &logrus.Logger{
		Out:          io.MultiWriter(file),
		Level:        logrus.DebugLevel,
		ReportCaller: true,
		Formatter: &formatter.Formatter{
			TimestampFormat: "2001-12-12 12:12:12",
			LogFormat:       "[%lvl%]: %time% - %msg%\n",
		},
	}
}

func Trace(args ...interface{}) {
	Logger.Trace(args)
}

func Debug(args ...interface{}) {
	Logger.Debug(args)
}

func Info(args ...interface{}) {
	Logger.Info(args)
}

func Warn(args ...interface{}) {
	Logger.Warn(args)
}

func Error(args ...interface{}) {
	Logger.Error(args)
}

func Fatal(args ...interface{}) {
	Logger.Fatal(args)
}

func Panic(args ...interface{}) {
	Logger.Panic(args)
}

func WithContext(ctx context.Context) {
	Logger.WithContext(ctx)
}
