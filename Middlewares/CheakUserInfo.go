package Middlewares

import (
	"KeTangPai/services/DC/UserCenter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

//检查用户的信息是否正确-用于登录

func CheakUserInfo(uservice UserCenter.UserCenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		user:=UserCenter.Userdb{}
		err:=c.ShouldBind(&user)//绑定用户数据
		if err!=nil{//绑定失败
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		ctx,_:=context.WithTimeout(context.Background(), serviceTimeLimit)
		tmp,err:=uservice.GetUserInfoByEmail(ctx,&UserCenter.S{S:user.Email})//获取后台记录
		if err!=nil{//查找失败
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}
		if tmp.Pwd!=user.Pwd{//密码对不上
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"password wrong !",
			})
			c.Abort()
			return
		}

		c.Set("id",tmp.Id)//验证通过,存入UID
	}
}