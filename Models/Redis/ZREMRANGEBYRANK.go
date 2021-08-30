package Redis

import (
	"fmt"
)

func (r RedisPool)ZREMRANGEBYRANK(key string,start,stop int32)error{
	rdb :=r.rpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("ZREMRANGEBYRANK",key,start,stop)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)ZREMRANGEBYRANK(key,member string)(int32,error): %s",err.Error())
	}
	return err
}

