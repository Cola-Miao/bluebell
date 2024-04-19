package main

import (
	"bluebell/config"
	"bluebell/dao/msq"
	"bluebell/dao/rdb"
	"bluebell/logger"
	"bluebell/utils"
	"flag"
	"fmt"
	"log/slog"
	"strings"
)

func initialize() error {
	var err error
	parseFlag()
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
	if err = utils.InitSFNode(); err != nil {
		return err
	}
	return nil
}

func parseFlag() {
	var model string
	flag.StringVar(&model, "model", "release", "choose release or debug model to start")
	flag.Parse()
	config.Cfg.Model = strings.ToUpper(model)
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
