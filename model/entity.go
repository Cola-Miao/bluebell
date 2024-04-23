package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID       int64          `db:"id"`
	UUID     int64          `db:"uuid"`
	Username string         `db:"username"`
	Hash     string         `db:"hash"`
	Email    sql.NullString `db:"email"`
}

type Community struct {
	ID            int64  `db:"id"`
	AdminUUID     int64  `db:"admin_uuid"`
	Administrator string `db:"administrator"`
	Name          string `db:"name" binding:"required,min=1,max=16"`
	Introduction  string `db:"introduction" binding:"required,max=512"`
}

type Article struct {
	ID           int64     `db:"id"`
	UUID         int64     `db:"uuid"`
	CommunityID  int64     `db:"community_id" json:"community_id" binding:"required"`
	AuthorUUID   int64     `db:"author_uuid"`
	Author       string    `db:"author"`
	Title        string    `db:"title" binding:"required,max=32" example:"testTitle"`
	Content      string    `db:"content" binding:"required" example:"if too long will got a truncated introduction"`
	Introduction string    `db:"introduction"`
	Score        float32   `db:"score"`
	CreateAt     time.Time `db:"create_at"`
	UpdateAt     time.Time `db:"update_at"`
}

type ArticleLite struct {
	ID           int64     `db:"id"`
	UUID         int64     `db:"uuid"`
	CommunityID  int64     `db:"community_id" json:"community_id" binding:"required"`
	AuthorUUID   int64     `db:"author_uuid"`
	Author       string    `db:"author"`
	Title        string    `db:"title" binding:"required,max=32"`
	Introduction string    `db:"introduction"`
	Score        float32   `db:"score"`
	CreateAt     time.Time `db:"create_at"`
	UpdateAt     time.Time `db:"update_at"`
}
