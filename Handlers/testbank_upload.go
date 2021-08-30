package Handlers

import (
	"KeTangPai/services/DC/TestBank"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Testbank_upload(t TestBank.TestBankClient)gin.HandlerFunc{
	return func(c *gin.Context){
		uid,err:=getInt("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		test:=TestBank.Testdb{}
		err=c.ShouldBind(&test)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		if test.Typ>=TestBank.TypNum{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"invalid type",
			})
			return
		}
		if test.Discipline>=TestBank.DisciplineNum{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"invalid type",
			})
			return
		}
		
		test.Uploader=uid
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=t.Upload(ctx,&TestBank.Test{
			Id:test.Id,
			Typ:test.Typ,
			Content:test.Content,
			Ans:test.Ans,
			Name:test.Name,
			Uploader:test.Uploader,
			Discipline: test.Discipline,
		})
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
