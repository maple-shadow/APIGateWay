package main

import (
	"context"
	"fmt"
	demo "github.com/maple-shadow/idlsvr/kitex_gen/demo"
)

// IdlServiceImpl implements the last service interface defined in the IDL.
type IdlServiceImpl struct{}

var idlInfoMap = make(map[string]*demo.IdlInfo)

// Register implements the IdlServiceImpl interface.
func (s *IdlServiceImpl) Register(ctx context.Context, idlReq *demo.IdlReq) (resp *demo.AddIdlResp, err error) {
	// TODO: Your code here...

	fmt.Println(idlReq.IdlName)
	fmt.Println(idlReq.IdlInfo.Content)
	fmt.Println(idlReq.IdlInfo.Includes)

	idlInfoMap[idlReq.IdlName] = idlReq.IdlInfo

	resp = new(demo.AddIdlResp)

	fmt.Println(idlInfoMap)
	resp.Message = "add idl success"
	return resp, nil
}

// Query implements the IdlServiceImpl interface.
func (s *IdlServiceImpl) Query(ctx context.Context, name string) (resp *demo.IdlInfo, err error) {
	// TODO: Your code here...

	fmt.Println("query: " + name)
	resp = new(demo.IdlInfo)

	resp = idlInfoMap[name]

	fmt.Println(resp)
	return resp, nil
}
