package Handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Id uint32`json:"id"`
	Role string`json:"role"`
	Name string`json:"name"`
	Uid uint32`json:"uid"`
	Email string`json:"email"`
}


func Setting() gin.HandlerFunc {
	return func(c *gin.Context){
		tmp:=User{}
		uid,err := getUint("uid",c)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		tmp.Uid=uid

		tmp.Id,err=getUint("id",c)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		tmp.Role,err=getStr("role",c)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		tmp.Email,err=getStr("email",c)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		tmp.Name,err=getStr("name",c)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK,gin.H{
			"return":tmp,
		})
		return
	}
}
