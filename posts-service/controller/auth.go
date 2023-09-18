package controller

//func AuthMiddleware(tokenAuth *jwtauth.JWTAuth) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		id := 0
//
//		var claims map[string]interface{}
//
//		token, err := c.Cookie("jwt")
//
//		if err == nil {
//			json.Unmarshal([]byte(token), &claims)
//
//			id, err = auth.ParseTokenId(token)
//			if err == nil {
//				c.Set("ID", id)
//				// log.Println(id, token, err)
//				return
//			}
//		}
//
//		err = auth.TokenValid(c.Request)
//		// log.Println(err)
//
//		if err != nil {
//			c.AbortWithStatus(http.StatusUnauthorized)
//		}
//	}
//}
