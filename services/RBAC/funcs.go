package RBAC

import (
	"google.golang.org/grpc"
	"net"
)

func Run(){
	ucs:=newRBACService()
	ucs.cache()//开始缓存
	lis,err:=net.Listen("tcp",Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterRBACServer(s,ucs)
	err=s.Serve(lis)
	panic(err)
}

func NewRBACConn() RBACClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:=NewRBACClient(conn)
	return c
}

