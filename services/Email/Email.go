package Email

import (
	"KeTangPai/services/Log"
	"context"
	"errors"
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

//发送邮件
func (e *EmailService)Send(c context.Context,in *Mail) (re *Empty,err error){
	Log.Send("Email.Send.info",in)
	select {
	case <-c.Done():
		Log.Send("Email.Send.error","timeout")
		return &Empty{},errors.New("timeout")
	default:
	}
	hp := strings.Split(e.host, ":")
	auth := smtp.PlainAuth("", e.user, e.pwd, hp[0])
	msg := []byte("To: " + in.To + "\r\nFrom: " + e.user + "\r\nSubject: " +in.Subject+ "\r\n" + "Content-Type: text/html; charset=UTF-8" + "\r\n\r\n" + in.Content)
	to:=strings.Split(in.To,";")
	err=smtp.SendMail(e.host,auth,e.user,to,msg)
	if err!=nil {
		Log.Send("Email.Send.error",err.Error())
	}
	return &Empty{},err
}

func (e *EmailService)mustEmbedUnimplementedEmailServer(){}
