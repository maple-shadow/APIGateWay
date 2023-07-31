// Code generated by Kitex v0.6.1. DO NOT EDIT.

package idlservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	demo "github.com/maple-shadow/idlsvr/kitex_gen/demo"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, idlReq *demo.IdlReq, callOptions ...callopt.Option) (r *demo.AddIdlResp, err error)
	Query(ctx context.Context, name string, callOptions ...callopt.Option) (r *demo.IdlInfo, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kIdlServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kIdlServiceClient struct {
	*kClient
}

func (p *kIdlServiceClient) Register(ctx context.Context, idlReq *demo.IdlReq, callOptions ...callopt.Option) (r *demo.AddIdlResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, idlReq)
}

func (p *kIdlServiceClient) Query(ctx context.Context, name string, callOptions ...callopt.Option) (r *demo.IdlInfo, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Query(ctx, name)
}
