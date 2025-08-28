package logger

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"github.com/tsfans/go/framework/utils"
)

var (
	logger *Logger
)

type Logger struct {
	*logrus.Logger
}

func init() {
	logger = &Logger{logrus.New()}
	logger.Level = logrus.DebugLevel
	logger.SetOutput(os.Stdout)
	logger.SetReportCaller(true)
	logger.SetFormatter(&LogFormater{})
}

func InitLogger(level, logFile string) {
	logPath := filepath.Dir(logFile)
	if _, err := os.Stat(logPath); err != nil {
		err = os.Mkdir(logPath, 0666)
		if err != nil {
			log.Panicf("Error creating log directory:%s", err)
		}
	}

	baseName := filepath.Base(logFile)
	suffix := filepath.Ext(baseName)
	prefix := strings.TrimRight(baseName, suffix)
	if hostname := os.Getenv("HOSTNAME"); hostname != "" {
		prefix = fmt.Sprintf("%s-%s", prefix, hostname)
	}

	logfmt := filepath.Join(logPath, prefix+"-%Y%m%d%H%M"+suffix)
	logWriter, err := rotatelogs.New(
		logfmt,
		rotatelogs.WithLinkName(filepath.Join(logPath, prefix+suffix)),
		rotatelogs.WithRotationTime(time.Duration(8)*time.Hour),
	)
	if err != nil {
		log.Panicf("Error creating log writer:%s", err)
	}

	logger.SetOutput(io.MultiWriter(os.Stdout, logWriter))
	logger.SetLevel(parseLogLevel(level))
}

type LogFormater struct{}

func (m *LogFormater) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.In(utils.CST).Format("2006-01-02 15:04:05.000")
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
	newLog := fmt.Sprintf("%s|%s|%s|%s|%d|%s|%s \n",
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

func parseLogLevel(level string) logrus.Level {
	switch level {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	case "fatal":
		return logrus.FatalLevel
	case "panic":
		return logrus.PanicLevel
	default:
		return logrus.InfoLevel
	}
}

func Get() *Logger {
	return logger
}
