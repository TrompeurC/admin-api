package log

import (
	"admin-api/common/config"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"time"
)

var log *logrus.Logger

var logToFile *logrus.Logger

// 日志文件名
var loggerFile string

// 设置文件地址
func setLogFile(file string) {
	loggerFile = file
}

// 初始化
func init() {
	logConfig := config.Config.Log
	setLogFile(filepath.Join(logConfig.Path, logConfig.Name))
}

// Log 方法调用
func Log() *logrus.Logger {
	//文件输出
	if config.Config.Log.Model == "file" {
		return logFile()
	}
	if log == nil {
		log = logrus.New()
		log.Out = os.Stdin
		log.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
		log.SetLevel(logrus.DebugLevel)
	}
	return log
}

// 日志写入文件
func logFile() *logrus.Logger {
	if logToFile == nil {
		logToFile = logrus.New()
		logToFile.SetLevel(logrus.DebugLevel)
		// 返回写日志对象logWriter
		logWriter, _ := rotatelogs.New(
			// 分割线的文件名称
			loggerFile+"_%Y%m%d.log",
			// 设置最大保存时间
			rotatelogs.WithMaxAge(30*24*time.Hour),
			// 设置日志切割时间间隔(1天)
			rotatelogs.WithRotationTime(24*time.Hour),
		)
		writeMap := lfshook.WriterMap{
			logrus.InfoLevel:  logWriter,
			logrus.FatalLevel: logWriter,
			logrus.DebugLevel: logWriter,
			logrus.WarnLevel:  logWriter,
			logrus.ErrorLevel: logWriter,
			logrus.PanicLevel: logWriter,
		}
		//设置时间格式
		lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
		// 新增 Hook
		logToFile.AddHook(lfHook)
	}
	return logToFile
}
