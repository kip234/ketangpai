package NetworkDisk

import (
	"google.golang.org/grpc"
	"net"
)

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

func NewTestBankConn() NetworkDiskClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:= NewNetworkDiskClient(conn)
	return c
}