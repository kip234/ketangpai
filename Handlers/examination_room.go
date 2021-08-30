package Handlers

import (
	"KeTangPai/services/DC/Exercise"
	"KeTangPai/services/Filter"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Examination_room(e Exercise.ExerciseClient,f Filter.FilterClient) gin.HandlerFunc {
	return func(c *gin.Context){
		tmp,ok:=c.GetQuery("eid")
		if !ok||tmp==""{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"missing necessary parameters",
			})
			return
		}
		eid,err:=strconv.Atoi(tmp)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		re,err:=e.GetExercisec(ctx,&Exercise.I{I: int32(eid)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		//判断时间是否能对得上
		if re.Typ==Exercise.LimitedDate && re.End<time.Now().Unix(){//日期限制且已经超时
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"the test has expired",
			})
			return
		}
		if re.Typ==Exercise.TimeLimit{//限时型
			re.Begin=time.Now().Unix()
			re.End=re.Begin+re.Duration
		}

		uid,err:=getInt("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		//开始建立websocket
		upgrader := websocket.Upgrader{
			Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			c.JSON(200, gin.H{
				"status": 10022,
				"info": "failed",
			})
			return
		}

		conn.WriteJSON(re)
		ctx =context.Background()
		if re.Typ!=Exercise.Unlimited{//有时间限制
			ctx,_=context.WithDeadline(ctx,time.Unix(re.End,0))
		}

		go func(ct context.Context) {
			for{
				select{
					case <-ct.Done():
						conn.WriteJSON("finish ")
						conn.Close()
						return
					default:
						_,p,err:=conn.ReadMessage()
						if err!=nil{//认为客户断开
							conn.Close()
							return
						}
						re,err:=f.Process(context.Background(),&Filter.FilterData{Data: p})
						if err!=nil{//
							conn.WriteJSON("The server is faulty. The link is down")
							conn.Close()
							return
						}
						e.SubmitAns(context.Background(),&Exercise.Submit{Exerciseid: int32(eid),Uploaderid: uid,Contents: string(re.Data)})
						conn.WriteJSON("submit successfully")
				}
			}
		}(ctx)
	}
}