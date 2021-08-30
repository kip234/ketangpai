package Middlewares

import (
	"KeTangPai/services/DC/UserCenter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func HaveClass(u  UserCenter.UserCenterClient) gin.HandlerFunc {
	return func(c *gin.Context){
		uid,err:=getInt("uid",c)//获取UID
		if err!=nil {//获取UID出错
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		ctx,_:=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		re,err:=u.GetUserClass(ctx,&UserCenter.Id{I:uid})//获取用户所属班级
		if err!=nil {//
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		if re.I<1{//没有所属班级
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"there is no class for you",
			})
			c.Abort()
			return
		}
		ctx,_=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		name,err:=u.GetClassName(ctx,&UserCenter.Id{I: re.I})
		if err!=nil {//
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("classname",name.S)
		c.Set("classid",re.I)//存入classID
	}
}
