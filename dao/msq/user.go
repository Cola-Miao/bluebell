package msq

import (
	"bluebell/model"
	"database/sql"
	"errors"
)

func CreateUser(u *model.User) error {
	query := `INSERT INTO user(
                uuid, username, hash
                )VALUES (?,?,?)`
	_, err := db.Exec(query, u.UUID, u.Username, u.Hash)
	return err
}

func UserExist(name string) (exist bool, err error) {
	var res int
	query := "SELECT 1 FROM user WHERE username = ?"
	if err = db.Get(&res, query, name); err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return true, err
}

func FindUserByName(name string) (*model.User, error) {
	var u model.User
	query := "SELECT * FROM user WHERE username = ?"
	if err := db.Get(&u, query, name); err != nil {
		return nil, err
	}
	return &u, nil
}
