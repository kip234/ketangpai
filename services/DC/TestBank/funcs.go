package TestBank

import (
	"google.golang.org/grpc"
	"net"
)

func Run(){
	jwt:= newTestBankService()
	lis,err:=net.Listen("tcp",Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterTestBankServer(s,jwt)
	err=s.Serve(lis)
	panic(err)
}

func NewTestBankConn() TestBankClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:= NewTestBankClient(conn)
	return c
}