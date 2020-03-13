
/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package middleware

import (
	"time"
	"vodka/core/router"

	middleware2 "github.com/dollarkillerx/vodka/middleware"
)

// 基础Prometheus 统计 1.请求数量2.请求错误3.请求耗时分布
func BasePrometheus(ctx *router.RouterContext) {
	startTime := time.Now()
	msg := ctx.GetPrometheusMsg()
	middleware2.Prometheus.IncrRequest(ctx.Context,msg.ServerName,msg.FuncName)
	ctx.Next()
	err := ctx.ErrGet()
	if err != nil {
		middleware2.Prometheus.IncrCode(ctx.Context,msg.ServerName,msg.FuncName,err)
	}
	middleware2.Prometheus.Latency(ctx.Context, msg.ServerName,
		msg.FuncName, time.Since(startTime).Nanoseconds()/1000)
}
