package Handlers

import (
	"KeTangPai/services/DC/Exercise"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"time"
)

func Grade(e  Exercise.ExerciseClient) gin.HandlerFunc {
	return func(c *gin.Context){
		s:=c.DefaultQuery("sid","")
		if s==""{
			uid,err:=getInt("uid",c)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"error":err.Error(),
				})
				return
			}
			ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
			re,err:=e.GetScores(ctx,&Exercise.I{I: uid})
			if err!=nil {
				c.JSON(http.StatusInternalServerError,gin.H{
					"error":err.Error(),
				})
				return
			}
			var data []*Exercise.Submit
			for {
				rec,err:=re.Recv()
				if err==io.EOF{
					break
				}
				if err!=nil {
					c.JSON(http.StatusInternalServerError,gin.H{
						"error":err.Error(),
					})
					return
				}
				data=append(data,rec)
			}
			c.JSON(http.StatusOK,gin.H{
				"return":data,
			})
			return
		}

		sid,err:=strconv.Atoi(s)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=e.GetScore(ctx,&Exercise.I{I:int32(sid)})
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
