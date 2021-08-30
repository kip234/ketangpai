package Handlers

import (
	"KeTangPai/services/DC/UserCenter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"

	//"strconv"
)

func Fire(u UserCenter.UserCenterClient) gin.HandlerFunc {
	return func(c *gin.Context){
		classid,err:=getInt("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		//自己检查一遍是不是自己的学生
		var Students []int32
		c.ShouldBindJSON(&Students)
		for _,i:=range Students{
			ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
			tmp,err:=u.GetUserClass(ctx,&UserCenter.Id{I:i})
			if err!=nil {
				c.JSON(http.StatusInternalServerError,gin.H{
					"error":err.Error(),
				})
				return
			}
			if tmp.I==classid{
				ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
				_,err:=u.FireStudent(ctx,&UserCenter.Id{I:i})
				if err!=nil {
					c.JSON(http.StatusInternalServerError,gin.H{
						"error":err.Error(),
					})
					return
				}
			}
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		class,err:=u.GetClassInfo(ctx,&UserCenter.Id{I:classid})
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
