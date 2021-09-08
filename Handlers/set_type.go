package Handlers

import (
	"KeTangPai/services/DC/KetangpaiDB"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Set_type(k KetangpaiDB.KetangpaiDBClient) gin.HandlerFunc {
	return func(c *gin.Context){
		typ,ok:=c.GetQuery("type")
		if !ok {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"missing necessary parameter",
			})
			return
		}
		typecode,err:=strconv.Atoi(typ)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		if typecode==KetangpaiDB.Administrator{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"illegal operation",
			})
			return
		}

		uid,err:=getInt("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit)
		re,err:=k.SetType(ctx,&KetangpaiDB.Member{Type: uint32(typecode),Uid: uint32(uid)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"return":re,
		})
	}
}
