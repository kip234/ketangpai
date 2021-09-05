package JWT

import (
	"google.golang.org/grpc"
	"net"
)

//启动服务
func Run(){
	jwt:=newJwtService()
	lis,err:=net.Listen("tcp",Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterJWTServer(s,jwt)
	err=s.Serve(lis)
	panic(err)
}

//生成调用连接
func NewJWTConn() JWTClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:=NewJWTClient(conn)
	return c
}
