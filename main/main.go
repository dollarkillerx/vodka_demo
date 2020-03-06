/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package main

import (
	"awesome/core/router"
	"awesome/generate"
	router2 "awesome/router"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	router := router.New()
	router2.Registry(router)                                // 路由注册
	pb.RegisterServiceServer(server, router.RegistryGRPC()) // 路由注册到GRPC处
	dial, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
	server.Serve(dial)
}
