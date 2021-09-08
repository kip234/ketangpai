package Handlers

import (
	"KeTangPai/services/DC/KetangpaiDB"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)
//新建一个班级-一个老师只能有一个班级，该操作会自动解散原有班级
func Create_class(u KetangpaiDB.KetangpaiDBClient) gin.HandlerFunc {
	return func(c *gin.Context){
		uid,err:=getInt("uid",c)//获取当前用户ID
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		class:=KetangpaiDB.Classdb{}
		err=c.ShouldBind(&class)//获取传回来的班级信息
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		class.Teacher=uint32(uid)//指明老师
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=u.CreateClass(ctx,&KetangpaiDB.Class{
			Classid:class.Classid,
			Teacher:class.Teacher,
			Name:class.Name,
			Students:class.Students,
		})//创建记录
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
