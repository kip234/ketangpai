package Redis

import "fmt"

func (r RedisPool)SISMEMBER(args ...interface{}) (bool,error) {
	rdb :=r.wpool.Get()
	defer rdb.Close()
	v,err:=rdb.Do("SISMEMBER",args...)
	if err != nil{
		err=fmt.Errorf("func (r RedisPool)DEL(key string) error: %s",err.Error())
	}
	re,ok:=v.(int64)
	if !ok {
		return false,fmt.Errorf("func (r RedisPool)SISMEMBER(args ...interface{}) (int,error) : Assertion error")
	}
	if re==1{
		return true,nil
	}else{
		return false,nil
	}
}
