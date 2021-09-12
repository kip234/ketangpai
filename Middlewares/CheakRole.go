package Middlewares

import (
	"KeTangPai/services/RBAC"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"net/http"
)

func CheakRole(r RBAC.RBACClient,path string) gin.HandlerFunc {
	return func(c *gin.Context){
		role,err:=getStr("role",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		ctx,_:=context.WithTimeout(context.TODO(),serviceTimeLimit)
		re,err:=r.CheakRolePath(ctx,&RBAC.RolesPaths{Roles: []string{role},Paths: []string{path}})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			c.Abort()
			return
		}

		if !re.Bools[0]{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"you don't have access.",
			})
			c.Abort()
			return
		}
	}
}
