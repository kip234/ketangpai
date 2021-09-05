package Middlewares

import (
	"KeTangPai/services/DC/UserCenter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//检查用户的信息是否正确-用于登录

func CheakUserInfo(uservice UserCenter.UserCenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		user:=UserCenter.Userdb{}
		err:=c.ShouldBind(&user)//绑定用户数据
		if err!=nil{//绑定失败
			c.JSON(http.StatusOK,gin.H{
				"content":err.Error(),
			})
			c.Abort()
			return
		}
		ctx,_:=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		tmp,err:=uservice.GetUserInfo(ctx,&UserCenter.Id{I:user.Uid})//获取后台记录
		if err!=nil{//查找失败
			c.JSON(http.StatusOK,gin.H{
				//"typ":Data.ErrTyp,
				"content":err.Error(),
			})
			c.Abort()
			return
		}
		if tmp.Pwd!=user.Pwd{//密码对不上
			c.JSON(http.StatusOK,gin.H{
				//"typ":Data.ErrTyp,
				"content":"password wrong !",
			})
			c.Abort()
			return
		}
		c.Set("uid",tmp.Uid)//验证通过,存入UID
	}
}