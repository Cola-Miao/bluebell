package utils

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func WebMessage(c *gin.Context, date any) {
	c.JSON(http.StatusOK, gin.H{
		"data": date,
	})
}

func WebErrorMessage(c *gin.Context, err error, mess string) {
	slog.Warn(mess, "error", err)
	c.JSON(http.StatusOK, gin.H{
		"error": mess,
	})
	c.Abort()
}
