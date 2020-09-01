package mlog
 
import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"time"
	"context"
)

type MLogger struct {
	Slogger *zap.SugaredLogger
}

var (
	Logger MLogger
)
var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}
 

func InitLog(path,level string) {
	fileName := path
	logLevel := GetLoggerLevel(level)
	syncWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   1024, //1G
		LocalTime: true,
		Compress:  true,
	})
	encoder := zap.NewProductionEncoderConfig()
	encoder.EncodeTime = LocalTimeEncoder
	core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(logLevel))
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	Logger.Slogger = logger.Sugar()
}
func GetLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}


func LocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder){
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// 下面实现了zap方法的封装
func (c *MLogger)Debug(ctx context.Context,args ...interface{}) {
	c.Slogger.Debug(getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)Debugf(template string, ctx context.Context,args ...interface{}) {
	c.Slogger.Debugf(template,getNewArgs(ctx ,args)...)
}

func (c *MLogger)Debugw(ctx context.Context,msg string,args ...interface{}) {
	m :=[]interface{}{}
	cname:="Context"
	m =append(m,cname)
	m = append(m,ctxTransMap(ctx))
	m = append(m ,args ...)
	c.Slogger.Debugw(msg,m ...)
}
 
func  (c *MLogger)Info(ctx context.Context,args ...interface{}) {
	c.Slogger.Info(getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)Infof(template string,ctx context.Context, args ...interface{}) {
	c.Slogger.Infof(template, getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)Infow(ctx context.Context,msg string,args ...interface{}){
	m :=[]interface{}{}
	cname:="Context"
	m =append(m,cname)
	m = append(m,ctxTransMap(ctx))
	m = append(m ,args ...)
	c.Slogger.Infow(msg,m ...)
}
func (c *MLogger)Warn(ctx context.Context,args ...interface{}) {
	c.Slogger.Warn(getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)Warnf(template string,ctx context.Context, args ...interface{}) {
	c.Slogger.Warnf(template, getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)Warnw(ctx context.Context,msg string,args ...interface{}){
	m :=[]interface{}{}
	cname:="Context"
	m =append(m,cname)
	m = append(m,ctxTransMap(ctx))
	m = append(m ,args ...)
	c.Slogger.Warnw(msg,m ...)
}
func (c *MLogger)Error(ctx context.Context,args ...interface{}) {
	c.Slogger.Error(getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)Errorf(template string,ctx context.Context, args ...interface{}) {
	c.Slogger.Errorf(template,  getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)Errorw(ctx context.Context,msg string,args ...interface{}){
	m :=[]interface{}{}
	cname:="Context"
	m =append(m,cname)
	m = append(m,ctxTransMap(ctx))
	m = append(m ,args ...)
	c.Slogger.Errorw(msg,m ...)
}
func (c *MLogger)DPanic(ctx context.Context,args ...interface{}) {
	c.Slogger.DPanic(getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)DPanicf(template string,ctx context.Context, args ...interface{}) {
	c.Slogger.DPanicf(template, getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)DPanicw(ctx context.Context,msg string,args ...interface{}){
	m :=[]interface{}{}
	cname:="Context"
	m =append(m,cname)
	m = append(m,ctxTransMap(ctx))
	m = append(m ,args ...)
	c.Slogger.DPanicw(msg,m ...)
}
func (c *MLogger)Panic(ctx context.Context,args ...interface{}) {
	c.Slogger.Panic(getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)Panicf(template string, ctx context.Context,args ...interface{}) {
	c.Slogger.Panicf(template, getNewArgs(ctx ,args)...)
}
func (c *MLogger)Panicw(ctx context.Context,msg string,args ...interface{}){
	m :=[]interface{}{}
	cname:="Context"
	m =append(m,cname)
	m = append(m,ctxTransMap(ctx))
	m = append(m ,args ...)
	c.Slogger.Panicw(msg,m ...)
}
func (c *MLogger)Fatal(ctx context.Context,args ...interface{}) {
	c.Slogger.Fatal(getNewArgs(ctx ,args)...)
}
 
func (c *MLogger)Fatalf(template string, ctx context.Context,args ...interface{}) {
	c.Slogger.Fatalf(template, getNewArgs(ctx ,args)...)
}
func (c *MLogger)Fatalw(ctx context.Context,msg string,args ...interface{}){
	m :=[]interface{}{}
	cname:="Context"
	m =append(m,cname)
	m = append(m,ctxTransMap(ctx))
	m = append(m ,args ...)
	c.Slogger.Fatalw(msg,m ...)
}

func getNewArgs(ctx context.Context,args ...interface{}) (newargs []interface{}){
	newargs= append(newargs, ctxTransMap(ctx))
	newargs= append(newargs, args ...)
	return newargs
}


func getKVNewArgs(ctx context.Context,args ...interface{}) (newargs []interface{}){
	cname:="Context"
	newargs =append(newargs,cname)
	newargs = append(newargs, ctxTransMap(ctx))
	newargs = append(newargs ,args ...)
	return newargs
}


func ctxTransMap(ctx context.Context) map[string]interface{}{
	m :=make(map[string]interface{})
	m["ClientIP"]=ctx.Value("ClientIP")
	m["RequestID"]=ctx.Value("RequestID")
	m["DentifyID"]=ctx.Value("DentifyID")
	m["ServersID"]=ctx.Value("ServersID")
	return m
}
