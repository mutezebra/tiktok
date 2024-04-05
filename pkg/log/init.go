package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Mutezebra/tiktok/config"
)

var LogrusObj *Logger

// InitLog initializes a logger and continuously updates its output
func InitLog() {
	logger := logrus.New()
	src, err := setOutPutFile()
	if err != nil {
		panic(err)
	}
	logger.Out = src
	if config.Conf.System.Status == "debug" {
		logger.SetLevel(logrus.DebugLevel)
	} else {
		logger.SetLevel(logrus.WarnLevel)
	}

	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return frame.Function, fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
	})

	LogrusObj = &Logger{
		Logger: logger,
		pool: sync.Pool{
			New: func() any {
				return new(strings.Builder)
			},
		},
	}

	go updateLogger()
}

func setOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		if config.Conf.System.OS == "windows" {
			logFilePath = dir + "\\logs\\"
		} else {
			logFilePath = dir + "/logs/"
		}
	}
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		err = os.Mkdir(logFilePath, 0777)
		if err != nil {
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := logFilePath + logFileName
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		_, err = os.Create(fileName)
		if err != nil {
			return nil, err
		}
	}
	src, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModeAppend)

	return src, err
}

// updateLogger update the output file for LogrusObj
func updateLogger() {
	oneday := int64(3600 * 24)
	for {
		now := time.Now().Unix()
		remain := oneday - (now % oneday)
		src, _ := setOutPutFile()
		LogrusObj.Out = src
		time.Sleep(time.Duration(remain))
	}
}
