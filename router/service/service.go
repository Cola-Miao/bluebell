package service

import (
	"bluebell/dao/msq"
	"bluebell/logic"
	"bluebell/model"
	"bluebell/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "ok",
	})
}

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
	utils.WebMessage(c, "login successful")
}

func TestFunc(c *gin.Context) {
	data := c.Query("data")
	res, err := msq.UserExist(data)

	c.JSON(http.StatusOK, gin.H{
		"result": res,
		"error":  err,
	})
}
