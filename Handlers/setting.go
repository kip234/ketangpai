package Handlers

import (
	"KeTangPai/services/DC/UserCenter"
	"KeTangPai/services/Email"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

//前端要注意每一个字段,在GET时会发送完整的信息,POST时也要回传完整信息
func Setting(u UserCenter.UserCenterClient,e Email.EmailClient) gin.HandlerFunc {
	return func(c *gin.Context){
		uid,err:=getInt("uid",c)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		email,err:=getStr("email",c)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		if c.Request.Method=="GET"{
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
			defer conn.Close()
			newu:=UserCenter.Userdb{}
			err=conn.ReadJSON(&newu)
			if err!=nil {//
				conn.WriteJSON(err)
				return
			}
			if newu.Uid!=uint32(uid){//UID对不上
				conn.WriteMessage(websocket.TextMessage,[]byte("Uid conflict"))
				return
			}

			//测试验证码
			ctx,cancle:=context.WithTimeout(context.Background(),emailTimeLimit*time.Second)
			rand.Seed(time.Now().Unix())
			code:=rand.Int()%1e6
			_,err=e.Send(context.Background(),&Email.Mail{Subject: subject,To: email,Content: fmt.Sprintf(body,code)})
			err=conn.WriteMessage(websocket.TextMessage,[]byte(fmt.Sprintf("验证码已发送给%s请在%d秒内提交验证码",email,emailTimeLimit)))
			if err != nil {
				panic(err)
				conn.WriteJSON(err)
				cancle()
				return
			}

			//收取验证码
			channel:=make(chan int)
			close(channel)
			select {
			case <-ctx.Done():
				conn.WriteMessage(websocket.TextMessage, []byte("结束"))
				return
			case <-channel:
				_, p, err := conn.ReadMessage()
				if err != nil {
					panic(err)
					conn.WriteJSON(err)
					cancle()
				}
				cd,err:=strconv.Atoi(string(p))
				if err != nil {
					panic(err)
					conn.WriteJSON(err)
					cancle()
				}
				if cd==code{
					break
				}else{
					conn.WriteMessage(websocket.TextMessage,[]byte("验证码错误"))
				}
			}

			ctx,_=context.WithTimeout(context.Background(),emailTimeLimit*time.Second)
			re,err:=u.RefreshingUserData(ctx,&UserCenter.Uuser{
				Uid:newu.Uid,
				Name:newu.Name,
				Pwd:newu.Pwd,
				Email:newu.Email,
			})
			if err!=nil {
				c.JSON(http.StatusInternalServerError,gin.H{
					"error":err.Error(),
				})
				return
			}
			//JWT只用到UID所以不用刷新
			conn.WriteJSON(re)
			return
		}

		if c.Request.Method=="POST"{
			ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
			newu,err:=u.GetUserInfo(ctx,&UserCenter.Id{I: uint32(uid)})
			if err!=nil {
				c.JSON(http.StatusInternalServerError,gin.H{
					"error":err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK,gin.H{
				"return":newu,
			})
			return
		}
	}
}
