package logger

import (
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// SetLogger Установка Logger-а
var (
	Info  *log.Logger
	Error *log.Logger
	Warn  *log.Logger
	Debug *log.Logger
)

const (
	LogInfo       = "logs/info.log"
	LogError      = "logs/error.log"
	LogWarning    = "logs/warning.log"
	LogDebug      = "logs/debug.log"
	LogMaxSize    = 25
	LogMaxBackups = 5
	LogMaxAge     = 30
	LogCompress   = true
)

func Init() *logrus.Logger {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		err = os.Mkdir("logs", 0755)
		if err != nil {
			return nil
		}
	}

	fileInfo, err := os.OpenFile(LogInfo, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil
	}
	fileError, err := os.OpenFile(LogError, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil
	}
	fileWarn, err := os.OpenFile(LogWarning, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil
	}
	fileDebug, err := os.OpenFile(LogDebug, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return nil
	}

	Info = log.New(fileInfo, "", log.Ldate|log.Lmicroseconds)
	Error = log.New(fileError, "", log.Ldate|log.Lmicroseconds)
	Warn = log.New(fileWarn, "", log.Ldate|log.Lmicroseconds)
	Debug = log.New(fileDebug, "", log.Ldate|log.Lmicroseconds)

	lumberLogInfo := &lumberjack.Logger{
		Filename:   LogInfo,
		MaxSize:    LogMaxSize, // megabytes
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,   //days
		Compress:   LogCompress, // disabled by default
		LocalTime:  true,
	}

	lumberLogError := &lumberjack.Logger{
		Filename:   LogError,
		MaxSize:    LogMaxSize, // megabytes
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,   //days
		Compress:   LogCompress, // disabled by default
		LocalTime:  true,
	}

	lumberLogWarn := &lumberjack.Logger{
		Filename:   LogWarning,
		MaxSize:    LogMaxSize, // megabytes
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,   //days
		Compress:   LogCompress, // disabled by default
		LocalTime:  true,
	}

	lumberLogDebug := &lumberjack.Logger{
		Filename:   LogDebug,
		MaxSize:    LogMaxSize, // megabytes
		MaxBackups: LogMaxBackups,
		MaxAge:     LogMaxAge,   //days
		Compress:   LogCompress, // disabled by default
		LocalTime:  true,
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, lumberLogInfo)

	Info.SetOutput(gin.DefaultWriter)
	Error.SetOutput(lumberLogError)
	Warn.SetOutput(lumberLogWarn)
	Debug.SetOutput(lumberLogDebug)

	return nil
}
