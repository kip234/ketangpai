package Email

import (
	"google.golang.org/grpc"
	"net"
)

func Run(){
	ucs:=newEmailService()
	lis,err:=net.Listen("tcp",Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterEmailServer(s,ucs)
	err=s.Serve(lis)
	panic(err)
}

func NewEmailConn() EmailClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:=NewEmailClient(conn)
	return c
}