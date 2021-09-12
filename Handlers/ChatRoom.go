package Handlers

import (
	"KeTangPai/Models"
	"KeTangPai/services/Filter"
	"KeTangPai/services/RankingList"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func ChatRoom(f Filter.FilterClient,rooms map[uint32]*Models.Room,r RankingList.RankingListClient)gin.HandlerFunc{
	return func(c *gin.Context){
		uid,err:=getUint("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}
		classid,err:=getUint("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}
		uname,err:=getStr("name",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}
		classname,err:=getStr("classname",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
		}

		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		//加入班级
		if _,ok:=rooms[classid];!ok{
			rooms[classid]=Models.NewRoom(classname)
			rooms[classid].Add(uid,conn)
			go rooms[classid].Run()
		}else{
			rooms[classid].Add(uid,conn)
		}

		go func(){
			defer func(){
				conn.Close()
				rooms[classid].Del(uid)
				if rooms[classid].ConnNum()==0{//
					delete(rooms,classid)//清理聊天室
					ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
					_,err=r.Dellist(ctx,&RankingList.Listname{Name: classname})//清空榜单
				}
				//Log.Printf("ChatRoom> leave %d\n",rooms[classid].ConnNum())
			}()
			in:=Models.In{}
			out:=Models.Out{
				Classid: classid,
				Classname: classname,
				Uid:uid,
				Uname: uname,
			}
			for {
				err:=conn.ReadJSON(&in)
				if err!=nil {
					log.Printf("ChatRoom> %s\n",err.Error())
					break
				}
				ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
				_,err=r.Flushlist(ctx,&RankingList.Flushin{Key: classname,Increment: 1,Member: uname})
				if err!=nil {
					log.Printf("ChatRoom> %s\n",err.Error())
					break
				}
				ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
				re,err:=f.Process(ctx,&Filter.FilterData{Data: []byte(in.Content)})
				if err!=nil {
					log.Printf("ChatRoom> %s\n",err.Error())
					break
				}
				out.Content=string(re.Data)
				out.In.To=in.To
				ctx,_=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
				list,err:=r.Getlistinfo(ctx,&RankingList.Listname{Name: classname})
				if err!=nil {
					log.Printf("ChatRoom> %s\n",err.Error())
					break
				}
				out.Ranks=list.List
				rooms[classid].Chan()<-&out
			}
		}()
	}
}
