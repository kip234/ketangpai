package Redis

import "fmt"

func (r RedisPool)ZINCRBY(key string,increment int32,member string)error{
	rdb :=r.wpool.Get()
	defer rdb.Close()
	_,err:=rdb.Do("ZINCRBY",key,increment,member)
	if err != nil{
		err=fmt.Errorf("(r RedisPool)ZINCRBY(args ...interface{})error: %s",err.Error())
	}
	return err
}