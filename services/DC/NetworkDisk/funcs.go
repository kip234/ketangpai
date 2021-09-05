package NetworkDisk

import (
	"google.golang.org/grpc"
	"net"
)

//启动服务
func Run(){
	jwt:= newNetworkDiskService()
	lis,err:=net.Listen("tcp", Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterNetworkDiskServer(s,jwt)
	err=s.Serve(lis)
	panic(err)
}

//创建调用连接
func NewTestBankConn() NetworkDiskClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:= NewNetworkDiskClient(conn)
	return c
}