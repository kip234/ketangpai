package Forum

import (
	"google.golang.org/grpc"
	"net"
)

func Run(){
	ucs:=newForumService()
	lis,err:=net.Listen("tcp",Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterForumServer(s,ucs)
	err=s.Serve(lis)
	panic(err)
}

func NewForumConn() ForumClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:=NewForumClient(conn)
	return c
}