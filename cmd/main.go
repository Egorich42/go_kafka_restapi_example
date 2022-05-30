package main

import (
	fastHttpRouter "github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
	"github.com/Egorich42/go_kafka_restapi_example/config"
	"github.com/Egorich42/go_kafka_restapi_example/container"
	"github.com/Egorich42/go_kafka_restapi_example/internal/api"
	"github.com/Egorich42/go_kafka_restapi_example/internal/router"
	"github.com/Egorich42/go_kafka_restapi_example/internal/service"
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

	// Define consumer conf

	zookeeperAddr := c.GetConfig().ZookeeperHost+":"+c.GetConfig().ZookeeperPort
	
	cgroup := "cgroup"
	topic := "coordinates"
	log.Print(cgroup, topic, zookeeperAddr, "@@@@@@@@@@@@@@@@@")
	go service.StartConsume(cgroup, topic, zookeeperAddr)

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
