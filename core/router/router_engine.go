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
	"errors"
)

type Router struct {
	router *serviceRouter
}

func New() *Router {
	return &Router{}
}

func (r *Router) RegistryGRPC() *serviceRouter {
	return r.router
}

func (r *Router) Run1(run1func ...Run1Func) {
	r.router.run1FuncSlice = append(r.router.run1FuncSlice, run1func...)
}

func (r *Router) Run2(run2func ...Run2Func) {
	r.router.run2FuncSlice = append(r.router.run2FuncSlice, run2func...)
}

func (r *Router) Run3(run3func ...Run3Func) {
	r.router.run3FuncSlice = append(r.router.run3FuncSlice, run3func...)
}

func (r *Router) Run4(run4func ...Run4Func) {
	r.router.run4FuncSlice = append(r.router.run4FuncSlice, run4func...)
}

func (r *Router) Run1Next(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	r.router.run1index++
	if r.router.run1index >= len(r.router.run1FuncSlice) {
		return nil,errors.New("next number of methods exceeded")
	}
	return r.router.run1FuncSlice[r.router.run1index](ctx,req)
}

type serviceRouter struct {
	run1index     int
	run2index     int
	run3index     int
	run4index     int
	run1FuncSlice []Run1Func
	run2FuncSlice []Run2Func
	run3FuncSlice []Run3Func
	run4FuncSlice []Run4Func
}



type Run1Func func(ctx context.Context, req *pb.Req) (*pb.Resp, error)

type Run2Func func(req *pb.Req, ser pb.Service_Run2Server) error

type Run3Func func(ser pb.Service_Run3Server) error

type Run4Func func(ser pb.Service_Run4Server) error

func (s *serviceRouter) Run1(ctx context.Context, req *pb.Req) (*pb.Resp, error) {
	return s.run1FuncSlice[0](ctx, req)
}

func (s *serviceRouter) Run2(req *pb.Req, ser pb.Service_Run2Server) error {
	return s.run2FuncSlice[0](req, ser)
}

func (s *serviceRouter) Run3(ser pb.Service_Run3Server) error {
	return s.run3FuncSlice[0](ser)
}

func (s *serviceRouter) Run4(ser pb.Service_Run4Server) error {
	return s.run4FuncSlice[0](ser)
}
