package main

import (
	"bluebell/config"
	"bluebell/dao/msq"
	"bluebell/dao/rdb"
	"bluebell/logger"
	"fmt"
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
