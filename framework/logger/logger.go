package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	logger *Logger
)

type Logger struct {
	*logrus.Logger
}

func init() {
	logger = &Logger{logrus.New()}
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "logs/server.log",
		MaxSize:    100,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	}
	logger.SetOutput(io.MultiWriter(os.Stdout, lumberjackLogger))
	logger.SetFormatter(&LogFormater{})
	logger.SetReportCaller(true)
}

type LogFormater struct{}

func (m *LogFormater) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.In(time.FixedZone("CST", 0)).Format("2006-01-02 15:04:05.000")
	hostName, _ := os.Hostname()
	level := getLogLevel(entry.Level)

	var fileName, functionName string
	var lineNumber int
	if entry.HasCaller() {
		fileName = filepath.Base(entry.Caller.File)
		lineNumber = entry.Caller.Line
		functionName = entry.Caller.Function
		if lastIndex := strings.LastIndex(functionName, "."); lastIndex >= 0 {
			functionName = functionName[lastIndex+1:]
		}
	} else {
		fileName = "Nil"
		lineNumber = 0
		functionName = "Nil"
	}
	newLog := fmt.Sprintf("[L]%s|%s|%s|%s|%d|%s|%s \n",
		timestamp, level, hostName, fileName, lineNumber, functionName, entry.Message)

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func getLogLevel(level logrus.Level) string {
	switch level {
	case logrus.DebugLevel:
		return "DBG"
	case logrus.InfoLevel:
		return "INF"
	case logrus.WarnLevel:
		return "WAR"
	case logrus.ErrorLevel:
		return "ERR"
	case logrus.FatalLevel:
		return "FAT"
	case logrus.PanicLevel:
		return "PAN"
	default:
		return "UNK"
	}
}

func GetLogger() *Logger {
	return logger
}
