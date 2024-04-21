package dao

import (
	"bluebell/dao/msq"
	"bluebell/dao/rdb"
	"fmt"
	"strconv"
)

func ArticleScoreToCold(uuid int64) error {
	us := strconv.Itoa(int(uuid))
	score, err := rdb.ArticleScore(us)
	if err != nil {
		return fmt.Errorf("get article score failed: %w", err)
	}
	if err = msq.UpdateArticleScore(uuid, score); err != nil {
		return fmt.Errorf("update article score failed: %w", err)
	}
	if err = rdb.DeleteArticle(us); err != nil {
		return fmt.Errorf("delete cache failed: %w", err)
	}
	return nil
}
