package logic

import (
	"bluebell/dao/msq"
	"bluebell/model"
	"bluebell/utils"
	"fmt"
	"strconv"
)

func CommunityList() ([]model.Community, error) {
	cs, err := msq.CommunityList()
	return cs, err
}

func CreateCommunity(c *model.Community) error {
	err := msq.CreateCommunity(c)
	return err
}

func FindCommunityByName(name string) (*model.Community, error) {
	c, err := msq.FindCommunityByName(name)
	return c, err
}

func CreateArticle(art *model.Article) error {
	exist, err := msq.CommunityExist(art.CommunityID)
	if err != nil {
		return fmt.Errorf("opration mysql failed: %w", err)
	}
	if !exist {
		return fmt.Errorf("community not exist")
	}

	r := []rune(art.Content)
	if len(r) > 128 {
		art.Introduction = string(r[:500]) + "..."
	} else {
		art.Introduction = string(r)
	}
	art.UUID = utils.GenerateUUID()

	err = msq.CreateArticle(art)
	return err
}

func ReadArticle(uuid int64) (*model.Article, error) {
	art, err := msq.FindArticleByUUID(uuid)
	return art, err
}

func ArticleList(offset, size string) ([]model.ArticleLite, error) {
	of, err := strconv.Atoi(offset)
	if err != nil {
		return nil, fmt.Errorf("parse query failed:%w", err)
	}
	sz, err := strconv.Atoi(size)
	if err != nil {
		return nil, fmt.Errorf("parse query failed:%w", err)
	}
	as, err := msq.ArticleList(of, sz)
	return as, err
}
