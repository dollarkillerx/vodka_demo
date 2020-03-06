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
)

type router struct {
	router *serviceRouter
}

func New() *router {
	return &router{}
}

func (r *router) Registry() *serviceRouter {
	return r.router
}

func (r *router) Run1(run1func ...Run1Func) {
	r.router.run1FuncSlice = append(r.router.run1FuncSlice, run1func...)
}

func (r *router) Run2(run2func ...Run2Func) {
	r.router.run2FuncSlice = append(r.router.run2FuncSlice, run2func...)
}

func (r *router) Run3(run3func ...Run3Func) {
	r.router.run3FuncSlice = append(r.router.run3FuncSlice, run3func...)
}

func (r *router) Run4(run4func ...Run4Func) {
	r.router.run4FuncSlice = append(r.router.run4FuncSlice, run4func...)
}

func (r *router) Run1Next() {
	r.router.run1index++
}

func (r *router) Run2Next() {
	r.router.run2index++
}

func (r *router) Run3Next() {
	r.router.run3index++
}

func (r *router) Run4Next() {
	r.router.run4index++
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

func (s *serviceRouter) Run1(ctx context.Context, req *pb.Req) (resp *pb.Resp, err error) {
	for i := 0; i < s.run1index; i++ {
		resp, err = s.run1FuncSlice[i](ctx, req)
	}
	return
}

func (s *serviceRouter) Run2(req *pb.Req, ser pb.Service_Run2Server) (err error) {
	for i := 0; i < s.run2index; i++ {
		err = s.run2FuncSlice[i](req, ser)
	}
	return
}

func (s *serviceRouter) Run3(ser pb.Service_Run3Server) (err error) {
	for i := 0; i < s.run3index; i++ {
		err = s.run3FuncSlice[i](ser)
	}
	return
}

func (s *serviceRouter) Run4(ser pb.Service_Run4Server) (err error) {
	for i := 0; i < s.run4index; i++ {
		err = s.run4FuncSlice[i](ser)
	}
	return
}
