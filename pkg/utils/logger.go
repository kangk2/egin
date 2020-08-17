package utils

import (
    "fmt"
    rotatelogs "github.com/lestrrat-go/file-rotatelogs"
    "github.com/rifflock/lfshook"
    "github.com/sirupsen/logrus"
    "os"
    "time"
)

type LoggerInstance struct {
}

var Logger LoggerInstance

func (l LoggerInstance) Info(message interface{}) {
    logrusLogger().WithFields(logrus.Fields{
        "message": message,
    }).Info()
}

func logrusLogger() *logrus.Logger {
    loggerConf := Config.Logger

    // 日志文件
    fileName := loggerConf.FileName

    // 写入文件
    src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
    if err != nil {
        fmt.Println("err", err)
    }

    // 实例化
    logger := logrus.New()

    // 设置输出
    logger.Out = src

    // 设置日志级别
    logger.SetLevel(logrus.Level(loggerConf.Level))

    // 设置 rotatelogs
    logWriter, err := rotatelogs.New(
        // 分割后的文件名称
        fileName+".%Y%m%d.log",

        // 生成软链，指向最新日志文件
        rotatelogs.WithLinkName(fileName),

        // 设置最大保存时间(7天)
        rotatelogs.WithMaxAge(7*24*time.Hour),

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

    lfHook := lfshook.NewHook(writeMap, &logrus.JSONFormatter{
        TimestampFormat: "2006-01-02 15:04:05",
    })

    // 新增 Hook
    logger.AddHook(lfHook)

    return logger
}
