// Code generated by Kitex v0.6.1. DO NOT EDIT.

package idlservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	demo "github.com/maple-shadow/httpsvr2.0/kitex_gen/demo"
)

func serviceInfo() *kitex.ServiceInfo {
	return idlServiceServiceInfo
}

var idlServiceServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "IdlService"
	handlerType := (*demo.IdlService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register": kitex.NewMethodInfo(registerHandler, newIdlServiceRegisterArgs, newIdlServiceRegisterResult, false),
		"Query":    kitex.NewMethodInfo(queryHandler, newIdlServiceQueryArgs, newIdlServiceQueryResult, false),
	}
	extra := map[string]interface{}{
		"PackageName": "demo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.6.1",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.IdlServiceRegisterArgs)
	realResult := result.(*demo.IdlServiceRegisterResult)
	success, err := handler.(demo.IdlService).Register(ctx, realArg.IdlReq)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newIdlServiceRegisterArgs() interface{} {
	return demo.NewIdlServiceRegisterArgs()
}

func newIdlServiceRegisterResult() interface{} {
	return demo.NewIdlServiceRegisterResult()
}

func queryHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.IdlServiceQueryArgs)
	realResult := result.(*demo.IdlServiceQueryResult)
	success, err := handler.(demo.IdlService).Query(ctx, realArg.Name)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newIdlServiceQueryArgs() interface{} {
	return demo.NewIdlServiceQueryArgs()
}

func newIdlServiceQueryResult() interface{} {
	return demo.NewIdlServiceQueryResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, idlReq *demo.IdlReq) (r *demo.AddIdlResp, err error) {
	var _args demo.IdlServiceRegisterArgs
	_args.IdlReq = idlReq
	var _result demo.IdlServiceRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Query(ctx context.Context, name string) (r *demo.IdlInfo, err error) {
	var _args demo.IdlServiceQueryArgs
	_args.Name = name
	var _result demo.IdlServiceQueryResult
	if err = p.c.Call(ctx, "Query", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
