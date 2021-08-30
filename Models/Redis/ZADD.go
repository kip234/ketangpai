package Redis

import "fmt"

func (r RedisPool)ZADD(args ...interface{})error{
	rdb :=r.wpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("ZADD",args...)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)ZADD(args ...interface{})error: %s",err.Error())
	}
	return err
}