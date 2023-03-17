package logtool

import (
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
	"time"
)

const (
	infoStr      = "%s\n[info] "
	warnStr      = "%s\n[warn] "
	errStr       = "%s\n[error] "
	traceErrStr  = "%s %s\n[error][rows:%v] %s"
	traceSlowStr = "%s %s\n[warn][rows:%v] %s"
	traceInfoStr = "%s\n[info][rows:%v] %s"
)

type sqlDbLogger struct {
	logLevel      logger.LogLevel
	slowThreshold time.Duration
}

func NewSqlDbLogger(slowThreshold int) *sqlDbLogger {
	return &sqlDbLogger{slowThreshold: time.Duration(slowThreshold) * time.Millisecond}
}

func (l *sqlDbLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.logLevel = level
	return l
}

func (l *sqlDbLogger) Info(_ context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Info {
		logx.Debugf(infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}
func (l *sqlDbLogger) Warn(_ context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Warn {
		logx.Infof(infoStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (l *sqlDbLogger) Error(_ context.Context, msg string, data ...interface{}) {
	if l.logLevel >= logger.Error {
		logx.Errorf(errStr+msg, append([]interface{}{utils.FileWithLineNum()}, data...)...)
	}
}

func (l *sqlDbLogger) Trace(_ context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.logLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.logLevel >= logger.Error && !errors.Is(err, gorm.ErrRecordNotFound):
		sql, rows := fc()
		if rows == -1 {
			logx.WithDuration(elapsed).Errorf(traceErrStr, utils.FileWithLineNum(), err, "-", sql)
		} else {
			logx.WithDuration(elapsed).Errorf(traceErrStr, utils.FileWithLineNum(), err, rows, sql)
		}
	case elapsed > l.slowThreshold && l.logLevel >= logger.Warn:
		sql, rows := fc()
		slowLog := fmt.Sprintf(">= %v", l.slowThreshold)
		if rows == -1 {
			logx.WithDuration(elapsed).Slowf(traceSlowStr, utils.FileWithLineNum(), slowLog, "-", sql)
		} else {
			logx.WithDuration(elapsed).Slowf(traceSlowStr, utils.FileWithLineNum(), slowLog, rows, sql)
		}
	case l.logLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			logx.WithDuration(elapsed).Infof(traceInfoStr, utils.FileWithLineNum(), "-", sql)
		} else {
			logx.WithDuration(elapsed).Infof(traceInfoStr, utils.FileWithLineNum(), rows, sql)
		}
	}
}
