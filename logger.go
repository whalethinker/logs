package logs

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	//logs2 "github.com/whale1017/momoyuer/biz/repository/logs"
	"go.uber.org/zap/zapcore"
)

var (
	//logger *zap.Logger
	Logger *zap.SugaredLogger
)

func init() {
	logger, err := zap.NewDevelopment(zap.Hooks(func(entry zapcore.Entry) error {
		//level := entry.Level.String()
		//time := entry.timeStep
		//caller := entry.Caller.String()
		//msg := entry.Message
		//stack := entry.Stack
		//err := logs2.Repository.AddLog(context.Background(), logs2.Log{
		//	Level:  level,
		//	timeStep:   time,
		//	Caller: caller,
		//	Msg:    msg,
		//	Stack:  stack,
		//})
		//
		//if err != nil {
		//	fmt.Println(fmt.Sprintf("errlog - logs add db failed, err:%v", err))
		//}
		return nil
	}))
	if err != nil {
		fmt.Println(fmt.Sprintf("[error] init logger error: %v", err))
		return
	}
	Logger = logger.
		Sugar().
		WithOptions(zap.AddCallerSkip(1))
	return
}

//func LogExit() error {
//	return logger.Sync()
//}

func CtxInfo(ctx context.Context, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	Logger.Infow(msg)
}

func CtxWarn(ctx context.Context, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	Logger.Warnw(msg)
}

func CtxError(ctx context.Context, template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	Logger.Errorw(msg)
}

func Info(template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	Logger.Infow(msg)
}

func Warn(template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	Logger.Warnw(msg)
}

func Error(template string, args ...interface{}) {
	msg := fmt.Sprintf(template, args...)
	Logger.Errorw(msg)
}
