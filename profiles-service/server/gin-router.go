package server

import (
	"app/controller"
	"github.com/go-chi/jwtauth/v5"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *GinServer) SetupRouter() {
	r := s.engine

	tokenAuth := jwtauth.New("HS256", []byte(s.config.Auth.JwtSigningKey), nil)

	api := r.Group("api/v1", controller.AuthMiddleware(tokenAuth))

	profile := api.Group("profile")
	{
		profile.GET(":id", s.handlers.GetProfile)
		profile.GET("short/:id", s.handlers.GetShortProfile)
		profile.PATCH("intro", s.handlers.ChangeIntro)
	}

	rel := api.Group("rel")
	{
		rel.GET("friends/:id", s.handlers.GetFriendsDetail)
		rel.GET("mutual-friends/:id", s.handlers.GetMutualFriends)
		rel.GET("mutual-type/:id", s.handlers.GetMutualAndType)
		rel.PUT(":id/:type", s.handlers.ChangeType)
	}

	api.GET("search", s.handlers.Search)

	//apiV1.GET("ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run(":80")

	r.GET("/manage/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
