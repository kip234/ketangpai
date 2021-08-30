package Middlewares

import (
	"KeTangPai/services/DC/JWT"
	"KeTangPai/services/DC/UserCenter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

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
		u,err:=jwt.CheckToken(ctx,&JWT.Token{Content: token})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"content":err.Error(),
			})
			c.Abort()
			return
		}
		ctx,_=context.WithTimeout(context.Background(), serviceTimeLimit*time.Second)
		user,err:=uc.GetUserInfo(ctx,&UserCenter.Id{I: u.Uid})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"content":err.Error(),
			})
			c.Abort()
			return
		}
		c.Set("uname",user.Name)
		c.Set("uid",u.Uid)//存入UID
		c.Set("email",user.Email)
	}
}