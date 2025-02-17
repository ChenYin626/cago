package logger

import (
	"context"

	"github.com/codfrm/cago/configs"
	"github.com/codfrm/cago/pkg/logger/loki"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var logger *zap.Logger

// Logger 日志组件,核心组件,必须提前注册
func Logger(ctx context.Context, config *configs.Config) error {
	cfg := &Config{}
	if err := config.Scan("logger", cfg); err != nil {
		return err
	}
	cfg.lokiOptions = append(cfg.lokiOptions,
		loki.AppendLabels(zap.String("app", config.AppName)),
		loki.AppendLabels(zap.String("version", config.Version)),
		loki.AppendLabels(zap.String("env", string(config.Env))),
	)
	cfg.debug = config.Debug
	l, err := NewWithConfig(ctx, cfg)
	if err != nil {
		return err
	}
	logger = l
	return nil
}

func SetLogger(l *zap.Logger) {
	logger = l
}

// Default 默认日志,尽量不要使用,会丢失上下文信息
func Default() *zap.Logger {
	return logger
}

func Ctx(ctx context.Context) *zap.Logger {
	log, ok := ctx.Value(loggerKey).(*zap.Logger)
	if !ok {
		if gctx, ok := ctx.(*gin.Context); ok {
			return gctx.Request.Context().Value(loggerKey).(*zap.Logger)
		}
		return logger
	}
	return log
}
