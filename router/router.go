package router

import (
	"bluebell/config"
	"bluebell/router/service"
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func SetupRouter() error {
	serverCfg, err := config.Cfg.Server()
	if err != nil {
		return fmt.Errorf("read server config failed: %w", err)
	}

	r := gin.Default()
	public := r.Group("/")
	{
		public.GET("/health", service.Health)
	}
	srv := http.Server{
		Addr:    serverCfg.Addr,
		Handler: r,
	}

	go func() {
		err = srv.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Info("server failed", "error: ", err.Error())
			panic(err)
		}
	}()
	slog.Info("server run")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	if err = srv.Shutdown(ctx); err != nil {
		return err
	}
	slog.Info("server shutdown")

	return nil
}