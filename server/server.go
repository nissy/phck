package server

import (
	"github.com/facebookgo/grace/gracehttp"
	"github.com/facebookgo/pidfile"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
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

func (srv *Server) Run() error {
	srv.routing()

	if len(srv.controller.Config.PIDFile) > 0 {
		pidfile.SetPidfilePath(srv.controller.Config.PIDFile)

		if err := pidfile.Write(); err != nil {
			return err
		}
	}

	std := standard.New(srv.controller.Config.Listen)
	std.SetHandler(srv.router)
	if err := gracehttp.Serve(std.Server); err != nil {
		return err
	}

	return nil
}
