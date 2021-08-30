package Handlers

import (
	"KeTangPai/services/DC/Exercise"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Mark(e  Exercise.ExerciseClient) gin.HandlerFunc {
	return func(c *gin.Context){
		uid,err:=getInt("uid",c)
		if nil!=err {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		s:=c.DefaultQuery("id","")
		if s=="" {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"missing necessary parameters",
			})
			return
		}
		v:=c.DefaultQuery("v","")
		if v=="" {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"missing necessary parameters",
			})
			return
		}
		sid,err:=strconv.Atoi(s)
		if nil!=err {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		value,err:=strconv.Atoi(v)
		if nil!=err {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		_,err=e.SetScore(ctx, &Exercise.Score{Submitid: int32(sid),Value: int32(value),Judge: uid})
		if nil!=err {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
	}
}
