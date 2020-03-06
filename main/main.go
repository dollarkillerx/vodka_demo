/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package main

import (
	"awesome/core/router"
	"awesome/generate"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()
	router := router.New()
	pb.RegisterServiceServer(server, router.Registry())
	dial, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalln(err)
	}
	server.Serve(dial)
}
