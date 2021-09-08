package Middlewares

import (
	"KeTangPai/services/DC/KetangpaiDB"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//判断该用户是否是管理员

func IsAdmin(u KetangpaiDB.KetangpaiDBClient) gin.HandlerFunc{
	return func(c *gin.Context){
		uid,err:=getInt("uid",c)//从上下文中获取UID
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"missing necessary parameter",
			})
			c.Abort()
			return
		}
		ctx,_:=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		re,err:=u.GetUserType(ctx,&KetangpaiDB.Uid{Uid: uint32(uid)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		if re.Typecode!=KetangpaiDB.Administrator{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"only administrators can access it",
			})
			c.Abort()
			return
		}
	}
}
