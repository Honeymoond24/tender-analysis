package logs

import (
	"github.com/Honeymoond24/tender-analysis/internal/application"
	"go.uber.org/zap"
)

type ZapLogger struct {
	sugar *zap.SugaredLogger
}

func (l *ZapLogger) Info(fields ...interface{}) {
	l.sugar.Info(fields...)
}

func (l *ZapLogger) Error(fields ...interface{}) {
	l.sugar.Error(fields...)
}

func (l *ZapLogger) Fatal(fields ...interface{}) {
	l.sugar.Fatal(fields...)
}

func NewLogger() (application.Logger, error) {
	rawLogger, err := zap.NewProduction()
	if err != nil {
		return nil, err
	}

	sugar := rawLogger.Sugar()
	return &ZapLogger{sugar: sugar}, nil
}
