package Handlers

import (
	"KeTangPai/services/DC/UserCenter"
	"KeTangPai/services/Email"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Retrieve(u UserCenter.UserCenterClient,e Email.EmailClient)gin.HandlerFunc{
	return func(c *gin.Context){
		email,ok:=c.GetQuery("email")
		if !ok {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"missing necessary parameter",
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		right,err:=u.UserIs_Exist(ctx,&UserCenter.S{S: email})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		if !right.Right{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":"the mailbox is not registered",
			})
			return
		}
		ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=u.GetUserInfoByEmail(ctx,&UserCenter.S{S:email})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		_,err=e.Send(ctx,&Email.Mail{Subject: subject,To: email,Content: fmt.Sprintf(body,re)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"return":"email has been sent",
		})
	}
}
