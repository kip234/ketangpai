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

//发布一条言论
func (f *ForumService)Speak(c context.Context,in *Message) (*Message, error){
	Log.Send("Forum.Speak.info",in)
	select {
	case<-c.Done():
		Log.Send("Forum.Speak.error","timeout")
		return &Message{},errors.New("timeout")
	default:
	}
	tmp:=Messagedb{Time: in.Time,Owner: in.Owner,Tosb: in.Tosb,Content: in.Content,Classid: in.Classid}
	err:=f.db.Model(Messagedb{}).Create(&tmp).Error
	if err!=nil {
		Log.Send("Forum.Speak.error",err.Error())
		return &Message{},err
	}
	in.Id=tmp.Id
	return in,nil
}

//获取该UID用户收到的消息
func (f *ForumService)GetMessage(c context.Context,in *Uid) (*Messages, error){
	Log.Send("Forum.GetMessage.info",in)
	select {
	case<-c.Done():
		Log.Send("Forum.GetMessage.error","timeout")
		return &Messages{},errors.New("timeout")
	default:
	}
	re:=Messages{}
	err:=f.db.Model(Messagedb{}).Where("tosb=?",in.Id).Find(&re.M).Error
	if err!=nil {
		Log.Send("Forum.GetMessage.error",err.Error())
		return &Messages{},err
	}
	return &re,nil
}

//获取所有记录
func (f *ForumService)GetHistory(c context.Context,in *Classid) (*Messages, error){
	Log.Send("Forum.GetHistory.info",in)
	select {
	case<-c.Done():
		Log.Send("Forum.GetHistory.error","timeout")
		return &Messages{},errors.New("timeout")
	default:
	}
	re:=Messages{}
	err:=f.db.Model(Messagedb{}).Where("classid=?",in.Id).Find(&re.M).Error
	if err!=nil {
		Log.Send("Forum.GetHistory.error",err.Error())
		return &Messages{},err
	}
	return &re,nil
}

func (f *ForumService)mustEmbedUnimplementedForumServer(){}
