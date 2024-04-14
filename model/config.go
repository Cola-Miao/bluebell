package model

type LogCfg struct {
	// lumberjack
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	LocalTime  bool
	Compress   bool
	// slog
	AddSource bool
	Level     int
	Type      int
}

type MysqlCfg struct {
	User     string
	Password string
	Addr     string
	DBName   string
}

type RedisCfg struct {
	User     string
	Password string
	Addr     string
}

type ServerCfg struct {
	Addr string
}
