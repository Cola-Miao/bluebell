package service

import (
	"bluebell/dao/msq"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "ok",
	})
}

func Private(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": c.GetString("username"),
	})
}

func TestFunc(c *gin.Context) {
	data := c.Query("data")
	res, err := msq.UserExist(data)

	c.JSON(http.StatusOK, gin.H{
		"result": res,
		"error":  err,
	})
}
