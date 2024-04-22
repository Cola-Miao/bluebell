package service

import (
	"bluebell/dao/rdb"
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
		utils.WebErrorMessage(c, err, model.ErrGetList.Error())
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
		utils.WebErrorMessage(c, err, "get community information")
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
		utils.WebErrorMessage(c, err, model.ErrParseForm.Error())
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
		utils.WebErrorMessage(c, err, "get article")
		return
	}
	utils.WebMessage(c, art)
}

func CreateArticle(c *gin.Context) {
	var art model.Article
	var err error
	if err = c.ShouldBindJSON(&art); err != nil {
		utils.WebErrorMessage(c, err, model.ErrParseForm.Error())
		return
	}

	art.Author = c.GetString("username")
	art.AuthorUUID = c.GetInt64("uuid")

	if err = logic.CreateArticle(&art); err != nil {
		utils.WebErrorMessage(c, err, "create article")
		return
	}
	if err = rdb.CreateArticle(art.UUID); err != nil {
		utils.WebErrorMessage(c, err, "cache article")
		return
	}
	utils.WebMessage(c, "create article success")
}

func ArticleList(c *gin.Context) {
	offset := c.Query("offset")
	size := c.Query("size")
	as, err := logic.ArticleList(offset, size)
	if err != nil {
		utils.WebErrorMessage(c, err, model.ErrGetList.Error())
		return
	}
	utils.WebMessage(c, as)
}

func ArticleListByCommunity(c *gin.Context) {
	comID := c.Param("id")
	offset := c.Query("offset")
	size := c.Query("size")
	as, err := logic.ArticleListByCommunity(comID, offset, size)
	if err != nil {
		utils.WebErrorMessage(c, err, model.ErrGetList.Error())
		return
	}
	utils.WebMessage(c, as)
}

func VoteForArticle(c *gin.Context) {
	var vf model.FormVote
	var err error
	if err = c.ShouldBindJSON(&vf); err != nil {
		utils.WebErrorMessage(c, err, model.ErrParseForm.Error())
		return
	}
	userID := c.GetInt64("uuid")
	if err = logic.VoteForArticle(vf.UUID, userID, vf.Score); err != nil {
		utils.WebErrorMessage(c, err, "vote for article failed")
		return
	}
	utils.WebMessage(c, "vote success")
}

func ArticleScore(c *gin.Context) {
	uuid := c.Query("uuid")
	score, err := logic.ArticleScore(uuid)
	if err != nil {
		utils.WebErrorMessage(c, err, "get score failed")
		return
	}
	utils.WebMessage(c, score)
}

func HighestScoreArticle(c *gin.Context) {
	offset := c.Query("offset")
	size := c.Query("size")
	as, err := logic.HighestScoreArticle(offset, size)
	if err != nil {
		utils.WebErrorMessage(c, err, model.ErrGetList.Error())
		return
	}
	utils.WebMessage(c, as)
}
