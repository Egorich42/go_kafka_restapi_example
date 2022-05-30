package api

import (
	"github.com/valyala/fasthttp"
	"github.com/Egorich42/go_kafka_restapi_example/container"
	"github.com/Egorich42/go_kafka_restapi_example/internal/service"
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
	reqBody := ctx.PostBody()
	t.service.SendCoords(reqBody)
	log.Printf("New coords send to query")
	return
}
