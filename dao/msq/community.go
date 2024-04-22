package msq

import (
	"bluebell/model"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func CommunityList() ([]model.Community, error) {
	var cs []model.Community
	query := "SELECT * FROM community"
	err := db.Select(&cs, query)
	return cs, err
}

func FindCommunityByName(name string) (*model.Community, error) {
	var c model.Community
	query := "SELECT * FROM community WHERE name = ? LIMIT 1"
	err := db.Get(&c, query, name)
	return &c, err
}

func CreateCommunity(c *model.Community) error {
	query := `INSERT INTO community(
              name,admin_uuid,introduction,administrator
			  )VALUES (?,?,?,?)`
	_, err := db.Exec(query, c.Name, c.AdminUUID, c.Introduction, c.Administrator)
	return err
}

func FindArticleByUUID(uuid int64) (*model.Article, error) {
	var art model.Article
	query := "SELECT * FROM article WHERE uuid = ? LIMIT 1"
	err := db.Get(&art, query, uuid)
	return &art, err
}

func CreateArticle(art *model.Article) error {
	query := `INSERT INTO article(
               uuid, community_id, author_uuid, author, title, content, introduction
			  ) VALUES (?,?,?,?,?,?,?)`
	_, err := db.Exec(query, art.UUID, art.CommunityID, art.AuthorUUID, art.Author, art.Title, art.Content, art.Introduction)
	return err
}

func CommunityExist(id int64) (exist bool, err error) {
	var res int
	query := "SELECT 1 FROM community WHERE id = ? LIMIT 1"
	if err = db.Get(&res, query, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = nil
		}
		return exist, err
	}
	return true, nil
}

func ArticleList(offset, size int) ([]model.ArticleLite, error) {
	var as []model.ArticleLite
	query := `SELECT id, uuid, community_id, author_uuid, author, title, introduction, create_at, update_at 
	FROM article LIMIT ?,?`
	err := db.Select(&as, query, offset, size)
	return as, err
}

func ArticleListByCommunity(communityID, offset, size int, desc bool) ([]model.ArticleLite, error) {
	var as []model.ArticleLite
	query := `SELECT id, uuid, community_id, author_uuid, author, title, introduction,score, create_at, update_at 
	FROM article WHERE community_id = ?`
	if desc {
		query += " ORDER BY uuid DESC LIMIT ?,?"
	} else {
		query += " LIMIT ?,?"
	}

	err := db.Select(&as, query, communityID, offset, size)
	return as, err
}

func ArticleListByUUID(uuid []int64) ([]model.ArticleLite, error) {
	var as []model.ArticleLite
	query, args, err := sqlx.In(`SELECT id, uuid, community_id, author_uuid, author, title, introduction, create_at, update_at  
	FROM article WHERE uuid IN (?)`, uuid)
	if err != nil {
		return nil, fmt.Errorf("parse slice failed: %w", err)
	}
	query = db.Rebind(query)
	err = db.Select(&as, query, args...)
	return as, err
}

func UpdateArticleScore(uuid int64, score float64) error {
	query := `UPDATE article SET score = ? WHERE uuid = ?`
	_, err := db.Exec(query, score, uuid)
	return err
}
