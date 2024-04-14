package model

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
