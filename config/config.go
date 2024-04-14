package config

import (
	"bluebell/model"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log/slog"
)

type config struct {
	vp *viper.Viper
}

var Cfg config

func (c *config) Init() error {
	vp := viper.New()
	vp.SetConfigFile("config.yaml")
	if err := vp.ReadInConfig(); err != nil {
		return err
	}
	c.vp = vp
	vp.WatchConfig()
	vp.OnConfigChange(func(in fsnotify.Event) {
		slog.Info("config file has changed", "file", in.Name)
	})
	return nil
}

func (c *config) Log() (*model.LogCfg, error) {
	var logCfg model.LogCfg
	err := c.vp.UnmarshalKey("Log", &logCfg)
	return &logCfg, err
}

func (c *config) Mysql() (*model.MysqlCfg, error) {
	var mysqlCfg model.MysqlCfg
	err := c.vp.UnmarshalKey("Mysql", &mysqlCfg)
	return &mysqlCfg, err
}

func (c *config) Redis() (*model.RedisCfg, error) {
	var redisCfg model.RedisCfg
	err := c.vp.UnmarshalKey("Redis", &redisCfg)
	return &redisCfg, err
}
