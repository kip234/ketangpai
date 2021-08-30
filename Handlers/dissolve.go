package Handlers

import (
	"KeTangPai/services/DC/Exercise"
	"KeTangPai/services/DC/UserCenter"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Dissolve(u  UserCenter.UserCenterClient,e  Exercise.ExerciseClient) gin.HandlerFunc {
	return func(c *gin.Context){
		classid,err:=getInt("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		_,err=u.DissolveClass(ctx,&UserCenter.Id{I: classid})
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		_,err=e.DelExercises(ctx,&Exercise.I{I: classid})
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
	}
}