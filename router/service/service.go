package service

import (
	"bluebell/dao/rdb"
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
	rdb.VoteForArticle("tA", "tU", 666)

	c.JSON(http.StatusOK, gin.H{
		//"result": res,
		//"error":  err,
	})
}
