package Handlers

import (
	"KeTangPai/services/DC/Exercise"
	"KeTangPai/services/DC/TestBank"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)
//不止作业，考试和作业没有区分
func Assign_homework(e  Exercise.ExerciseClient,t TestBank.TestBankClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, err := getUint("uid", c) //获取当前用户UID
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		classid, err := getUint("classid", c) //获取用户的班级
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		auto, ok := c.GetQuery("auto") //是否自动生成
		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "missing necessary parameter",
			})
			return
		}

		tmp := Exercise.Exercisedb{}
		err = c.ShouldBind(&tmp) //绑定前面传过来等作业信息
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		if tmp.Typ>=Exercise.TypNum{
			c.JSON(http.StatusBadRequest, gin.H{
				"error":"Invalid type",
			})
			return
		}
		//补全信息
		tmp.Ownerid = uid
		tmp.Classid = classid

		var re *Exercise.ExerciseData

		if auto == "1" {//自动生成
			subjective, ok := c.GetQuery("subjective")
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "missing necessary parameter",
				})
				return
			}
			sub,err:=strconv.Atoi(subjective)
			if err!=nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}

			objective, ok := c.GetQuery("objective") //是否自动生成
			if !ok {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "missing necessary parameter",
				})
				return
			}
			obj,err:=strconv.Atoi(objective)
			if err!=nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit)
			ids,err:=t.GenerateTest(ctx,&TestBank.Testconf{SubjectiveItem: uint32(sub),ObjectiveItem: uint32(obj),Discipline: tmp.Discipline})
			if err!=nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			tmp.Content=append(tmp.Content,ids.Tests...)//追加自动生成的内容
			tmp.Ans=ids.Ans
		}

		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit)
		re, err= e.AddExercise(ctx, &Exercise.ExerciseData{
			Classid:tmp.Classid,
			Ownerid:tmp.Ownerid,
			Content:tmp.Content,
			Typ:tmp.Typ,
			Begin:tmp.Begin,
			End:tmp.End,
			Duration:tmp.Duration,
			Name:tmp.Name,
			Ans: tmp.Ans,
		}) //添加记录
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"return":re,
		})
	}
}
