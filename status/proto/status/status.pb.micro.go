// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/status/status.proto

package api_status

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	proto1 "github.com/micro/go-micro/v3/api/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v3/api"
	client "github.com/micro/go-micro/v3/client"
	server "github.com/micro/go-micro/v3/server"
	microClient "github.com/micro/micro/v3/service/client"
	microServer "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option
var _ = microServer.Handle
var _ = microClient.Call

// Api Endpoints for Status service

func NewStatusEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Status service

type StatusService interface {
	Call(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
}

type statusService struct {
	name string
}

func NewStatusService(name string) StatusService {
	return &statusService{name: name}
}

func (c *statusService) Call(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := microClient.NewRequest(c.name, "Status.Call", in)
	out := new(proto1.Response)
	err := microClient.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Status service

type StatusHandler interface {
	Call(context.Context, *proto1.Request, *proto1.Response) error
}

func RegisterStatusHandler(hdlr StatusHandler, opts ...server.HandlerOption) error {
	type status interface {
		Call(ctx context.Context, in *proto1.Request, out *proto1.Response) error
	}
	type Status struct {
		status
	}
	h := &statusHandler{hdlr}
	return microServer.Handle(microServer.NewHandler(&Status{h}, opts...))
}

type statusHandler struct {
	StatusHandler
}

func (h *statusHandler) Call(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.StatusHandler.Call(ctx, in, out)
}
