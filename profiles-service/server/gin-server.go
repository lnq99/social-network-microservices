package server

import (
	"app/config"
	"app/controller"
	"app/service"
	"app/util"

	limits "github.com/gin-contrib/size"
	"github.com/gin-gonic/gin"
)

type GinServer struct {
	util.BaseServer
	engine   *gin.Engine
	handlers *controller.Controller
	config   *config.ServerConfig
}

func NewGinServer(service *service.Service, cfg *config.ServerConfig) util.Server {
	// engine.SetMode(engine.ReleaseMode)
	engine := gin.New()

	// Middleware
	engine.Use(gin.Logger())
	engine.Use(gin.Recovery())
	engine.Use(limits.RequestSizeLimiter(1000))

	ctrl := controller.NewGinController(service)
	s := GinServer{
		engine:   engine,
		handlers: ctrl,
		config:   cfg,
	}
	s.SetupRouter()
	return &s
}

func (s *GinServer) Run() {
	cfg := util.DefaultConfig()
	cfg.Addr = s.config.Host + ":" + s.config.Port
	cfg.Handler = s.engine
	s.InitHttpServer(cfg)
	s.BaseServer.Run()
}
