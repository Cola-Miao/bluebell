package middleware

import (
	"bluebell/utils"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := c.Get("uuid")
		if ok {
			c.Next()
			return
		}

		jwt, err := c.Cookie("jwt")
		if err != nil {
			utils.WebErrorMessage(c, err, "get cookie failed")
			return
		}

		u, newJWT, err := utils.ParseJWT(jwt)
		if err != nil {
			utils.WebErrorMessage(c, err, "parse jwt failed")
			return
		}
		if newJWT != "" {
			utils.SetJWT(c, newJWT)
		}

		c.Set("uuid", u.UUID)
		c.Set("username", u.Username)
		c.Next()
		return
	}
}
