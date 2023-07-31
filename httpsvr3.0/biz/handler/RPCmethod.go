package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	kclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/loadbalance"
	etcd "github.com/kitex-contrib/registry-etcd"
	"github.com/maple-shadow/httpsvr2.0/kitex_gen/demo/idlservice"
)

var cliMap = make(map[string]genericclient.Client)

func RpcMethod(ctx context.Context, c *app.RequestContext) {
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		panic("get http req failed")
	}

	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic(err)
		//panic("get custom req failed")
	}

	jsonStr, err := json.Marshal(customReq.Body)
	jsonReq := string(jsonStr)
	fmt.Println(jsonReq)
	//fmt.Println(req)
	//c.JSON(consts.StatusOK, req)
	//fmt.Println(req)

	targetName := c.Param("svrName")

	cli := cliMap[targetName]

	if cli == nil {
		cli = getIdlInfo(ctx, targetName)
	}

	if cli == nil {
		c.JSON(consts.StatusOK, "can not find "+targetName)
		return
	}

	fmt.Println(targetName)
	fmt.Println(c.Param("methodName"))

	resp, err := cli.GenericCall(ctx, c.Param("methodName"), jsonReq)

	//fmt.Println(resp)
	c.JSON(consts.StatusOK, resp)
	//fmt.Println(resp)
}

func AddIdlInfo(ctx context.Context, c *app.RequestContext) {
	httpReq, err := adaptor.GetCompatRequest(c.GetRequest())
	if err != nil {
		panic("get http req failed")
	}

	customReq, err := generic.FromHTTPRequest(httpReq)
	if err != nil {
		panic(err)
		//panic("get custom req failed")
	}

	jsonStr, err := json.Marshal(customReq.Body)
	jsonReq := string(jsonStr)
	fmt.Println(jsonReq)
	//fmt.Println(req)
	//c.JSON(consts.StatusOK, req)
	//fmt.Println(req)

	cli := initIdlGenericClient("idlSvr")

	resp, err := cli.GenericCall(ctx, "Register", jsonReq)

	getIdlInfo(ctx, c.Param("idlName"))

	//fmt.Println(resp)
	c.JSON(consts.StatusOK, resp)
	//fmt.Println(resp)
}

func getIdlInfo(ctx context.Context, name string) genericclient.Client {
	idlCli, err := idlservice.NewClient("idlSvr", kclient.WithHostPorts("127.0.0.1:9999"))
	if err != nil {
		panic("err new client:" + err.Error())
	}

	fmt.Println("get: " + name)

	idlResp, err := idlCli.Query(ctx, name)

	//fmt.Println(idlResp)
	//
	//fmt.Println(idlResp.Content)
	if idlResp == nil {
		return nil
	}

	if err != nil {
		panic("err query rpc server:" + err.Error())
	}

	newCli := initGenericClient(name, idlResp.Content, idlResp.Includes)

	cliMap[name] = newCli

	return cliMap[name]
}

func initGenericClient(svrName string, content string, includes map[string]string) genericclient.Client {
	// 本地文件 idl 解析
	// YOUR_IDL_PATH thrift 文件路径: 举例 ./idl/example.thrift
	// includeDirs: 指定 include 路径，默认用当前文件的相对路径寻找 include
	//p, err := generic.NewThriftFileProvider("../example_service.thrift")

	//p, err := generic.NewThriftFileProvider("../idlsvr2.0/idl.thrift")
	//if err != nil {
	//	panic(err)
	//}
	p, err := generic.NewThriftContentProvider(content, includes)
	if err != nil {
		panic(err)
	}

	// 构造 http 类型的泛化调用
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic("err in resolver")
		//log.Fatal(err)
	}

	cli, err := genericclient.NewClient(
		svrName,
		g,
		kclient.WithResolver(r),
		kclient.WithTag("Cluster", "xxx"),
		kclient.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()),
	)
	if err != nil {
		//return nil
		panic(err)
	}
	return cli
}

func initIdlGenericClient(svrName string) genericclient.Client {
	// 本地文件 idl 解析
	// YOUR_IDL_PATH thrift 文件路径: 举例 ./idl/example.thrift
	// includeDirs: 指定 include 路径，默认用当前文件的相对路径寻找 include
	//p, err := generic.NewThriftFileProvider("../example_service.thrift")

	//p, err := generic.NewThriftFileProvider("../idlsvr2.0/idl.thrift")
	//if err != nil {
	//	panic(err)
	//}

	path := "a/b/main.thrift"
	content := `
namespace go demo



struct IdlInfo {
    1: string content(go.tag = 'json:"content"'),
    2: map<string, string> includes(go.tag = 'json:"includes"'),
}

struct IdlReq {
    1: string idlName(api.body='name'),
    2: IdlInfo idlInfo(api.body='info'),
}

struct AddIdlResp {
    1: bool success(api.body='success'),
    2: string message(api.body='message'),
}

service IdlService {
    AddIdlResp Register(1: IdlReq idlReq)(api.post = '/add-ldl-info')
    IdlInfo Query(1: string name)(api.get = '/query')
}
	`
	includes := map[string]string{
		path: content,
	}

	p, err := generic.NewThriftContentProvider(content, includes)
	if err != nil {
		panic(err)
	}

	// 构造 http 类型的泛化调用
	g, err := generic.JSONThriftGeneric(p)
	if err != nil {
		panic(err)
	}

	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic("err in resolver")
		//log.Fatal(err)
	}

	cli, err := genericclient.NewClient(
		svrName,
		g,
		kclient.WithResolver(r),
		kclient.WithTag("Cluster", "xxx"),
		kclient.WithLoadBalancer(loadbalance.NewWeightedRandomBalancer()),
	)
	if err != nil {
		panic(err)
	}
	return cli
}
