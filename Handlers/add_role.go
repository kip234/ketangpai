package Handlers

import (
	"KeTangPai/services/RBAC"
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Add_role(r RBAC.RBACClient) gin.HandlerFunc {
	return func(c *gin.Context){
		var role []string
		err:=c.ShouldBind(&role)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		log.Println(role)
		ctx,_:=context.WithTimeout(context.TODO(),serviceTimeLimit)
		re,err:=r.AddRole(ctx,&RBAC.Roles{Roles: role})
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
