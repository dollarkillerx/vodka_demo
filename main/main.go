
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
	"github.com/dollarkillerx/vodka/server"
)

func main() {
	v := vodka.New()
	router.ServerAddr = server.Config.Addr
	app := router.New()
	app.Use(middleware2.BasePrometheus)  // 注册全局中间件  基础Prometheus
	router2.Registry(app)
	pb.RegisterServiceServer(v.RegisterServer(), app.RegistryGRPC())

	if server.Config.Prometheus.SwitchOn {
		go middleware.Prometheus.Run(server.Config.Prometheus.Addr)
	}
	log.Println(v.Run(router.ServerAddr))
}
