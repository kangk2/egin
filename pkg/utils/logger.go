package utils

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

type LoggerInstance struct {
	_channel string
}

var Logger LoggerInstance

var logrusInstance *logrus.Logger

func (l LoggerInstance) Channel(channel string) LoggerInstance {
	l._channel = channel
	return l
}

func (l LoggerInstance) Info(message interface{}, content ...interface{}) {
	logrusInstance.WithFields(logrus.Fields{
		"content": content,
		"channel": l._channel,
	}).Info(message)
}

func (l LoggerInstance) Error(message interface{}, content ...interface{}) {
	logrusInstance.WithFields(logrus.Fields{
		"content": content,
		"channel": l._channel,
	}).Error(message)
}

func init() {
	logrusInstance = loggerFactory()
	Logger = LoggerInstance{"app"}
}

func loggerFactory() *logrus.Logger {
	conf := Config.Logger
	switch conf.Type {
	case "stdout":
		return stdoutLogger()
	case "file":
		return fileLogger()
	default:
		return stdoutLogger()
	}
}

func stdoutLogger() *logrus.Logger {
	loggerConf := Config.Logger

	logger := logrus.New()

	logWriter := io.Writer(os.Stdout)
	logger.Out = logWriter

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.SetLevel(logrus.Level(loggerConf.Level))

	return logger
}

func fileLogger() *logrus.Logger {
	loggerConf := Config.Logger

	fileName := loggerConf.FileName

	logger := logrus.New()

	logWriter, _ := rotatelogs.New(
		// 分割后的文件名称
		fileName+".%Y%m%d.log",

		// 生成软链，指向最新日志文件
		rotatelogs.WithLinkName(fileName),

		// 设置最大保存时间(7天)
		rotatelogs.WithMaxAge(7*24*time.Hour),

		// 设置日志切割时间间隔(1天)
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	logger.Out = logWriter

	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	logger.SetLevel(logrus.Level(loggerConf.Level))

	return logger
}

// TODO
func esLogger() {}

func mongoLogger() {}
