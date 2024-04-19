package logic

import (
	"bluebell/dao/msq"
	"bluebell/model"
	"bluebell/utils"
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
	// TODO :check community exist
	r := []rune(art.Content)
	if len(r) > 128 {
		art.Introduction = string(r[:500]) + "..."
	} else {
		art.Introduction = string(r)
	}
	art.UUID = utils.GenerateUUID()

	err := msq.CreateArticle(art)
	return err
}

func ReadArticle(uuid int64) (*model.Article, error) {
	art, err := msq.FindArticleByUUID(uuid)
	return art, err
}
