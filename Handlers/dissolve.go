package Handlers

import (
	"KeTangPai/services/DC/Exercise"
	"KeTangPai/services/DC/KetangpaiDB"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Dissolve(u KetangpaiDB.KetangpaiDBClient,e  Exercise.ExerciseClient) gin.HandlerFunc {
	return func(c *gin.Context){
		classid,err:=getUint("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		_,err=u.DissolveClass(ctx,&KetangpaiDB.Classid{Classid: uint32(classid)})
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		_,err=e.DelExercises(ctx,&Exercise.I{I: uint32(classid)})
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
	}
}