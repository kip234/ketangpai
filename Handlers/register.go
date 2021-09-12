//暂时没弄
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

//注册后会返回UID
func Register(uservice UserCenter.UserCenterClient,e Email.EmailClient ) gin.HandlerFunc {
	return func(c *gin.Context) {
		grader:=websocket.Upgrader{
			Subprotocols: []string{c.GetHeader("Sec-WebSocket-Protocol")},
			CheckOrigin: func(r *http.Request) bool {
			return true
		}}
		conn,err:=grader.Upgrade(c.Writer,c.Request,nil)
		if err!=nil{
			log.Println(err)
			c.JSON(200, gin.H{
				"status": 10022,
				"info":   "failed",
			})
			panic(err)
			return
		}
		defer conn.Close()

		user:=UserCenter.Userdb{}
		err=conn.ReadJSON(&user)
		if err!=nil {
			panic(err)
			conn.WriteJSON(err)
			return
		}
		//邮箱是否已经注册？
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit)
		ok,err:=uservice.UserIs_Exist(ctx,&UserCenter.S{S: user.Email})
		if err != nil {
			panic(err)
			conn.WriteJSON(err)
			return
		}
		if ok.Right {
			conn.WriteMessage(websocket.TextMessage,[]byte("该邮箱已经注册！"))
			return
		}

		//测试邮箱地址
		ctx,cancle:=context.WithTimeout(context.Background(),emailTimeLimit)
		rand.Seed(time.Now().Unix())
		code:=rand.Int()%1e6
		_,err=e.Send(context.Background(),&Email.Mail{Subject: subject,To: user.Email,Content: fmt.Sprintf(body,code)})
		if err != nil {
			panic(err)
			conn.WriteJSON(err)
			cancle()
			return
		}
		err=conn.WriteMessage(websocket.TextMessage,[]byte(fmt.Sprintf("验证码已发送给%s请及时提交验证码",user.Email)))
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

		//用户中心进行记录
		ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit)
		id,err:=uservice.CreatUser(ctx,&UserCenter.Uuser{
			//Id:user.Id,
			Pwd:user.Pwd,
			Email:user.Email,
		})
		if err!=nil{
			panic(err)
			conn.WriteJSON(err)
			cancle()
		}

		err=conn.WriteJSON(id)
		if err!=nil{
			panic(err)
			conn.WriteJSON(err)
			cancle()
		}
	}
}

