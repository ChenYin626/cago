package logger

import (
	"context"
	"net/url"
	"os"

	"github.com/codfrm/cago/config"
	"github.com/codfrm/cago/pkg/logger/loki"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Config struct {
	Level string      `yaml:"level"`
	Debug bool        `yaml:"debug"`
	Loki  *LokiConfig `yaml:"loki"`
}

func InitWithConfig(ctx context.Context, config *config.Config, opts ...Option) (*zap.Logger, error) {
	cfg := &Config{}
	if err := config.Scan("logger", cfg); err != nil {
		return nil, err
	}
	if cfg.Level != "" {
		opts = append(opts, Level(cfg.Level))
	}
	if cfg.Debug {
		opts = append(opts, Debug())
	}
	if cfg.Loki != nil {
		opts = append(opts, WithLoki(&LokiConfig{
			Level: cfg.Loki.Level,
			Url:   cfg.Loki.Url,
		}))
	}
	return Init(ctx, opts...)
}

func Init(ctx context.Context, opt ...Option) (*zap.Logger, error) {
	options := &Options{}
	for _, o := range opt {
		o(options)
	}
	core := make([]zapcore.Core, 0, 1)
	level := toLevel(options.level)
	levelEnable := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= level
	})
	if options.debug {
		core = append(core, zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			levelEnable,
		))
	} else {
		core = append(core, zapcore.NewCore(
			zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
			zapcore.AddSync(os.Stdout),
			levelEnable,
		))
	}
	if options.loki != nil {
		u, err := url.Parse(options.loki.Url)
		if err != nil {
			return nil, err
		}
		level := toLevel(options.loki.Level)
		lokiCore, err := loki.NewLokiCore(ctx, u, zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return lvl >= level
		}))
		if err != nil {
			return nil, err
		}
		core = append(core, lokiCore)
	}

	logger := zap.New(zapcore.NewTee(core...))
	if options.fields != nil {
		logger = logger.With(options.fields...)
	}
	return logger, nil
}

func toLevel(level string) zapcore.Level {
	switch level {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	}
	return zap.InfoLevel
}
