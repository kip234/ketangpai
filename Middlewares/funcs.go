package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//一些几乎经常用的函数

//用于从上下文中取出 int32 的数据(最早是取UID,现在还要取classID等)
func getInt(key string,c *gin.Context) (uid int32,err error) {
	v,ok:=c.Get(key)
	if !ok {
		uid=-1
		err = fmt.Errorf("Missing UID field")
		return
	}
	uid,ok = v.(int32)
	if !ok {
		uid=-1
		err = fmt.Errorf("Assertion failure")
		return
	}
	return uid,nil
}
