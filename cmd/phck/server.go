package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/nissy/bon"
	"github.com/nissy/bon/middleware"
	"github.com/nissy/phck/controller"
)

type server struct {
	router *bon.Mux
}

func newRouter(cfg *config) *bon.Mux {
	r := bon.NewRouter()

	health := controller.NewHealth(cfg.Health)
	r.Use(middleware.Timeout(10 * time.Second))
	r.Get("/", health.HealthCheck)

	return r
}

func newServer(cfg *config) *server {
	return &server{
		router: newRouter(cfg),
	}
}

func (srv *server) writePIDFIle(filepath string) error {
	return ioutil.WriteFile(filepath, []byte(strconv.Itoa(os.Getpid())), os.ModePerm)
}

func (srv *server) listenAndServe(listen string) error {
	return gracehttp.Serve(&http.Server{
		Addr:    listen,
		Handler: srv.router,
	})
}
