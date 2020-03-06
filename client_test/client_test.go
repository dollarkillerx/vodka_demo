/**
*@program: vodka_demo
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 16:49
 */
package client_test

import (
	pb "awesome/generate"
	"context"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestConnGRPC(t *testing.T) {
	dial, err := grpc.Dial("0.0.0.0:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer dial.Close()
	client := pb.NewServiceClient(dial)
	run1, err := client.Run1(context.TODO(), &pb.Req{Msg: "hello"})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(run1)
}
