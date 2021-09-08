package Handlers

import (
	"KeTangPai/services/DC/JWT"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Login(jwt JWT.JWTClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid,err:=getInt("uid",c)//获取UID
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		t,err:=jwt.RefreshToken(ctx,&JWT.Juser{Uid: uint32(uid)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{//把令牌返回给客户
			"return":t.Content,
		})
	}
}