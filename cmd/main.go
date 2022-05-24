package main

import (
	fastHttpRouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/Egorich42/testserver/config"
	"github.com/Egorich42/testserver/container"
	"github.com/Egorich42/testserver/internal/api"
	"github.com/Egorich42/testserver/internal/router"
	"github.com/Egorich42/testserver/internal/service"
	"log"
)

func main() {
	// Define new config
	initiatorConfig, err := config.NewAppConfig()
	if err != nil {
		log.Fatal("Unable to read envs")
	}

	// Define new container
	c := container.NewContainer(initiatorConfig)

	// Define router
	r := fastHttpRouter.New()
	apiGroup := r.Group("/v1")
	{
		initiatorService := service.NewAppService(c)
		initiatorController := api.NewAppController(c, initiatorService)
		initiatorRouter := router.NewAppRouter(initiatorController, apiGroup)
		initiatorRouter.Register()
		log.Print("Finished registering Transaction Initiator services")
	}

	addr := c.GetConfig().Host+":"+c.GetConfig().Port
	log.Fatal(fasthttp.ListenAndServe(addr, r.Handler))
}
