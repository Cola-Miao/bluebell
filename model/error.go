package model

import "errors"

var (
	ErrReadInConfig = errors.New("read in config failed")
	ErrUserExist    = errors.New("username existed")
	ErrSetCache     = errors.New("set cache failed")
	ErrGetCache     = errors.New("get cache failed")
	ErrParseForm    = errors.New("parse form failed")
	ErrParseQuery   = errors.New("parse query failed")
	ErrGetList      = errors.New("get list failed")
)
