/**
*@program: vodka_demo
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 14:43
 */
package controller

import (
	"awesome/core/router"
	pb "awesome/generate"
	"fmt"
	"time"
)

func Run1(ctx *router.RouterContext) {
	startTime := time.Now().UnixNano()

	ctx.Next()

	endTime := time.Now().UnixNano()

	seconds := float64((endTime - startTime) / 1e9)
	fmt.Println("Seconds: ", seconds)
}

func Run2(ctx *router.RouterContext) {
	context := ctx.Ctx.(*router.Run1FuncContext)
	fmt.Println(context.Req)
	context.Resp = &pb.Resp{
		Msg: "hello",
	}
}
