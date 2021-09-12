package Handlers

import (
	"KeTangPai/services/RBAC"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add_routes(r RBAC.RBACClient) gin.HandlerFunc {
	return func(c *gin.Context){
		var ra []string
		err:=c.ShouldBind(&ra)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.TODO(),serviceTimeLimit)
		_,err=r.RefreshPaths(ctx,&RBAC.Paths{Paths: ra})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"return":"nil",
		})
	}
}
