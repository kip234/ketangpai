package RankingList

import (
	"google.golang.org/grpc"
	"net"
)

func Run(){
	ucs:=newRankingListService()
	lis,err:=net.Listen("tcp",Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterRankingListServer(s,ucs)
	err=s.Serve(lis)
	panic(err)
}

func NewRankingListConn() RankingListClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:=NewRankingListClient(conn)
	return c
}
