package server

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/ngc224/phck/config"
	"github.com/ngc224/phck/controller"
)

type Server struct {
	router     *echo.Echo
	config     *config.Config
	controller *controller.Controller
}

func NewServer(c *config.Config) *Server {
	return &Server{
		router:     echo.New(),
		config:     c,
		controller: controller.NewController(c),
	}
}

func (srv *Server) routing() {
	srv.router.GET("/", srv.controller.HealthCheck)
}

func WritePIDFile(filepath string) error {
	return ioutil.WriteFile(filepath, []byte(strconv.Itoa(os.Getpid())), os.ModePerm)
}

func (srv *Server) Run() error {
	srv.routing()

	if len(srv.config.PIDFile) > 0 {
		if err := WritePIDFile(srv.config.PIDFile); err != nil {
			return err
		}
	}

	std := standard.New(srv.config.Listen)
	std.SetHandler(srv.router)
	if err := gracehttp.Serve(std.Server); err != nil {
		return err
	}

	return nil
}
