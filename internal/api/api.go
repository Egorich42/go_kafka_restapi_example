package api

import (
	"github.com/valyala/fasthttp"
	"github.com/Egorich42/testserver/container"
	"github.com/Egorich42/testserver/internal/service"
	"log"
)

type AppController struct {
	container container.Container
	service   service.AppService
}

func NewAppController(container container.Container, service service.AppService) *AppController {
	return &AppController{
		container: container,
		service:   service,
	}
}

func (t *AppController) SendCoords(ctx *fasthttp.RequestCtx) {
	log.Print("i catch location")
	reqBody := ctx.PostBody()
	t.service.SendCoords(reqBody)
	log.Print(reqBody)
	return
}
