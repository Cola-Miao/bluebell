package model

type FormSignup struct {
	Username   string `binding:"required,min=1,max=8"`
	Password   string `binding:"required,min=8,max=64"`
	RePassword string `binding:"eqfield=Password"`
}

type FormLogin struct {
	Username string `binding:"required,min=1,max=8"`
	Password string `binding:"required,min=8,max=64"`
}
