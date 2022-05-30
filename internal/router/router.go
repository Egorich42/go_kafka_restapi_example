package router

import (
	"github.com/fasthttp/router"
	"github.com/Egorich42/go_kafka_restapi_example/internal/api"
)

type AppRouter struct {
	controller *api.AppController
	router     *router.Group
}

func NewAppRouter(c *api.AppController, r *router.Group) *AppRouter {
	return &AppRouter{
		controller: c,
		router:     r,
	}
}

func (a *AppRouter) Register() {
	a.router.POST("/coordinates", a.controller.SendCoords)
}
