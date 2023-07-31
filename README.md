# CloudWeGo API网关大作业说明文档

## 1、接口

### 1.1、http网关接口

http网关一共实现了两个接口，一个用于提供网关服务的接口，一个用于管理IDL资源的接口。

#### 1.1.1、网关服务接口

r.POST("/:svrName/:methodName", handler.RpcMethod)

func RpcMethod(ctx context.Context, c *app.RequestContext)

该接口用于提供网关服务，svrName是目标服务器注册的名字，methodName是目标方法的名字。

#### 1.1.2、IDL管理接口

r.POST("/idl-info/add/:idlName", handler.AddIdlInfo)

func AddIdlInfo(ctx context.Context, c *app.RequestContext)

该接口用于管理IDL资源，idlName是需要添加或改动IDL文件的服务器注册的名字，当服务器IDL文件发生变动时可以通过该接口修改IDL管理服务器中的内容，同时http网关也会在内部添加或更新相关客户端。

### 1.2、IDL管理服务器接口

IDL管理服务器一共实现了两个接口，一个用于添加或修改IDL文件内容，一个用于获取目标服务器的IDL文件内容

#### 1.2.1、添加或修改接口

func (s *IdlServiceImpl) Register(ctx context.Context, idlReq *demo.IdlReq) (resp *demo.AddIdlResp, err error)

该接口用于添加或修改IDL文件信息

#### 1.2.2、查询接口

func (s *IdlServiceImpl) Query(ctx context.Context, name string) (resp *demo.IdlInfo, err error)

该接口用于查询IDL文件信息，不对外开放，只能由http网关调用

## 2、部署步骤

1、开启etcd服务

2、开启httpSvr服务

3、开启idlSvr服务

4、开启各个rpcSvr服务