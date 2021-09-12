package Handlers

import (
	"KeTangPai/services/DC/KetangpaiDB"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add(k KetangpaiDB.KetangpaiDBClient) gin.HandlerFunc {
	return func(c *gin.Context){
		classid,err:=getUint("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		var students []uint32
		err=c.ShouldBindJSON(&students)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		for _,i:=range students{
			ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit)
			_,err=k.AddStudent(ctx,&KetangpaiDB.Member{Classid: uint32(classid),Uid: i})
			if err!=nil {
				c.JSON(http.StatusInternalServerError,gin.H{
					"error":err.Error(),
				})
				return
			}
		}
	}
}
