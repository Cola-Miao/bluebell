package logger

import (
	"bluebell/config"
	"errors"
	"gopkg.in/natefinch/lumberjack.v2"
	"log/slog"
)

const (
	JSON = iota
	Text
)

func Init() (err error) {
	logCfg, err := config.Cfg.Log()
	if err != nil {
		return err
	}

	logger := lumberjack.Logger{
		Filename:   logCfg.Filename,
		MaxSize:    logCfg.MaxSize,
		MaxAge:     logCfg.MaxAge,
		MaxBackups: logCfg.MaxBackups,
		LocalTime:  logCfg.LocalTime,
		Compress:   logCfg.Compress,
	}

	ho := slog.HandlerOptions{
		AddSource:   logCfg.AddSource,
		Level:       slog.Level(logCfg.Level),
		ReplaceAttr: nil,
	}

	var handle slog.Handler
	switch logCfg.Type {
	case JSON:
		handle = slog.NewJSONHandler(&logger, &ho)
	case Text:
		handle = slog.NewTextHandler(&logger, &ho)
	default:
		return errors.New("undefined value")
	}
	l := slog.New(handle)
	slog.SetDefault(l)

	return nil
}
