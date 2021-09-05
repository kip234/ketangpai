package Exercise

import (
	"google.golang.org/grpc"
	"net"
)

//启动服务
func Run(){
	jwt:= newExerciseService()
	lis,err:=net.Listen("tcp", Addr)
	if err!=nil {
		panic(err)
	}
	s:=grpc.NewServer()
	RegisterExerciseServer(s,jwt)
	err=s.Serve(lis)
	panic(err)
}
//生成调用连接
func NewExerciseConn() ExerciseClient {
	conn,err:=grpc.Dial(Addr,grpc.WithInsecure())
	if err!=nil {
		panic(err)
	}
	c:= NewExerciseClient(conn)
	return c
}