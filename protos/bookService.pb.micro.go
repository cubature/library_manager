// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: library_manager/protos/bookService.proto

package protos

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

// Api Endpoints for BookService service

func NewBookServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for BookService service

type BookService interface {
	ListAll(ctx context.Context, in *Empty, opts ...client.CallOption) (*BookList, error)
}

type bookService struct {
	c    client.Client
	name string
}

func NewBookService(name string, c client.Client) BookService {
	return &bookService{
		c:    c,
		name: name,
	}
}

func (c *bookService) ListAll(ctx context.Context, in *Empty, opts ...client.CallOption) (*BookList, error) {
	req := c.c.NewRequest(c.name, "BookService.ListAll", in)
	out := new(BookList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for BookService service

type BookServiceHandler interface {
	ListAll(context.Context, *Empty, *BookList) error
}

func RegisterBookServiceHandler(s server.Server, hdlr BookServiceHandler, opts ...server.HandlerOption) error {
	type bookService interface {
		ListAll(ctx context.Context, in *Empty, out *BookList) error
	}
	type BookService struct {
		bookService
	}
	h := &bookServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&BookService{h}, opts...))
}

type bookServiceHandler struct {
	BookServiceHandler
}

func (h *bookServiceHandler) ListAll(ctx context.Context, in *Empty, out *BookList) error {
	return h.BookServiceHandler.ListAll(ctx, in, out)
}