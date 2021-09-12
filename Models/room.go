package Models

import (
	"github.com/gorilla/websocket"
)

//实时交流用的房间结构

type Room struct{
	num uint32
	name string//标识
	in chan *Out//等待输入消息
	links map[uint32] *websocket.Conn
}

func (r *Room)Add(uid uint32,conn *websocket.Conn){
	if _,ok:=r.links[uid];!ok{
		r.num+=1
	}else{
		r.links[uid].WriteMessage(websocket.TextMessage,[]byte("The account is online somewhere else. The link is about to be disconnected"))
		r.links[uid].Close()
	}
	r.links[uid]=conn
}

func (r *Room)Del(uid uint32){
	if _,ok:=r.links[uid];ok{
		r.num-=1
	}
	delete(r.links,uid)
}

func (r *Room)Chan() chan *Out{
	return r.in
}

func (r *Room)IsExit(uid uint32) (ok bool) {
	_,ok=r.links[uid]
	return
}

func (r *Room)Run(){
	for {
		if r.num==0{//没有连接的时候自动退出
			break
		}
		select {
		case data := <-r.in:
			data.Online=r.num//刷新在线人数
			if len(data.To)!=0 {
				for _,i:=range data.To{
					if _,ok:=r.links[i];!ok{
						continue
					}
					err:=r.links[i].WriteJSON(data)
					if err!=nil {//认为用户断开
						r.Del(i)
					}
				}
				err:=r.links[data.Uid].WriteJSON(data)
				if err!=nil {//认为用户断开
					r.Del(data.Uid)
				}
			}else{
				for uid,i:=range r.links{
					err:=i.WriteJSON(data)
					if err!=nil {//认为用户断开
						r.Del(uid)
					}
				}
			}
		}
	}
}

func (r *Room)ConnNum()uint32{
	return r.num
}

func NewRoom(Name string) *Room {
	return &Room{
		num: 0,
		name: Name,
		in:make(chan *Out),
		links:make(map[uint32] *websocket.Conn),
	}
}
