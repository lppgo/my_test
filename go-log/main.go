package main

import (
	"fmt"
	"log"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	Logger *zap.SugaredLogger
)

type MyLogger struct {
	*zap.SugaredLogger
	// trace
	// span
}

// func init() {
// 	logpath := "./log/"
// 	logfileName := "log.log"
// 	zaplogger.LogConf(logpath, logfileName)
// }

func main() {
	logpath := "."
	logfileName := "log.log"
	srvName := "go-elk"
	LogConf(srvName, logpath, logfileName)

	tick := time.NewTicker(time.Second * 1)
	for {
		select {
		case <-tick.C:
			Logger.Debug("logger one ")
			Logger.Info("logger two ")
			Logger.Errorw("logger 3 ", "traceID", 10001, "spanID", 18988, "orderBook", 8189891)
		default:
			// Logger.Info("main_info", "logger default ")

		}
	}

	//test
	// mlog := MyLogger{}
	// mlog.Errorw(msg string, keysAndValues ...interface{})

}

func LogConf(srvname, logpath, logfileName string) {
	if logfileName == "" {
		logfileName = "log.log"
	}

	//  日志切分
	hook := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s-%s", logpath, srvname, logfileName), //filePath
		MaxSize:    1,                                                      // megabytes
		MaxBackups: 5,
		MaxAge:     10,    //days
		Compress:   false, // disabled by default
		LocalTime:  true,
	}
	defer hook.Close()
	/*zap 的 Config 非常的繁琐也非常强大，可以控制打印 log 的所有细节，因此对于我们开发者是友好的，有利于二次封装。
	  但是对于初学者则是噩梦。因此 zap 提供了一整套的易用配置，大部分的姿势都可以通过一句代码生成需要的配置。
	*/
	// enConfig := zap.NewProductionEncoderConfig() //生成配置
	enConfig := zapcore.EncoderConfig{
		MessageKey:    "msg",
		TimeKey:       "ts",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		FunctionKey:   zapcore.OmitKey,
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		// EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 时间格式
	// enConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	// YYYY-MM-DD HH:mm:ss.SSS
	enConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	// enConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

	w := zapcore.AddSync(hook)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(enConfig), //编码器配置
		w,                                //打印到控制台和文件
		zap.DebugLevel,                   //日志等级
	)

	// logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	logger := zap.New(core, zap.AddCaller()) //AddCaller()为显示文件名和行号
	// _log := log.New(hook, "", log.LstdFlags)
	Logger = logger.Sugar().Named(srvname)
	// Logger = logger.Sugar()
	// _log.Println("log config completed...")
	log.Println("log config completed...")
}
