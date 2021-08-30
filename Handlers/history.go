package Handlers

import (
	"KeTangPai/services/Forum"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func History(f Forum.ForumClient )gin.HandlerFunc{
	return func(c *gin.Context){
		classid,err:=getInt("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=f.GetHistory(ctx,&Forum.Classid{Id: classid})
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
