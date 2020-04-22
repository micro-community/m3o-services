// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: github.com/micro/services/m3o/api/proto/m3o.proto

package go_micro_api_m3o

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for AccountService service

func NewAccountServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for AccountService service

type AccountService interface {
	Read(ctx context.Context, in *ReadAccountRequest, opts ...client.CallOption) (*ReadAccountResponse, error)
}

type accountService struct {
	c    client.Client
	name string
}

func NewAccountService(name string, c client.Client) AccountService {
	return &accountService{
		c:    c,
		name: name,
	}
}

func (c *accountService) Read(ctx context.Context, in *ReadAccountRequest, opts ...client.CallOption) (*ReadAccountResponse, error) {
	req := c.c.NewRequest(c.name, "AccountService.Read", in)
	out := new(ReadAccountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for AccountService service

type AccountServiceHandler interface {
	Read(context.Context, *ReadAccountRequest, *ReadAccountResponse) error
}

func RegisterAccountServiceHandler(s server.Server, hdlr AccountServiceHandler, opts ...server.HandlerOption) error {
	type accountService interface {
		Read(ctx context.Context, in *ReadAccountRequest, out *ReadAccountResponse) error
	}
	type AccountService struct {
		accountService
	}
	h := &accountServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&AccountService{h}, opts...))
}

type accountServiceHandler struct {
	AccountServiceHandler
}

func (h *accountServiceHandler) Read(ctx context.Context, in *ReadAccountRequest, out *ReadAccountResponse) error {
	return h.AccountServiceHandler.Read(ctx, in, out)
}

// Api Endpoints for ProjectService service

func NewProjectServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for ProjectService service

type ProjectService interface {
	Create(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*CreateProjectResponse, error)
	Update(ctx context.Context, in *UpdateProjectRequest, opts ...client.CallOption) (*UpdateProjectResponse, error)
	List(ctx context.Context, in *ListProjectsRequest, opts ...client.CallOption) (*ListProjectsResponse, error)
}

type projectService struct {
	c    client.Client
	name string
}

func NewProjectService(name string, c client.Client) ProjectService {
	return &projectService{
		c:    c,
		name: name,
	}
}

func (c *projectService) Create(ctx context.Context, in *CreateProjectRequest, opts ...client.CallOption) (*CreateProjectResponse, error) {
	req := c.c.NewRequest(c.name, "ProjectService.Create", in)
	out := new(CreateProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectService) Update(ctx context.Context, in *UpdateProjectRequest, opts ...client.CallOption) (*UpdateProjectResponse, error) {
	req := c.c.NewRequest(c.name, "ProjectService.Update", in)
	out := new(UpdateProjectResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *projectService) List(ctx context.Context, in *ListProjectsRequest, opts ...client.CallOption) (*ListProjectsResponse, error) {
	req := c.c.NewRequest(c.name, "ProjectService.List", in)
	out := new(ListProjectsResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProjectService service

type ProjectServiceHandler interface {
	Create(context.Context, *CreateProjectRequest, *CreateProjectResponse) error
	Update(context.Context, *UpdateProjectRequest, *UpdateProjectResponse) error
	List(context.Context, *ListProjectsRequest, *ListProjectsResponse) error
}

func RegisterProjectServiceHandler(s server.Server, hdlr ProjectServiceHandler, opts ...server.HandlerOption) error {
	type projectService interface {
		Create(ctx context.Context, in *CreateProjectRequest, out *CreateProjectResponse) error
		Update(ctx context.Context, in *UpdateProjectRequest, out *UpdateProjectResponse) error
		List(ctx context.Context, in *ListProjectsRequest, out *ListProjectsResponse) error
	}
	type ProjectService struct {
		projectService
	}
	h := &projectServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProjectService{h}, opts...))
}

type projectServiceHandler struct {
	ProjectServiceHandler
}

func (h *projectServiceHandler) Create(ctx context.Context, in *CreateProjectRequest, out *CreateProjectResponse) error {
	return h.ProjectServiceHandler.Create(ctx, in, out)
}

func (h *projectServiceHandler) Update(ctx context.Context, in *UpdateProjectRequest, out *UpdateProjectResponse) error {
	return h.ProjectServiceHandler.Update(ctx, in, out)
}

func (h *projectServiceHandler) List(ctx context.Context, in *ListProjectsRequest, out *ListProjectsResponse) error {
	return h.ProjectServiceHandler.List(ctx, in, out)
}
