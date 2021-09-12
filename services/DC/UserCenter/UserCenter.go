package UserCenter

import (
	"KeTangPai/services/Log"
	"context"
	"errors"
	"gorm.io/gorm"
)

type UserCenterService struct {
	db *gorm.DB//MySQL连接
}

func newUserCenterService()*UserCenterService{//创建一个默认的服务
	InitGorm()//初始化MySQL连接
	return &UserCenterService{DB}
}

//创建一个用户
func(u *UserCenterService)CreatUser(c context.Context,user *Uuser) (*Uuser, error){
	Log.Send("UserCenter.CreatUser.info",user)
	select {
	case <-c.Done():
		Log.Send("UserCenter.CreatUser.error","timeout")
		return &Uuser{},errors.New("timeout")
	default:
	}
	tmp:=Userdb{Pwd: user.Pwd,Email: user.Email}
	err:=u.db.Model(&Userdb{}).Create(&tmp).Error
	if err!=nil {
		Log.Send("UserCenter.CreatUser.error",err.Error())
	}
	user.Id=tmp.Id
	return user, err
}

//依据UID获取用户所有信息
func(u *UserCenterService)GetUserInfo(c context.Context,in *Id) (*Uuser, error){
	Log.Send("UserCenter.GetUserInfo.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetUserInfo.error","timeout")
		return &Uuser{},errors.New("timeout")
	default:
	}
	user:=Uuser{}
	err:=u.db.Model(&Userdb{}).Where("id=?",in.I).Find(&user).Error
	if err!=nil {
		Log.Send("UserCenter.GetUserInfo.error",err.Error())
	}
	return &user, err
}

//凭借邮箱获取用户所有信息
func(u *UserCenterService)GetUserInfoByEmail(c context.Context,in *S) (*Uuser, error){
	Log.Send("UserCenter.GetUserInfoByEmail.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetUserInfoByEmail.error","timeout")
		return &Uuser{},errors.New("timeout")
	default:
	}
	user:=Uuser{}
	err:=u.db.Model(&Userdb{}).Where("email=?",in.S).Find(&user).Error
	if err!=nil {
		Log.Send("UserCenter.GetUserInfoByEmail.error",err.Error())
	}
	return &user, err
}

//凭借用户UID获取用户密码
func(u *UserCenterService)GetUserPwd(c context.Context,in *Id) (*S, error){
	Log.Send("UserCenter.GetUserPwd.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetUserPwd.error","timeout")
		return &S{S:""},errors.New("timeout")
	default:
	}
	var pwd string
	err:=u.db.Model(&Userdb{}).Where("id=?",in.I).Select("pwd").Find(&pwd).Error
	if err!=nil {
		Log.Send("UserCenter.GetUserPwd.error",err.Error())
	}
	return &S{S:pwd}, err
}

//凭借UID获取用户名
func(u *UserCenterService)GetUserName(c context.Context,in *Id) (*S, error){
	Log.Send("UserCenter.GetUserName.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetUserName.error","timeout")
		return &S{S:""},errors.New("timeout")
	default:
	}
	var name string
	err:=u.db.Model(&Userdb{}).Where("id=?",in.I).Select("name").Find(&name).Error
	if err!=nil {
		Log.Send("UserCenter.GetUserName.error",err.Error())
	}
	return &S{S:name},err
}

//凭借UID获取用户邮箱
func(u *UserCenterService)GetUserEmail(c context.Context,in *Id) (*S, error){
	Log.Send("UserCenter.GetUserEmail.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetUserEmail.error","timeout")
		return &S{S:""},errors.New("timeout")
	default:
	}
	var email string
	err:=u.db.Model(&Userdb{}).Where("id=?",in.I).Select("email").Find(&email).Error
	if err!=nil {
		Log.Send("UserCenter.GetUserEmail.error",err.Error())
	}
	return &S{S:email},err
}

//凭借邮箱判断用户是否存在
func(u *UserCenterService)UserIs_Exist(c context.Context,in *S) (*Right, error){
	Log.Send("UserCenter.UserIs_Exist.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.UserIs_Exist.error","timeout")
		return &Right{Right:false},errors.New("timeout")
	default:
	}
	tmpu:=Uuser{}
	err:=u.db.Model(&Userdb{}).Where("email=?",in.S).Find(&tmpu).Error
	if err!=nil {
		Log.Send("UserCenter.UserIs_Exist.error",err.Error())
		return &Right{Right:false},err
	}
	if tmpu.Id>0{
		return &Right{Right:true},err
	}
	return &Right{Right:false},err
}

//用传入的用户数据更新数据库数据
func (u *UserCenterService)RefreshingUserData(c context.Context,user *Uuser) (*Uuser, error) {
	Log.Send("UserCenter.RefreshingUserData.info",user)
	select {
	case <-c.Done():
		Log.Send("UserCenter.RefreshingUserData.error","timeout")
		return &Uuser{},errors.New("timeout")
	default:
	}
	if user.GetId()<1{
		return &Uuser{},errors.New("id illegal")
	}
	err:=u.db.Model(&Userdb{}).Where("id=?",user.Id).Updates(
		Userdb{
			Id: user.Id,
			Pwd: user.Pwd,
			Email:user.Email,
		}).Error
	if err!=nil{
		Log.Send("UserCenter.RefreshingUserData.error",err.Error())
	}
	return user,err
}




func(u *UserCenterService)mustEmbedUnimplementedUserCenterServer(){}
