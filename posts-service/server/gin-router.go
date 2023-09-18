package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//var tokenAuth = jwtauth.New("HS256", []byte(cfg.Auth.JwtSigningKey), nil)

func (s *GinServer) SetupRouter() {
	r := s.engine

	api := r.Group("api/v1")
	//api := r.Group("api/v1", controller.AuthMiddleware(ctrl.auth))
	//handler := jwtauth.Verifier(tokenAuth)
	//api.Use(gin.WrapH())

	post := api.Group("post")
	{
		post.GET(":id", s.handlers.GetPost)
		post.GET("u/:id", s.handlers.GetPostByUserId)
		post.POST("", s.handlers.PostPost)
		post.DELETE(":id", s.handlers.DeletePost)
	}

	react := api.Group("react")
	{
		react.GET(":post_id", s.handlers.GetReaction)
		react.GET("u/:u_id", s.handlers.GetReactionByUserPost)
		react.PUT(":post_id/:type", s.handlers.PutReaction)
	}

	cmt := api.Group("cmt")
	{
		cmt.GET(":id", s.handlers.GetTreeComment)
		cmt.POST("", s.handlers.PostComment)
	}

	photo := api.Group("photo")
	{
		photo.GET(":id", s.handlers.GetPhoto)
		photo.GET("u/:id", s.handlers.GetPhotoByUserId)
	}

	r.GET("/manage/health", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
}
