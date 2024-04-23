package model

type FormSignup struct {
	Username   string `binding:"required,min=1,max=8" example:"testName"`
	Password   string `binding:"required,min=8,max=64" example:"testPassword"`
	RePassword string `binding:"eqfield=Password"  example:"testPassword"`
}

type FormLogin struct {
	Username string `binding:"required,min=1,max=8" example:"testName"`
	Password string `binding:"required,min=8,max=64" example:"testPassword"`
}

type FormVote struct {
	UUID  int64   `binding:"required" example:"1781231541096022016"`
	Score float64 `binding:"required,gte=1,lte=5" example:"4"`
}
