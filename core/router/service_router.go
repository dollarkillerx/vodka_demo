/**
*@Program: vodka
*@MicroServices Framework: https://github.com/dollarkillerx
 */
package router

import (
	"context"
	"log"
	pb "vodka/generate"
)

type Router struct {
	router *serviceRouter
}

func New() *Router {
	return &Router{
		router: &serviceRouter{

			Run1FuncSlice: make([]RunFunc, 0),

			Run2FuncSlice: make([]RunFunc, 0),

			Run3FuncSlice: make([]RunFunc, 0),

			Run4FuncSlice: make([]RunFunc, 0),
		},
	}
}

func (r *Router) RegistryGRPC() *serviceRouter {
	return r.router
}

func (r *Router) Run1(Run1func ...RunFunc) {
	r.router.Run1FuncSlice = append(r.router.Run1FuncSlice, Run1func...)
}

func (r *Router) Run2(Run2func ...RunFunc) {
	r.router.Run2FuncSlice = append(r.router.Run2FuncSlice, Run2func...)
}

func (r *Router) Run3(Run3func ...RunFunc) {
	r.router.Run3FuncSlice = append(r.router.Run3FuncSlice, Run3func...)
}

func (r *Router) Run4(Run4func ...RunFunc) {
	r.router.Run4FuncSlice = append(r.router.Run4FuncSlice, Run4func...)
}

type serviceRouter struct {
	Run1FuncSlice []RunFunc

	Run2FuncSlice []RunFunc

	Run3FuncSlice []RunFunc

	Run4FuncSlice []RunFunc
}

type RouterContextItem interface {
	_routerContext()
}

type RouterContext struct {
	Ctx      RouterContextItem
	funcList []RunFunc
	index    int
	psg      *PrometheusMsg
}

type PrometheusMsg struct {
	FuncName    string // 方法名称
	ServerName  string // 服务名称
	Environment string // 环境 开发 or 测试
	Cluster     string // 集群名称
	EngineRoom  string // 机房
	TraceId     string // 分布式追踪id
	RespIP      string // 服务端IP
	ReqIP       string // 客户端IP
}

func (r *RouterContext) GetPrometheusMsg() *PrometheusMsg {
	return r.psg
}

func (r *RouterContext) Next() {
	r.index++
	if r.index <= len(r.funcList) {
		r.funcList[r.index-1](r)
	} else {
		log.Println("RouterContext Next  what ???")
	}
}

type Run1FuncContext struct {
	Ctx  context.Context
	Req  *pb.Req
	Resp *pb.Resp

	Err error
}

type Run2FuncContext struct {
	Req *pb.Req
	Ser pb.Service_Run2Server

	Err error
}

type Run3FuncContext struct {
	Ser pb.Service_Run3Server

	Err error
}

type Run4FuncContext struct {
	Ser pb.Service_Run4Server

	Err error
}

func (r *Run1FuncContext) _routerContext() {}

func (r *Run2FuncContext) _routerContext() {}

func (r *Run3FuncContext) _routerContext() {}

func (r *Run4FuncContext) _routerContext() {}

type RunFunc func(ctx *RouterContext)

// 下面是主题方法

func (s *serviceRouter) Run1(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	routerContext := RouterContext{
		Ctx: &Run1FuncContext{
			Ctx:  ctx,
			Req:  req,
			Resp: nil,
			Err:  nil,
		},
		funcList: s.Run1FuncSlice,
		index:    0,
		psg: &PrometheusMsg{
			FuncName: "Run1",
		},
	}

	routerContext.Next()
	funcContext := routerContext.Ctx.(*Run1FuncContext)
	return funcContext.Resp, funcContext.Err
}

func (s *serviceRouter) Run2(req *pb.Req, ser pb.Service_Run2Server) error {
	routerContext := RouterContext{
		Ctx: &Run2FuncContext{
			Req: req,
			Ser: ser,
			Err: nil,
		},
		funcList: s.Run2FuncSlice,
		index:    0,
	}

	routerContext.Next()
	funcContext := routerContext.Ctx.(*Run2FuncContext)
	return funcContext.Err
}

func (s *serviceRouter) Run3(ser pb.Service_Run3Server) error {
	routerContext := RouterContext{
		Ctx: &Run3FuncContext{
			Err: nil,
			Ser: ser,
		},
		funcList: s.Run3FuncSlice,
		index:    0,
	}

	routerContext.Next()
	funcContext := routerContext.Ctx.(*Run3FuncContext)
	return funcContext.Err
}

func (s *serviceRouter) Run4(ser pb.Service_Run4Server) error {
	routerContext := RouterContext{
		Ctx: &Run4FuncContext{
			Err: nil,
			Ser: ser,
		},
		funcList: s.Run4FuncSlice,
		index:    0,
	}

	routerContext.Next()
	funcContext := routerContext.Ctx.(*Run4FuncContext)
	return funcContext.Err
}
