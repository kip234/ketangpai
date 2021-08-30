package Redis

import "fmt"

func (r RedisPool)ZREVRANGE(key string,start,stop int32)(re []string,err error){
	rdb :=r.rpool.Get()
	defer rdb.Close()
	v,err:=rdb.Do("ZREVRANGE",key,start,stop,"WITHSCORES")
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)ZREVRANGE(args ...interface{})(re []string,err error): %s",err.Error())
	}
	for _,i:=range v.([]interface{}){
		re=append(re,string(i.([]uint8)))
	}
	return re,err
}