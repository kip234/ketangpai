package Email

import (
	"context"
	"errors"
	"log"
	"net/smtp"
	"strings"
)

type EmailService struct{
	host string//服务器地址
	user string//用户
	pwd string//密码
}

func newEmailService()*EmailService{
	return &DefaultEmailService
}

func (e *EmailService)Send(c context.Context,in *Mail) (re *Empty,err error){
	log.Printf("Send: %+v\n",in)
	defer func(){
		if err != nil{
			log.Printf("Send> %s\n",err.Error())
		}
	}()
	log.Printf("Send: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("Send> timeout\n")
		return &Empty{},errors.New("timeout")
	default:
	}
	hp := strings.Split(e.host, ":")
	auth := smtp.PlainAuth("", e.user, e.pwd, hp[0])
	msg := []byte("To: " + in.To + "\r\nFrom: " + e.user + "\r\nSubject: " +in.Subject+ "\r\n" + "Content-Type: text/html; charset=UTF-8" + "\r\n\r\n" + in.Content)
	to:=strings.Split(in.To,";")
	err=smtp.SendMail(e.host,auth,e.user,to,msg)
	return &Empty{},err
}

func (e *EmailService)mustEmbedUnimplementedEmailServer(){}
