
/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package main

import (
	"log"
	"vodka/core/router"
	"vodka/generate"
	middleware2 "vodka/middleware"
	router2 "vodka/router"
	
	"github.com/dollarkillerx/vodka"
	"github.com/dollarkillerx/vodka/middleware"
)

func main() {
	v := vodka.New()
	router.ServerAddr = ":8080"
	app := router.New()
	app.Use(middleware2.BasePrometheus)  // 注册全局中间件  基础Prometheus
	router2.Registry(app)
	pb.RegisterServiceServer(v.RegisterServer(), app.RegistryGRPC())

	go middleware.Prometheus.Run(":8085")
	log.Println(v.Run(router.ServerAddr))
}
