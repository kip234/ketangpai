package Handlers

import (
	"KeTangPai/services/RBAC"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Cache(r RBAC.RBACClient)gin.HandlerFunc{
	return func(c *gin.Context){
		ctx,_:=context.WithTimeout(context.TODO(),serviceTimeLimit)
		_,err:=r.Cache(ctx,&RBAC.Empty{})
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"return":"success",
		})
	}
}
