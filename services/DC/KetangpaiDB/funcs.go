package KetangpaiDB

import (
	"google.golang.org/grpc"
	"net"
)

//启动服务
func Run(){
	jwt:= newKetangpaiDBService()
	lis,err:=net.Listen("tcp", Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterKetangpaiDBServer(s,jwt)
	err=s.Serve(lis)
	panic(err)
}

//创建调用连接
func NewKetangpaiDBConn() KetangpaiDBClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:= NewKetangpaiDBClient(conn)
	return c
}
