package Handlers

import (
	"KeTangPai/services/DC/Exercise"
	"context"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)
//查看班级里面的作业
func Assignment(e Exercise.ExerciseClient) gin.HandlerFunc{
	return func(c *gin.Context){
		classid,err:=getUint("classid",c)//获取当前用户班级
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=e.GetExercises(ctx,&Exercise.I{I: classid})//获取该班级布置过的所有作业
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		var data []*Exercise.ExerciseData
		for  {
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
}
