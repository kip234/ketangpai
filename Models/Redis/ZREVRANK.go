package Redis

import (
	"errors"
	"fmt"
)

func (r RedisPool)ZREVRANK(key,member string)(uint64,error){
	rdb :=r.rpool.Get()
	defer rdb.Close()
	v,err:=rdb.Do("ZREVRANK",key,member)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)ZREVRANK(key,member string)(re []string,err error): %s",err.Error())
	}
	if _,ok:=v.(uint64);!ok{
		return 0,errors.New("type conflict")
	}
	return v.(uint64),err
}
