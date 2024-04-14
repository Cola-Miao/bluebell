package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
)

func main() {
	if err := initialize(); err != nil {
		slog.Error("server start failed", "error: ", err.Error())
		panic(err)
	}
	slog.Info("server run")
	fmt.Println("server run")
}
