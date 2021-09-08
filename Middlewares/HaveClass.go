package Middlewares

import (
	"KeTangPai/services/DC/KetangpaiDB"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//判断该用户是否有所属班级


func HaveClass(u KetangpaiDB.KetangpaiDBClient) gin.HandlerFunc {
	return func(c *gin.Context){
		uid,err:=getInt("uid",c)//获取UID
		if err!=nil {//获取UID出错
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		ctx,_:=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		re,err:=u.GetUserClass(ctx,&KetangpaiDB.Uid{Uid:uint32(uid)})//获取用户所属班级
		if err!=nil {//
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		if re.Classid<1{//没有所属班级
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"there is no class for you",
			})
			c.Abort()
			return
		}
		ctx,_=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		name,err:=u.GetClassName(ctx,re)
		if err!=nil {//
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		c.Set("classname",name.Name)
		c.Set("classid",re.Classid)//存入classID
	}
}
