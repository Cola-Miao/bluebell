package main

import (
	"bluebell/router"
	"log/slog"
)

func main() {
	defer closure()
	var err error
	if err = initialize(); err != nil {
		slog.Error("infrastructure init failed", "error: ", err.Error())
		panic(err)
	}
	if err = router.SetupRouter(); err != nil {
		slog.Error("server start failed", "error: ", err.Error())
		panic(err)
	}
	return
}
