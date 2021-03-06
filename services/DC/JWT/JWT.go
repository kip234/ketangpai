package JWT

import (
	"KeTangPai/Models/Redis"
	"KeTangPai/services/Log"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Iss string 	`json:"iss"`//签发人
	Exp uint 	`json:"exp"`//过期时间
	Sub string 	`json:"sub"`//主题
	Aud uint32 	`json:"aud"`//用户ID
	Nbf uint 	`json:"ndf"`//生效时间
	Iat int64 	`json:"iat"`//签发时间
	Jti uint 	`json:"jti"`//编号
}

type Jwt struct{
	Header Header
	Payload Payload
	Secret string
}

type JwtService struct{
	redis *Redis.RedisPool
}

func newJwtService()*JwtService{
	DefaultRedis.Init()
	jwt:=JwtService{redis: &DefaultRedis}
	return &jwt
}

func (j *JwtService) RefreshToken(c context.Context,U *Juser) (*Token, error) {
	Log.Send("JWT.RefreshToken.info",U)
	select {
	case <-c.Done():
		Log.Send("JWT.RefreshToken.error","timeout")
		return &Token{},errors.New("timeout")
	default:
	}
	//给部分字段赋上特殊值
	DefaultJwt.Payload.Aud=U.Id
	DefaultJwt.Payload.Iat=time.Now().Unix()
	header,err:=json.Marshal(DefaultJwt.Header)
	if err!=nil {
		Log.Send("JWT.RefreshToken.error",err.Error())
		return &Token{Content: ""},err
	}
	Header1:=base64.StdEncoding.EncodeToString(header)
	payload,err:=json.Marshal(DefaultJwt.Payload)
	if err!=nil {
		Log.Send("JWT.RefreshToken.error",err.Error())
		return &Token{Content: ""},err
	}

	Payload1:=base64.StdEncoding.EncodeToString(payload)
	hash := hmac.New(sha256.New,[]byte(DefaultJwt.Secret))
	hash.Write(
		[]byte(Header1+ "."+
			Payload1+ "."))
	t:=Token{Content: Header1+"."+Payload1+"."+base64.StdEncoding.EncodeToString(hash.Sum(nil))}
	err=j.redis.SET(U.Id,t.Content)
	if err!=nil {
		Log.Send("JWT.RefreshToken.error",err.Error())
	}
	return &t,err
}

func (j *JwtService) CheckToken(c context.Context,t *Token) (*Juser, error) {
	Log.Send("JWT.CheckToken.info",t)
	select {
	case <-c.Done():
		Log.Send("JWT.CheckToken.error","timeout")
		return &Juser{},errors.New("timeout")
	default:
	}
	hps:=strings.Split(t.Content,".")//分割token的三部分
	if len(hps)!=3 {//长度不够？
		return &Juser{},fmt.Errorf("RefreshHP Signature error")
	}

	p,err:=base64.StdEncoding.DecodeString(hps[1])//反序列化paylo
	if err!=nil{
		Log.Send("JWT.CheckToken.error",err.Error())
		return &Juser{},err
	}
	err = json.Unmarshal(p,&DefaultJwt.Payload)
	r,err:=j.redis.GET(strconv.Itoa(int(DefaultJwt.Payload.Aud)))//查找token记录
	if err!=nil{
		Log.Send("JWT.CheckToken.error",err.Error())
		return &Juser{},err
	}
	if r!=t.Content{//与记录不符
		//log.Printf("CheckToken> inconsistent with record\n")
		return &Juser{},errors.New("inconsistent with record")
	}
	return &Juser{Id: DefaultJwt.Payload.Aud},err
}

//清除JWT记录
func (j *JwtService) DelJwt(c context.Context,in *Juser) (*Token, error){
	Log.Send("JWT.DelJwt.info",in)
	select {
	case <-c.Done():
		Log.Send("JWT.DelJwt.error","timeout")
		return &Token{},errors.New("timeout")
	default:
	}
	err:=j.redis.DEL(strconv.Itoa(int(in.Id)))
	if err!=nil {
		Log.Send("JWT.DelJwt.error",err.Error())
	}
	return &Token{},err
}

func(j *JwtService)mustEmbedUnimplementedJWTServer(){}