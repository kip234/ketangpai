package Handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

//用于从上下文中取出UID
func getUint(key string,c *gin.Context) (uid uint32,err error) {
	v,ok:=c.Get(key)
	if !ok {
		uid=0
		err = fmt.Errorf("Missing %s field",key)
		return
	}
	uid,ok = v.(uint32)
	if !ok {
		uid=0
		err = fmt.Errorf("Assertion failure")
		return
	}
	return uid,nil
}

func getStr(key string,c *gin.Context) (str string,err error) {
	v,ok:=c.Get(key)
	if !ok {
		err = fmt.Errorf("Missing %s field",key)
		return
	}
	str,ok = v.(string)
	if !ok {
		err = fmt.Errorf("Assertion failure")
		return
	}
	return str,nil
}