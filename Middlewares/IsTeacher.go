package Middlewares

import (
	"KeTangPai/services/DC/UserCenter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//判断用户是否是老师

func IsTeacher(uservice UserCenter.UserCenterClient) gin.HandlerFunc {
	return func (c *gin.Context){
		uid,err:=getInt("uid",c)
		if err!=nil {
			c.JSON(http.StatusOK,gin.H{
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
		if typ.I!=UserCenter.Teacher{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"only teacher can access it",
			})
			c.Abort()
			return
		}
	}
}
