package Handlers

import (
	"KeTangPai/services/RBAC"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Add_user_role(r RBAC.RBACClient) gin.HandlerFunc {
	return func(c *gin.Context){
		ur:=RBAC.UserRoledb{}
		err:=c.ShouldBind(&ur)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.TODO(),serviceTimeLimit)
		_,err=r.RefreshUserRole(ctx,&RBAC.UsersRoles{Uids: []uint32{ur.Uid},Roles: []string{ur.Role}})
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
