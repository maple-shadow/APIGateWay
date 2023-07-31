package main

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/maple-shadow/rpcsvr/kitex_gen/kitex/test/server/exampleservice"
	_ "github.com/maple-shadow/rpcsvr/kitex_gen/kitex/test/server/exampleservice"
	"log"
	_ "log"
	"net"

	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	//svr := server0.NewServer(new(ExampleServiceImpl))
	//
	//err := svr.Run()
	//
	//if err != nil {
	//	log.Println(err.Error())
	//}

	// 本地文件 idl 解析
	// YOUR_IDL_PATH thrift 文件路径: e.g. ./idl/example.thrift
	//p, err := generic.NewThriftFileProvider("../example_service.thrift")
	//if err != nil {
	//	panic(err)
	//}
	//// 构造 JSON 请求和返回类型的泛化调用
	//g, err := generic.JSONThriftGeneric(p)
	//if err != nil {
	//	panic(err)
	//}

	//svr := genericserver.NewServer(new(GenericServiceImpl), g, opts...)
	//if err != nil {
	//	panic(err)
	//}
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8890")
	//svr := demo.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	ri := &registry.Info{
		ServiceName: "BServiceName",
		Tags: map[string]string{
			"Cluster": "xxx",
		},
	}

	svr := exampleservice.NewServer(
		new(ExampleServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "BServiceName",
		}),
		server.WithRegistryInfo(ri),
		server.WithServiceAddr(addr),
	)

	err = svr.Run()
	if err != nil {
		panic(err)
	}
	// resp is a JSON string
}

//type GenericServiceImpl struct {
//}
//
//func (g *GenericServiceImpl) GenericCall(ctx context.Context, method string, request interface{}) (response interface{}, err error) {
//	// use jsoniter or other json parse sdk to assert request
//	m := request.(string)
//	fmt.Printf("Recv: %v\n", m)
//	if method == "ExampleMethod" {
//		req := &server2.ExampleReq{
//			Msg:  m,
//			Base: nil,
//		}
//		var s *ExampleServiceImpl
//		resp, err := s.ExampleMethod(ctx, req)
//		if err != nil {
//			return nil, err
//		}
//		return resp, nil
//	}
//	return "{\"Msg\": \"Hello,world\"}", nil
//}
