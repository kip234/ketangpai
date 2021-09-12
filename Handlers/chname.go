package Handlers

import (
	"KeTangPai/services/DC/KetangpaiDB"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Chname(k KetangpaiDB.KetangpaiDBClient)gin.HandlerFunc{
	return func(c *gin.Context){
		name,ok:=c.GetQuery("name")
		if !ok {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"missing necessary parameters",
			})
			return
		}
		if name==""{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"the length of name is invalid",
			})
			return
		}

		uid,err:=getUint("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		ctx,_:=context.WithTimeout(context.TODO(),serviceTimeLimit)
		re,err:=k.SetUserName(ctx,&KetangpaiDB.User{Name: name,Uid: uid})
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"return":re,
		})
	}
}
