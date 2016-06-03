package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/fasthttp"
	"github.com/ngc224/hck/config"
	"github.com/ngc224/hck/controller"
)

type Server struct {
	router     *echo.Echo
	controller *controller.Controller
}

func NewServer(c *config.Config) *Server {
	return &Server{
		router:     echo.New(),
		controller: controller.NewController(c),
	}
}

func (srv *Server) routing() {
	srv.router.GET("/", srv.controller.HealthCheck)
}

func (srv *Server) Run() {
	srv.routing()
	srv.router.Run(fasthttp.New(srv.controller.Config.Listen))
}
