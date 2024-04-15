package logic

import (
	"bluebell/dao/msq"
	"bluebell/model"
	"bluebell/utils"
)

func Signup(sf *model.FormSignup) (err error) {
	exist, err := msq.UserExist(sf.Username)
	if err != nil {
		return err
	}
	if exist {
		return model.ErrUserExist
	}

	hash, err := utils.Password(sf.Password).Encode()
	if err != nil {
		return err
	}

	u := model.User{
		UUID:     utils.GenerateUUID(),
		Username: sf.Username,
		Hash:     hash,
	}
	if err = msq.CreateUser(&u); err != nil {
		return err
	}
	return nil
}
