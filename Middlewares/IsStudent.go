package Middlewares

import (
	"KeTangPai/services/DC/KetangpaiDB"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//该用户是否是学生

func IsStudent(uservice KetangpaiDB.KetangpaiDBClient) gin.HandlerFunc {
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
		typ,err:=uservice.GetUserType(ctx,&KetangpaiDB.Uid{Uid:uint32(uid)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		if typ.Typecode!=KetangpaiDB.Student{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"only students can access it",
			})
			c.Abort()
			return
		}
	}
}
