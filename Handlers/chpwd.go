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

func Chpwd(u UserCenter.UserCenterClient,e Email.EmailClient)gin.HandlerFunc{
	return func(c *gin.Context){
		email,err:=getStr("email",c)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		id,err:=getUint("id",c)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}

		pwd,ok:=c.GetQuery("pwd")
		if !ok {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"missing necessary parameters",
			})
			return
		}
		if len(pwd)<6{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"the password length is invalid",
			})
			return
		}

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

		//测试验证码
		ctx,cancle:=context.WithTimeout(context.Background(),emailTimeLimit)
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
			Id:id,
			Pwd:pwd,
			Email:email,
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
}
