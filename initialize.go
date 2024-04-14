package main

import (
	"bluebell/config"
	"bluebell/dao/msq"
	"bluebell/dao/rdb"
	"bluebell/logger"
	"fmt"
	"log/slog"
)

func initialize() error {
	var err error
	if err = config.Cfg.Init(); err != nil {
		return fmt.Errorf("config init failed: %w", err)
	}
	if err = logger.Init(); err != nil {
		return fmt.Errorf("logger init failed: %w", err)
	}
	if err = msq.Init(); err != nil {
		return fmt.Errorf("mysql init failed: %w", err)
	}
	if err = rdb.Init(); err != nil {
		return fmt.Errorf("redis init failed: %w", err)
	}
	return nil
}

func closure() {
	var err error
	if err = msq.Close(); err != nil {
		slog.Warn("close mysql failed", "error", err.Error())
	}
	if err = rdb.Close(); err != nil {
		slog.Warn("close redis failed", "error", err.Error())
	}
	return
}
