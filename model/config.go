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
