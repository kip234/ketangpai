package UserCenter

import (
	"google.golang.org/grpc"
	"net"
)

func Run(){
	ucs:=newUserCenterService()
	lis,err:=net.Listen("tcp",Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterUserCenterServer(s,ucs)
	err=s.Serve(lis)
	panic(err)
}

func NewUCSConn() UserCenterClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:=NewUserCenterClient(conn)
	return c
}