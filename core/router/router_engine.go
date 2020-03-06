/**
*@program: awesomeProject
*@description: https://github.com/dollarkillerx
*@author: dollarkiller [dollarkiller@dollarkiller.com]
*@create: 2020-03-06 13:51
 */
package router

import (
	pb "awesome/generate"
	"context"
	"log"
)

type Router struct {
	router *serviceRouter
}

func New() *Router {
	return &Router{
		router: &serviceRouter{
			run1FuncSlice: make([]RunFunc, 0),
			run2FuncSlice: make([]RunFunc, 0),
			run3FuncSlice: make([]RunFunc, 0),
			run4FuncSlice: make([]RunFunc, 0),
		},
	}
}

func (r *Router) RegistryGRPC() *serviceRouter {
	return r.router
}

func (r *Router) Run1(run1func ...RunFunc) {
	r.router.run1FuncSlice = append(r.router.run1FuncSlice, run1func...)
}

func (r *Router) Run2(run2func ...RunFunc) {
	r.router.run2FuncSlice = append(r.router.run2FuncSlice, run2func...)
}

func (r *Router) Run3(run3func ...RunFunc) {
	r.router.run3FuncSlice = append(r.router.run3FuncSlice, run3func...)
}

func (r *Router) Run4(run4func ...RunFunc) {
	r.router.run4FuncSlice = append(r.router.run4FuncSlice, run4func...)
}

type serviceRouter struct {
	run1FuncSlice []RunFunc
	run2FuncSlice []RunFunc
	run3FuncSlice []RunFunc
	run4FuncSlice []RunFunc
}

type RouterContextItem interface {
	_routerContext()
}

type RouterContext struct {
	Ctx      RouterContextItem
	funcList []RunFunc
	index    int
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
	Err  error
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

func (s *serviceRouter) Run1(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	routerContext := RouterContext{
		Ctx: &Run1FuncContext{
			Ctx:  ctx,
			Req:  req,
			Resp: nil,
			Err:  nil,
		},
		funcList: s.run1FuncSlice,
		index:    0,
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
		funcList: s.run2FuncSlice,
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
		funcList: s.run3FuncSlice,
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
		funcList: s.run4FuncSlice,
		index:    0,
	}

	routerContext.Next()
	funcContext := routerContext.Ctx.(*Run4FuncContext)
	return funcContext.Err
}
