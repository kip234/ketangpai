package Forum

import (
	"KeTangPai/services/Log"
	"context"
	"errors"
	"gorm.io/gorm"
)

type ForumService struct {
	db *gorm.DB
}

func newForumService()*ForumService{//创建一个默认的服务
	InitGorm()//初始化MySQL连接
	return &ForumService{DB}
}

func (f *ForumService)Speak(c context.Context,in *Message) (*Message, error){
	Log.Send("Forum.Speak.info",in)
	//log.Printf("Speak: %+v\n",in)
	select {
	case<-c.Done():
		Log.Send("Forum.Speak.error","timeout")
		//log.Printf("Speak> timeout\n")
		return &Message{},errors.New("timeout")
	default:
	}
	tmp:=Messagedb{Time: in.Time,Owner: in.Owner,Tosb: in.Tosb,Content: in.Content,Classid: in.Classid}
	err:=f.db.Model(Messagedb{}).Create(&tmp).Error
	if err!=nil {
		Log.Send("Forum.Speak.error",err.Error())
		//log.Printf("Speak> %s\n",err.Error())
		return &Message{},err
	}
	in.Id=tmp.Id
	return in,nil
}

func (f *ForumService)GetMessage(c context.Context,in *Uid) (*Messages, error){
	Log.Send("Forum.GetMessage.info",in)
	//log.Printf("GetMessage: %+v\n",in)
	select {
	case<-c.Done():
		Log.Send("Forum.GetMessage.error","timeout")
		//log.Printf("GetMessage> timeout\n")
		return &Messages{},errors.New("timeout")
	default:
	}
	re:=Messages{}
	err:=f.db.Model(Messagedb{}).Where("tosb=?",in.Id).Find(&re.M).Error
	if err!=nil {
		Log.Send("Forum.GetMessage.error",err.Error())
		//log.Printf("GetMessage> %s\n",err.Error())
		return &Messages{},err
	}
	return &re,nil
}

func (f *ForumService)GetHistory(c context.Context,in *Classid) (*Messages, error){
	Log.Send("Forum.GetHistory.info",in)
	//log.Printf("GetHistory: %+v\n",in)
	select {
	case<-c.Done():
		Log.Send("Forum.GetHistory.error","timeout")
		//log.Printf("GetHistory> timeout\n")
		return &Messages{},errors.New("timeout")
	default:
	}
	re:=Messages{}
	err:=f.db.Model(Messagedb{}).Where("classid=?",in.Id).Find(&re.M).Error
	if err!=nil {
		Log.Send("Forum.GetHistory.error",err.Error())
		//log.Printf("GetHistory> %s\n",err.Error())
		return &Messages{},err
	}
	return &re,nil
}

func (f *ForumService)mustEmbedUnimplementedForumServer(){}
