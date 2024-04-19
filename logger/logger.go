package logger

import (
	"bluebell/config"
	"errors"
	"fmt"
	"gopkg.in/natefinch/lumberjack.v2"
	"log/slog"
	"os"
)

const (
	JSON = iota
	Text
)

func Init() (err error) {
	logCfg, err := config.Cfg.Log()
	if err != nil {
		return fmt.Errorf("read logger config failed: %w", err)
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
	if config.Cfg.Model == "DEBUG" {
		handle = slog.NewJSONHandler(os.Stdout, &ho)
	} else {
		switch logCfg.Type {
		case JSON:
			handle = slog.NewJSONHandler(&logger, &ho)
		case Text:
			handle = slog.NewTextHandler(&logger, &ho)
		default:
			return errors.New("undefined value")
		}
	}
	l := slog.New(handle)
	slog.SetDefault(l)

	return nil
}
