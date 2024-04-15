package model

import "database/sql"

type User struct {
	ID       int64
	UUID     int64
	Username string
	Hash     string
	Email    sql.NullString
	//	TODO: Add Delete_At ?
}
