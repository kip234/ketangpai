package Handlers

import (
	"KeTangPai/services/DC/NetworkDisk"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func File_contents(n NetworkDisk.NetworkDiskClient)gin.HandlerFunc{
	return func(c *gin.Context){
		classid,err:= getUint("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=n.GetContents(ctx,&NetworkDisk.Classid{Id: uint32(classid)})
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
