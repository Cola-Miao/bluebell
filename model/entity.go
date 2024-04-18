package model

import "database/sql"

type User struct {
	ID       int64
	UUID     int64
	Username string
	Hash     string
	Email    sql.NullString
}

type Community struct {
	ID            int64
	AdminUUID     int64
	Administrator string
	Name          string
	Introduction  string
}
