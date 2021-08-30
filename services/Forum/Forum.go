package Forum

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

type ForumService struct {
	db *gorm.DB
}

func newForumService()*ForumService{//创建一个默认的服务
	InitGorm()//初始化MySQL连接
	return &ForumService{DB}
}

func (f *ForumService)Speak(c context.Context,in *Message) (*Message, error){
	log.Printf("Speak: %+v\n",in)
	select {
	case<-c.Done():
		log.Printf("Speak> timeout\n")
		return &Message{},errors.New("timeout")
	default:
	}
	tmp:=Messagedb{Time: in.Time,Owner: in.Owner,Tosb: in.Tosb,Content: in.Content,Classid: in.Classid}
	err:=f.db.Model(Messagedb{}).Create(&tmp).Error
	if err!=nil {
		log.Printf("Speak> %s\n",err.Error())
		return &Message{},err
	}
	in.Id=tmp.Id
	return in,nil
}

func (f *ForumService)GetMessage(c context.Context,in *Uid) (*Messages, error){
	log.Printf("GetMessage: %+v\n",in)
	select {
	case<-c.Done():
		log.Printf("GetMessage> timeout\n")
		return &Messages{},errors.New("timeout")
	default:
	}
	re:=Messages{}
	err:=f.db.Model(Messagedb{}).Where("tosb=?",in.Id).Find(&re.M).Error
	if err!=nil {
		log.Printf("GetMessage> %s\n",err.Error())
		return &Messages{},err
	}
	return &re,nil
}

func (f *ForumService)GetHistory(c context.Context,in *Classid) (*Messages, error){
	log.Printf("GetHistory: %+v\n",in)
	select {
	case<-c.Done():
		log.Printf("GetHistory> timeout\n")
		return &Messages{},errors.New("timeout")
	default:
	}
	re:=Messages{}
	err:=f.db.Model(Messagedb{}).Where("classid=?",in.Id).Find(&re.M).Error
	if err!=nil {
		log.Printf("GetHistory> %s\n",err.Error())
		return &Messages{},err
	}
	return &re,nil
}

func (f *ForumService)mustEmbedUnimplementedForumServer(){}
