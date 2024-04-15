package model

import "errors"

var (
	ErrReadInConfig = errors.New("read in config failed")
	ErrUserExist    = errors.New("username existed")
)
