package msq

import "bluebell/model"

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
