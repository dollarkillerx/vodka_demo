/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package main

import (
	"log"
	"vodka/core/router"
	"vodka/generate"
	router2 "vodka/router"

	"github.com/dollarkillerx/vodka"
)

func main() {
	v := vodka.New()
	router := router.New()
	router2.Registry(router)
	pb.RegisterServiceServer(v.RegisterServer(), router.RegistryGRPC())

	log.Println(v.Run(":8080"))
}
