package Handlers

import (
	"KeTangPai/services/DC/JWT"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Logout(j JWT.JWTClient)gin.HandlerFunc{
	return func(c *gin.Context){
		id, err:=getUint("id",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		_,err=j.DelJwt(context.Background(),&JWT.Juser{Id: uint32(id)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		c.JSON(http.StatusInternalServerError,gin.H{
			"return":"success",
		})
	}
}
