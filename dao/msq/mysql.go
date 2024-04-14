package msq

import (
	"bluebell/config"
	"bluebell/model"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var msq *sqlx.DB

func Init() error {
	mysqlCfg, err := config.Cfg.Mysql()
	if err != nil {
		return fmt.Errorf("read mysql config failed: %w", err)
	}
	if err = createDB(mysqlCfg); err != nil {
		return fmt.Errorf("create database failed: %w", err)
	}
	if err = connectDB(mysqlCfg); err != nil {
		return fmt.Errorf("connect database failed: %w", err)
	}
	return nil
}

func connectDB(mysqlCfg *model.MysqlCfg) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True",
		mysqlCfg.User, mysqlCfg.Password, mysqlCfg.Addr, mysqlCfg.DBName)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	msq = db
	return nil
}

func createDB(mysqlCfg *model.MysqlCfg) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/",
		mysqlCfg.User, mysqlCfg.Password, mysqlCfg.Addr)

	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", mysqlCfg.DBName)
	if _, err = db.Exec(query); err != nil {
		return err
	}
	return nil
}
