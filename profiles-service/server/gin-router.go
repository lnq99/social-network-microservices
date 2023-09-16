package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *GinServer) SetupRouter() {
	r := s.engine

	v1 := r.Group("api/v1")

	profile := v1.Group("profile")
	{
		profile.GET(":id", s.handlers.GetProfile)
		profile.GET("short/:id", s.handlers.GetShortProfile)
		profile.PATCH("intro", s.handlers.ChangeIntro)
	}

	rel := v1.Group("rel")
	{
		rel.GET("friends/:id", s.handlers.GetFriendsDetail)
		rel.GET("mutual-friends/:id", s.handlers.GetMutualFriends)
		rel.GET("mutual-type/:id", s.handlers.GetMutualAndType)
		rel.PUT(":id/:type", s.handlers.ChangeType)
	}

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
