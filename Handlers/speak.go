package Handlers

import (
	"KeTangPai/services/Filter"
	"KeTangPai/services/Forum"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Speak(f Forum.ForumClient,fi Filter.FilterClient)gin.HandlerFunc{
	return func(c *gin.Context){
		uid,err:=getUint("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		classid,err:=getUint("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		tmp:=Forum.Messagedb{}
		err=c.ShouldBind(&tmp)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}


		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		filtered,err:=fi.Process(ctx,&Filter.FilterData{Data: []byte(tmp.Content)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		tmp.Content=string(filtered.Data)
		tmp.Classid=uint32(classid)
		tmp.Owner=uint32(uid)
		tmp.Time=time.Now().Unix()
		ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=f.Speak(ctx,&Forum.Message{
			Id:tmp.Id,
			Owner:tmp.Owner,
			Tosb:tmp.Tosb,
			Content:tmp.Content,
			Time:tmp.Time,
			Classid:tmp.Classid,
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
