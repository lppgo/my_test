package zaplogger

import (
	"fmt"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * srvName 服务名
 */
func LogConf(srvName, logpath, logfileName string) {
	now := time.Now()
	if logfileName == "" {
		logfileName = "log.log"
	}

	hook := &lumberjack.Logger{
		// Filename:   fmt.Sprintf("%s/%04d%02d%02d-%02d%02d%02d-%s", logpath, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), logfileName), //filePath
		Filename:   fmt.Sprintf("%s/%s-%04d%02d%02d-%s", logpath, srvName, now.Year(), now.Month(), now.Day(), logfileName), //filePath
		MaxSize:    500,                                                                                                     //MB                                                                                                                                                      // megabytes
		MaxBackups: 10,
		MaxAge:     100,   //days
		Compress:   false, // disabled by default
	}
	defer hook.Close()

	enConfig := zap.NewProductionEncoderConfig() //生成配置

	// 时间格式
	// enConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	enConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

	level := zap.InfoLevel
	w := zapcore.AddSync(hook)
	core := zapcore.NewCore(
		// zapcore.NewConsoleEncoder(enConfig), //编码器配置
		zapcore.NewJSONEncoder(enConfig), //编码器配置
		w,                                //打印到控制台和文件
		level,                            //日志等级
	)

	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Logger = logger.Sugar().Named(srvName)
	Logger.Info("log config completed...")
}
