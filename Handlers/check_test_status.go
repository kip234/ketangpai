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
//查看作业状态-成绩分析？
func Check_test_status(e Exercise.ExerciseClient) gin.HandlerFunc {
	return func(c *gin.Context){
		s:=c.DefaultQuery("id","")//获取前面传来的考试ID
		if s==""{//如果没有就列出该班级的所有考试记录
			classid,err:=getInt("classid",c)
			if err!=nil {
				c.JSON(http.StatusBadRequest,gin.H{
					"error":err.Error(),
				})
				return
			}
			ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
			re,err:=e.GetExercises(ctx,&Exercise.I{I: classid})//获取考试记录
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
		//展示指定考试的情况
		id,err:=strconv.Atoi(s)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=e.GetClassSubmit(ctx,&Exercise.I{I: int32(id)})//获取提交记录
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		var data []*Exercise.Submit
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
	}
}
