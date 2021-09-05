package Middlewares

import (
	"KeTangPai/services/DC/JWT"
	"KeTangPai/services/DC/UserCenter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//检查JWT是否被修改，如果一切顺利会把用户名，用户ID以及用户注册用的邮箱放入上下文

//读取
func CheakJWT(jwt JWT.JWTClient,uc UserCenter.UserCenterClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("token")//找token
		if token == ""{//没有token
			c.JSON(http.StatusBadRequest,gin.H{
				"content":"token == nil",
			})
			c.Abort()
			return
		}
		ctx,_:=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		u,err:=jwt.CheckToken(ctx,&JWT.Token{Content: token})//检查传入的token
		if err!=nil {//对方出问题
			c.JSON(http.StatusInternalServerError,gin.H{
				"content":err.Error(),
			})
			c.Abort()
			return
		}
		ctx,_=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		user,err:=uc.GetUserInfo(ctx,&UserCenter.Id{I: u.Uid})//获取用户的相关信息
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"content":err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("uname",user.Name)//存入用户名
		c.Set("uid",u.Uid)//存入UID
		c.Set("email",user.Email)//存入邮箱
	}
}