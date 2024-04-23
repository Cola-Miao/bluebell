package service

import (
	"bluebell/dao/msq"
	"bluebell/logic"
	"bluebell/model"
	"bluebell/utils"
	"github.com/gin-gonic/gin"
)

// Signup godoc
//
//	@Summary		User Signup
//	@Description	get signup info json from request body
//	@Description	validate username existed, the length of data and re-password
//	@Description	create new user with encode password to mysql.bluebell.user
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			SignupForm	body		model.FormSignup	true	"User Signup Form"
//	@Success		200			{object}	string				"signup successful"
//	@Failure		400			{string}	string
//	@Router			/signup [post]
func Signup(c *gin.Context) {
	var sf model.FormSignup
	var err error
	if err = c.ShouldBindJSON(&sf); err != nil {
		utils.WebErrorMessage(c, err, "incorrect input")
		return
	}
	if err = logic.Signup(&sf); err != nil {
		utils.WebErrorMessage(c, err, "signup failed")
		return
	}
	utils.WebMessage(c, "signup successful")
}

// Login godoc
//
//	@Summary		User Login
//	@Description	get login info json from request body
//	@Description	validate username existed, the length of data and compare password
//	@Description	generate and set jwt to cookie
//	@Tags			account
//	@Accept			json
//	@Produce		json
//	@Param			LoginForm	body		model.FormLogin	true	"User Login Form"
//	@Success		200			{object}	string			"login successful"
//	@Failure		400			{string}	string
//	@Router			/login [post]
func Login(c *gin.Context) {
	var lf model.FormLogin
	var err error
	if err = c.ShouldBindJSON(&lf); err != nil {
		utils.WebErrorMessage(c, err, "incorrect input")
		return
	}
	u, err := msq.FindUserByName(lf.Username)
	if err != nil {
		utils.WebErrorMessage(c, err, "user not exist")
		return
	}
	if err = utils.Password(lf.Password).Compare(u.Hash); err != nil {
		utils.WebErrorMessage(c, err, "wrong password")
		return
	}
	token, err := utils.GenerateJWT(u)
	if err != nil {
		utils.WebErrorMessage(c, err, "generate jwt failed")
		return
	}
	utils.SetJWT(c, token)
	utils.WebMessage(c, "login successful")
}
