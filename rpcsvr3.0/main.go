package main

import (
	"github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/maple-shadow/rpcsvr3.0/kitex_gen/kitex/test/server/exampleservice"
	"log"
	"net"

	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8891")
	//svr := demo.NewServer(new(StudentServiceImpl), server.WithServiceAddr(addr))

	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	ri := &registry.Info{
		ServiceName: "CServiceName",
		Tags: map[string]string{
			"Cluster": "xxx",
		},
	}

	svr := exampleservice.NewServer(
		new(ExampleServiceImpl),
		server.WithRegistry(r),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
			ServiceName: "CServiceName",
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
