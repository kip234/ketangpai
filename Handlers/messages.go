package Handlers

import (
	"KeTangPai/services/Forum"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Messages(f Forum.ForumClient)gin.HandlerFunc{
	return func(c *gin.Context){
		uid,err:=getInt("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=f.GetMessage(ctx,&Forum.Uid{Id: uid})
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"return":re,
		})
	}
}
