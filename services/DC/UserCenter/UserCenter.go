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
	tmp:=Userdb{Name:user.Name,Pwd: user.Pwd,Type: user.Type,Classid: user.Classid,Email: user.Email}
	err:=u.db.Model(&Userdb{}).Create(&tmp).Error
	if err!=nil {
		Log.Send("UserCenter.CreatUser.error",err.Error())
	}
	user.Uid=tmp.Uid
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
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Find(&user).Error
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
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("pwd").Find(&pwd).Error
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
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("name").Find(&name).Error
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
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("email").Find(&email).Error
	if err!=nil {
		Log.Send("UserCenter.GetUserEmail.error",err.Error())
	}
	return &S{S:email},err
}

//凭借UID获取用户所处班级
func(u *UserCenterService)GetUserClass(c context.Context,in *Id) (*Id, error){
	Log.Send("UserCenter.GetUserClass.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetUserClass.error","timeout")
		return &Id{I:0},errors.New("timeout")
	default:
	}
	var class uint32
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("classid").Find(&class).Error
	if err!=nil {
		Log.Send("UserCenter.GetUserClass.error",err.Error())
	}
	return &Id{I:class},err
}

//凭借UID获取用户类型
func(u *UserCenterService)GetUserType(c context.Context,in *Id) (*Id, error){
	Log.Send("UserCenter.GetUserType.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetUserType.error","timeout")
		return &Id{I:0},errors.New("timeout")
	default:
	}
	var typ uint32
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("type").Find(&typ).Error
	if err!=nil {
		Log.Send("UserCenter.GetUserType.error",err.Error())
	}
	return &Id{I:typ},err
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
	if tmpu.Uid>0{
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
	if user.GetUid()<1{
		//log.Printf("RefreshingUserData> uid illegal\n")
		return &Uuser{},errors.New("uid illegal")
	}
	err:=u.db.Model(&Userdb{}).Where("uid=?",user.Uid).Updates(
		Userdb{
			Uid: user.Uid,
			Name: user.Name,
			Pwd: user.Pwd,
			Type: user.Type,
			Classid: user.Classid,
			Email:user.Email,
		}).Error
	if err!=nil{
		Log.Send("UserCenter.RefreshingUserData.error",err.Error())
	}
	return user,err
}

///创建一个班级
func (u *UserCenterService)CreateClass(c context.Context,class *Class) (*Class, error) {
	Log.Send("UserCenter.CreateClass.info",class)
	select {
	case <-c.Done():
		Log.Send("UserCenter.CreateClass.error","timeout")
		return &Class{},errors.New("timeout")
	default:
	}
	//检查该老师有没有旧班级,有的话就先解散
	err:=u.db.Transaction(func(tx *gorm.DB) error {
		var classid uint32
		err:=tx.Model(Classdb{}).Where("Teacher=?",class.Teacher).Select("Classid").Find(&classid).Error
		if classid==0 {//没有直接退出进行下一步
			return nil
		}
		//如果有
		_,err=u.DissolveClass(context.Background(),&Id{I: classid})
		return err
	})
	if err!=nil {
		Log.Send("UserCenter.CreateClass.error",err.Error())
		return &Class{},err
	}


	//创建新班级
	err=u.db.Transaction(func(tx *gorm.DB) error {
		tmp:=Classdb{Teacher: class.Teacher,Name: class.Name}
		err:=tx.Model(&Classdb{}).Create(&tmp).Error
		if err!=nil {
			return err
		}
		class.Classid=tmp.Classid
		for _,i:=range class.Students{//更改学生记录
			err=tx.Model(&Userdb{}).Where("uid=?",i).Update("classid",class.GetClassid()).Error
			if err!=nil {
				return err
			}
		}
		//更改老师记录
		err=tx.Model(&Userdb{}).Where("uid=?",class.Teacher).Update("classid",class.GetClassid()).Error
		return err
	})
	if err!=nil {
		Log.Send("UserCenter.CreateClass.error",err.Error())
	}
	return class,err
}

