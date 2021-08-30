package Filter

import (
	"google.golang.org/grpc"
	"net"
)


func Run(){
	err,f:=newFilter()
	if err!=nil{
		panic(err)
	}
	lis,err:=net.Listen("tcp",Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterFilterServer(s,f)
	err=s.Serve(lis)
	if err!=nil {
		panic(err)
	}
}

func NewFilterConn() FilterClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:=NewFilterClient(conn)
	return c
}
