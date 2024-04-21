package logic

import (
	"bluebell/dao/msq"
	"bluebell/dao/rdb"
	"bluebell/model"
	"bluebell/utils"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
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

func ArticleListByCommunity(communityID, offset, size string) ([]model.ArticleLite, error) {
	comID, err := strconv.Atoi(communityID)
	if err != nil {
		return nil, fmt.Errorf("parse query failed:%w", err)
	}
	of, err := strconv.Atoi(offset)
	if err != nil {
		return nil, fmt.Errorf("parse query failed:%w", err)
	}
	sz, err := strconv.Atoi(size)
	if err != nil {
		return nil, fmt.Errorf("parse query failed:%w", err)
	}
	as, err := msq.ArticleListByCommunity(comID, of, sz)
	return as, err
}

func VoteForArticle(artID, userID int64, score float64) error {
	as := strconv.Itoa(int(artID))
	us := strconv.Itoa(int(userID))
	hot, err := rdb.HotArticle(as)
	if err != nil && !errors.Is(err, redis.Nil) {
		return fmt.Errorf("query hot article failed: %w", err)
	}
	if !hot {
		return errors.New("the article has expired ")
	}

	if err = rdb.VoteForArticle(as, us, score); err != nil {
		return fmt.Errorf("vote for article failed: %w", err)
	}
	return nil
}

func ArticleScore(uuid string) (float64, error) {
	score, err := rdb.ArticleScore(uuid)
	if err != nil {
		return 0, fmt.Errorf("get score failed: %w", err)
	}
	return score, nil
}
