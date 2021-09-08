package Handlers

import (
	"KeTangPai/services/DC/KetangpaiDB"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	//"strconv"
)

func Fire(u KetangpaiDB.KetangpaiDBClient) gin.HandlerFunc {
	return func(c *gin.Context){
		classid,err:=getInt("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		//自己检查一遍是不是自己的学生
		var Students []uint32
		c.ShouldBindJSON(&Students)
		for _,i:=range Students{
			ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
			tmp,err:=u.GetUserClass(ctx,&KetangpaiDB.Uid{Uid:i})
			if err!=nil {
				c.JSON(http.StatusInternalServerError,gin.H{
					"error":err.Error(),
				})
				return
			}
			if tmp.Classid==uint32(classid){
				ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
				_,err:=u.FireStudent(ctx,&KetangpaiDB.Uid{Uid:i})
				if err!=nil {
					c.JSON(http.StatusInternalServerError,gin.H{
						"error":err.Error(),
					})
					return
				}
			}
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		class,err:=u.GetClassInfo(ctx,&KetangpaiDB.Classid{Classid:uint32(classid)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"return":class,
		})
	}
}
