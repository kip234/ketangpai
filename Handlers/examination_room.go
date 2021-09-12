package Handlers

import (
	"KeTangPai/services/DC/Exercise"
	"KeTangPai/services/DC/TestBank"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

func Examination_room(e Exercise.ExerciseClient,t TestBank.TestBankClient) gin.HandlerFunc {
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
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit)
		re,err:=e.GetExercisec(ctx,&Exercise.I{I: uint32(eid)})
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		anss:=make(map[uint32]string)
		err=json.Unmarshal(re.Ans,&anss)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			log.Println(err.Error())
			return
		}

		//判断时间是否能对得上
		if re.Typ==Exercise.LimitedDate && re.End<time.Now().Unix(){//日期限制且已经超时
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"the test has expired",
			})
			return
		}

		uid,err:=getUint("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		if re.Typ==Exercise.TimeLimit{//限时型
			re.Begin=time.Now().Unix()
			re.End=re.Begin+int64(re.Duration)
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
			var ans []Exercise.SubCont
			for{
				select{
					case <-ct.Done():
						conn.WriteJSON("finish ")
						conn.Close()
						return
					default:
						conn.ReadJSON(&ans)
						if err!=nil{//认为客户断开
							conn.Close()
							return
						}
						content:=make([]string,len(ans))
						for i,_:=range ans{
							if s,ok:=anss[ans[i].Testid];ok{
								if s==ans[i].Content{//答案正确
									ans[i].Status=true
								}else{
									ans[i].Status=false
								}
							}
							p,err:=json.Marshal(ans[i])
							if err!=nil{//认为客户断开
								conn.Close()
								return
							}
							content[i]=string(p)
						}
						ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit)
						e.SubmitAns(ctx,&Exercise.Submit{Exerciseid: uint32(eid),Uploaderid: uid,Contents: content})
						conn.WriteJSON("submit successfully")
				}
			}
		}(ctx)
	}
}