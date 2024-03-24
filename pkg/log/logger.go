package log

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/Mutezebra/tiktok/config"
)

var LogrusObj *Logger

type Logger struct {
	*logrus.Logger
}

// Panic refactors the panic function within logrus. It records
// the stack info and outputs it to logrus`s 'Out'
func (l *Logger) Panic(v any) {
	const depth = 32
	var pcs [depth]uintptr
	n := runtime.Callers(2, pcs[:])
	st := make([]uintptr, n)
	st = pcs[0:n]
	var str strings.Builder
	str.WriteString(fmt.Sprintf("%v\n", v))
	for _, pc := range st {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\t%s:%d %s\n", file, line, fn.Name()))
	}
	fmt.Println(v)
	l.Logger.Panic(str.String())
}

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

	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.TextFormatter{ // 记录的是打日志的位置
		TimestampFormat: "2006-01-02 15:04:05",
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			fileName := path.Base(frame.File)
			return frame.Function, fmt.Sprintf("%s:%d", fileName, frame.Line)
		},
	})
	logger.Infoln("logs init success!")
	LogrusObj = &Logger{logger}
	go UpdateLogger()
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

// UpdateLogger update the output file for LogrusObj
func UpdateLogger() {
	oneday := int64(3600 * 24)
	for {
		now := time.Now().Unix()
		remain := oneday - (now % oneday)
		src, _ := setOutPutFile()
		LogrusObj.Out = src
		time.Sleep(time.Duration(remain))
	}
}

// Restore replaces "\\n" with "\n" in the panic msg.
func Restore(msg string) string {
	return strings.ReplaceAll(strings.ReplaceAll(msg, "\\n", "\n"), "\\t", "\t")
}
