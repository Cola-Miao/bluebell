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

// CommunityList godoc
//
//	@Summary		Get community list
//	@Description	get community list from mysql.bluebell.community
//	@Tags			community
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		400	{string}	string
//	@Router			/community [get]
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

// CommunityInfo godoc
//
//	@Summary		Get community information
//	@Description	get path specified community's information from mysql.bluebell.community
//	@Tags			community
//	@Accept			json
//	@Produce		json
//	@Param			community_name	path		string	true	"specified community"
//	@Success		200				{object}	string
//	@Failure		400				{string}	string
//	@Router			/community/{name} [get]
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

// CreateCommunity godoc
//
//	@Summary		Create community
//	@Description	check community_name whether to repeat
//	@Description	create a new community, default admin is creator
//	@Tags			community, auth
//	@Accept			json
//	@Produce		json
//	@Param			community	body		model.Community	true	"community name & introduction"
//	@Success		200			{object}	string
//	@Failure		400			{string}	string
//	@Router			/create_community [post]
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

// ReadArticle godoc
//
//	@Summary		Get article information
//	@Description	get path uuid specified article's information from mysql.bluebell.article
//	@Tags			community
//	@Accept			json
//	@Produce		json
//	@Param			uuid	path		string	true	"article's uuid"
//	@Success		200		{object}	string
//	@Failure		400		{string}	string
//	@Router			/article/{uuid} [get]
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

// CreateArticle godoc
//
//	@Summary		Create article
//	@Description	only need get article title and content from request body
//	@Description	add creator id, name and generate uuid, introduction to article info
//	@Description	store to mysql.bluebell.article, at this time add create_at and update_at
//	@Description	add article_uuid to zset article:time and article:score
//	@Tags			community, auth
//	@Accept			json
//	@Produce		json
//	@Param			article	body		model.Article	true	"article information"
//	@Success		200		{object}	string
//	@Failure		400		{string}	string
//	@Router			/create_article [post]
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

// ArticleList godoc
//
//	@Summary		Get article list without content
//	@Description	get article list, need input offset and (page)size in query
//	@Tags			community
//	@Accept			json
//	@Produce		json
//	@Param			offset	query		string	true	"mysql.bluebell.article offset"
//	@Param			size	query		string	true	"page size"
//	@Success		200		{object}	string
//	@Failure		400		{string}	string
//	@Router			/article [get]
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

// ArticleListByCommunity godoc
//
//	@Summary		Get specified community's article list without content
//	@Description	get article list, need input community_id in path and offset ,(page)size in query
//	@Tags			community
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string	true	"community id"
//	@Param			offset	query		string	true	"mysql.bluebell.article offset"
//	@Param			size	query		string	true	"page size"
//	@Success		200		{object}	string
//	@Failure		400		{string}	string
//	@Router			/community/{id} [get]
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

// VoteForArticle godoc
//
//	@Summary		vote form article
//	@Description	vote 0~5 score for article
//	@Description	voter will store to zset article:voter:{article_uuid}
//	@Description	repeat vote be checked and got correct score
//	@Description	vote for expired article will be refuse
//	@Tags			community, auth
//	@Accept			json
//	@Produce		json
//	@Param			voteForm	body		model.FormVote	true	"article_uuid and score "
//	@Success		200			{object}	string
//	@Failure		400			{string}	string
//	@Router			/article_vote [post]
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