//凭借classID获取班级信息
func (u *UserCenterService)GetClassInfo(c context.Context,in *Id) (*Class, error) {
	Log.Send("UserCenter.GetClassInfo.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetClassInfo.error","timeout")
		return &Class{},errors.New("timeout")
	default:
	}
	class:=Classdb{}
	err:=u.db.Model(&Classdb{}).Where("classid=?",in.I).Find(&class).Error//获取基础信息
	if err!=nil {
		Log.Send("UserCenter.GetClassInfo.error",err.Error())
		return &Class{},err
	}
	err=u.db.Model(&Userdb{}).Select("uid").Where("classid=?",in.I).Find(&class.Students).Error//获取UID
	if err!=nil {
		Log.Send("UserCenter.GetClassInfo.error",err.Error())
	}
	return &Class{Classid: class.Classid,Teacher: class.Teacher,Name: class.Name,Students: class.Students},err
}

//凭借classID获取班级教师
func (u *UserCenterService)GetClassTeacher(c context.Context,in *Id) (*Id, error){
	Log.Send("UserCenter.GetClassTeacher.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetClassTeacher.error","timeout")
		return &Id{},errors.New("timeout")
	default:
	}
	var teacher uint32
	err:=u.db.Model(&Classdb{}).Select("teacher").Where("classid",in.I).Find(&teacher).Error
	if err!=nil {
		Log.Send("UserCenter.GetClassTeacher.error",err.Error())
	}
	return &Id{I:teacher},err
}

//凭借classID获取班级名
func (u *UserCenterService)GetClassName(c context.Context,in *Id) (*S, error){
	Log.Send("UserCenter.GetClassName.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.GetClassName.error","timeout")
		return &S{},errors.New("timeout")
	default:
	}
	var name string
	err:=u.db.Model(&Classdb{}).Where("classid=?",in.I).Select("name").Find(&name).Error
	if err!=nil {
		Log.Send("UserCenter.GetClassName.error",err.Error())
	}
	return &S{S:name},err
}

//解散班级
func (u *UserCenterService)DissolveClass(c context.Context,in *Id) (*Empty, error) {
	Log.Send("UserCenter.DissolveClass.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.DissolveClass.error","timeout")
		return &Empty{},errors.New("timeout")
	default:
	}
	err:=u.db.Transaction(func(tx *gorm.DB) error {
		//用-1表示班级已解散
		err:=tx.Model(&Userdb{}).Where("Classid=?",in.I).Update("Classid",-1).Error
		if err!=nil {
			return err
		}
		err=tx.Model(Classdb{}).Delete(Classdb{Classid: in.I}).Error
		return err
	})
	if err!=nil {
		Log.Send("UserCenter.DissolveClass.error",err.Error())
	}
	return &Empty{},err
}

//依靠传入的数据刷新班级数据
func (u *UserCenterService)RefreshingClassData(c context.Context,in *Class) (*Class, error) {
	Log.Send("UserCenter.RefreshingClassData.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.RefreshingClassData.error","timeout")
		return &Class{},errors.New("timeout")
	default:
	}
	err:=u.db.Model(Classdb{}).Where("classid=?",in.GetClassid()).Updates(
		Classdb{
			Classid: in.Classid,
			Teacher: in.Teacher,
			Name: in.Name,
		}).Error
	if err!=nil {
		Log.Send("UserCenter.RefreshingClassData.error",err.Error())
	}
	return in,err
}

//凭借UID将该用户从其所属班级中移除
func (u *UserCenterService)FireStudent(c context.Context,in *Id) (*Class, error){
	Log.Send("UserCenter.FireStudent.info",in)
	select {
	case <-c.Done():
		Log.Send("UserCenter.FireStudent.error","timeout")
		return &Class{},errors.New("timeout")
	default:
	}
	var classid uint32
	var re *Class
	err:=u.db.Transaction(func(tx *gorm.DB)error{
		err:=tx.Model(&Userdb{}).Where("uid=?",in.I).Select("classid").Find(&classid).Error//获取classID
		if err!=nil {
			return err
		}
		err=tx.Model(&Userdb{}).Where("uid=?",in.I).Update("classid",-1).Error//退出班级
		if err!=nil {
			return err
		}
		re,err=u.GetClassInfo(context.Background(),&Id{I: classid})
		return err
	})
	if err!=nil {
		Log.Send("UserCenter.FireStudent.error",err.Error())
	}
	return re,err
}

func(u *UserCenterService)mustEmbedUnimplementedUserCenterServer(){}
