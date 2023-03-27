package logtool

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type AsynqLogger struct {
	logger logx.Logger
}

func NewAsynqLogger() *AsynqLogger {
	return &AsynqLogger{logger: logx.WithContext(context.Background())}
}

func (l AsynqLogger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

func (l AsynqLogger) Info(args ...interface{}) {
	l.logger.Info(args)
}

func (l AsynqLogger) Warn(args ...interface{}) {
	l.logger.Slow(args)
}

func (l AsynqLogger) Error(args ...interface{}) {
	l.logger.Error(args)
}

func (l AsynqLogger) Fatal(args ...interface{}) {
	l.logger.Error("fatal: ", args)
}
