package service

import (
	"bluebell/logic"
	"bluebell/model"
	"bluebell/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommunityList(c *gin.Context) {
	cs, err := logic.CommunityList()
	if err != nil {
		utils.WebErrorMessage(c, err, "get community list failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"communities": cs,
	})
}

func CommunityInfo(c *gin.Context) {
	name := c.Param("name")
	cm, err := logic.FindCommunityByName(name)
	if err != nil {
		utils.WebErrorMessage(c, err, "get community information failed")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"community": cm,
	})
}

func CreateCommunity(c *gin.Context) {
	var cm model.Community
	var err error
	if err = c.ShouldBindJSON(&cm); err != nil {
		utils.WebErrorMessage(c, err, "parse form failed")
		return
	}

	cm.Administrator = c.GetString("username")
	cm.AdminUUID = c.GetInt64("uuid")
	if err = logic.CreateCommunity(&cm); err != nil {
		utils.WebErrorMessage(c, err, "create community failed")
		return
	}
	utils.WebMessage(c, "create community successful")
}

func ReadArticle(c *gin.Context) {
	uuidStr := c.Param("uuid")
	uuid, err := strconv.Atoi(uuidStr)
	if err != nil {
		utils.WebErrorMessage(c, err, "parse article id failed")
		return
	}
	art, err := logic.ReadArticle(int64(uuid))
	if err != nil {
		utils.WebErrorMessage(c, err, "get article failed")
		return
	}
	utils.WebMessage(c, art)
}

func CreateArticle(c *gin.Context) {
	var art model.Article
	if err := c.ShouldBindJSON(&art); err != nil {
		utils.WebErrorMessage(c, err, "parse form failed")
		return
	}

	art.Author = c.GetString("username")
	art.AuthorUUID = c.GetInt64("uuid")

	if err := logic.CreateArticle(&art); err != nil {
		utils.WebErrorMessage(c, err, "create article failed")
		return
	}
	utils.WebMessage(c, "create article success")
}
