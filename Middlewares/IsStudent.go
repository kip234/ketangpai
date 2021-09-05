package Middlewares

import (
	"KeTangPai/services/DC/UserCenter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//该用户是否是学生

func IsStudent(uservice UserCenter.UserCenterClient) gin.HandlerFunc {
	return func(c *gin.Context){
		uid,err:=getInt("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		ctx,_:=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		typ,err:=uservice.GetUserType(ctx,&UserCenter.Id{I:uid})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		if typ.I!=UserCenter.Student{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"only students can access it",
			})
			c.Abort()
			return
		}
	}
}
