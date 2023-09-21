package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-chi/jwtauth/v5"
	"net/http"
)

func AuthMiddleware(tokenAuth *jwtauth.JWTAuth) gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenStr, err := c.Cookie("jwt")

		if err == nil {
			token, err := tokenAuth.Decode(tokenStr)
			//fmt.Printf("%+v\n", token)

			if err == nil {
				idStr, ok := token.Get("ID")
				if ok {
					c.Set("ID", int(idStr.(float64)))
				}
			}
		}

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}
