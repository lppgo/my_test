package main

import (
	"context"

	"git.yeeuu.com/yeeuu/my_test/099_temp_test/logger"
)

var (
	mlog logger.MLogger
)

func init() {
	path := "./stdout.log"
	logger.InitLog(path, "debug")
	mlog = logger.Logger
}

func main() {
	//

	ctx := context.WithValue(context.Background(), "ClientIP", "127.0.0.1")
	ctx = context.WithValue(ctx, "RequestID", "fcvuosqav8912")
	ctx = context.WithValue(ctx, "DentifyID", 1000)
	ctx = context.WithValue(ctx, "ServersID", "mountain_go")
	// mlog := logger.MLogger
	str := "hello ccc"
	// logger.Logger.Info(ctx context.Context, args ...interface{})
	// mlog := logger.Logger
	mlog.Info(ctx, str)
	mlog.Infof("template", ctx, str)
	mlog.Debug(ctx, "debug一个信息")
	mlog.Debugf("模板", ctx, "debug模板信息")

	// mlog.Debugw("消息", ctx, "k1", "v1")
	// mlog.Infow("消息", ctx, "k1", "v1")
	// mlog.Warnw("消息", ctx, "k1", "v1")
	// mlog.DPanicw("消息", ctx, "k1", "v1")
	mlog.DPanicw("消息", ctx, "k1", "v1")
	// mlog.Debugw("消息", "Context", ctx, "k1", "v1")
	// mlog.Warn(ctx context.Context, args ...interface{})
	// mlog.Warnf(template string, ctx context.Context, args ...interface{})
	// mlog.Info(ctx context.Context, args ...interface{})

}
