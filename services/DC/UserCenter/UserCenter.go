package UserCenter

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

type UserCenterService struct {
	db *gorm.DB//MySQL连接
}

func newUserCenterService()*UserCenterService{//创建一个默认的服务
	InitGorm()//初始化MySQL连接
	return &UserCenterService{DB}
}

func(u *UserCenterService)CreatUser(c context.Context,user *Uuser) (*Uuser, error){
	log.Printf("CreatUser: %+v\n",user)
	select {
	case <-c.Done():
		log.Printf("CreatUser> timeout\n")
		return &Uuser{},errors.New("timeout")
	default:
	}
	tmp:=Userdb{Name:user.Name,Pwd: user.Pwd,Type: user.Type,Classid: user.Classid,Email: user.Email}
	err:=u.db.Model(&Userdb{}).Create(&tmp).Error
	if err!=nil {
		log.Printf("CreatUser> %s\n",err.Error())
	}
	user.Uid=tmp.Uid
	return user, err
}

func(u *UserCenterService)GetUserInfo(c context.Context,in *Id) (*Uuser, error){
	log.Printf("GetUserInfo: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetUserInfo> timeout\n")
		return &Uuser{},errors.New("timeout")
	default:
	}
	user:=Uuser{}
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Find(&user).Error
	if err!=nil {
		log.Printf("GetUserInfo> %s\n",err.Error())
	}
	return &user, err
}

func(u *UserCenterService)GetUserInfoByEmail(c context.Context,in *S) (*Uuser, error){
	log.Printf("GetUserInfoByEmail: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetUserInfoByEmail> timeout\n")
		return &Uuser{},errors.New("timeout")
	default:
	}
	user:=Uuser{}
	err:=u.db.Model(&Userdb{}).Where("email=?",in.S).Find(&user).Error
	if err!=nil {
		log.Printf("GetUserInfoByEmail> %s\n",err.Error())
	}
	return &user, err
}

func(u *UserCenterService)GetUserPwd(c context.Context,in *Id) (*S, error){
	log.Printf("GetUserPwd: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetUserPwd> timeout\n")
		return &S{S:""},errors.New("timeout")
	default:
	}
	var pwd string
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("pwd").Find(&pwd).Error
	if err!=nil {
		log.Printf("GetUserPwd> %s\n",err.Error())
	}
	return &S{S:pwd}, err
}

func(u *UserCenterService)GetUserName(c context.Context,in *Id) (*S, error){
	log.Printf("GetUserName: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetUserName> timeout\n")
		return &S{S:""},errors.New("timeout")
	default:
	}
	var name string
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("name").Find(&name).Error
	if err!=nil {
		log.Printf("GetUserName> %s\n",err.Error())
	}
	return &S{S:name},err
}

func(u *UserCenterService)GetUserEmail(c context.Context,in *Id) (*S, error){
	log.Printf("GetUserEmail: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetUserEmail> timeout\n")
		return &S{S:""},errors.New("timeout")
	default:
	}
	var email string
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("email").Find(&email).Error
	if err!=nil {
		log.Printf("GetUserEmail> %s\n",err.Error())
	}
	return &S{S:email},err
}

func(u *UserCenterService)GetUserClass(c context.Context,in *Id) (*Id, error){
	log.Printf("GetUserClass: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetUserClass> timeout\n")
		return &Id{I:-1},errors.New("timeout")
	default:
	}
	var class int32
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("classid").Find(&class).Error
	if err!=nil {
		log.Printf("GetUserClass> %s\n",err.Error())
	}
	return &Id{I:class},err
}

func(u *UserCenterService)GetUserType(c context.Context,in *Id) (*Id, error){
	log.Printf("GetUserType: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetUserType> timeout\n")
		return &Id{I:-1},errors.New("timeout")
	default:
	}
	var typ int32
	err:=u.db.Model(&Userdb{}).Where("uid=?",in.I).Select("type").Find(&typ).Error
	if err!=nil {
		log.Printf("GetUserType> %s\n",err.Error())
	}
	return &Id{I:typ},err
}

func(u *UserCenterService)UserIs_Exist(c context.Context,in *S) (*Right, error){
	log.Printf("UserIs_Exist: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("UserIs_Exist> timeout\n")
		return &Right{Right:false},errors.New("timeout")
	default:
	}
	tmpu:=Uuser{}
	err:=u.db.Model(&Userdb{}).Where("email=?",in.S).Find(&tmpu).Error
	if err!=nil {
		log.Printf("UserIs_Exist> %s\n",err.Error())
		return &Right{Right:false},err
	}
	return &Right{Right:true},err
}

func (u *UserCenterService)RefreshingUserData(c context.Context,user *Uuser) (*Uuser, error) {
	log.Printf("RefreshingUserData: %+v\n",user)
	select {
	case <-c.Done():
		log.Printf("RefreshingUserData> timeout\n")
		return &Uuser{},errors.New("timeout")
	default:
	}
	if user.GetUid()<1{
		log.Printf("RefreshingUserData> uid illegal\n")
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
		log.Printf("RefreshingUserData> %s\n",err.Error())
	}
	return user,err
}

func (u *UserCenterService)CreateClass(c context.Context,class *Class) (*Class, error) {
	log.Printf("CreateClass: %+v\n",class)
	select {
	case <-c.Done():
		log.Printf("CreateClass> timeout\n")
		return &Class{},errors.New("timeout")
	default:
	}
	//检查该老师有没有旧班级,有的话就先解散
	err:=u.db.Transaction(func(tx *gorm.DB) error {
		var classid int32
		err:=tx.Model(Classdb{}).Where("Teacher=?",class.Teacher).Select("Classid").Find(&classid).Error
		if classid==0 {//没有直接退出进行下一步
			return nil
		}
		//如果有
		_,err=u.DissolveClass(context.Background(),&Id{I: classid})
		return err
	})
	if err!=nil {
		log.Printf("CreateClass> %s\n",err.Error())
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
		log.Printf("CreateClass> %s\n",err.Error())
	}
	return class,err
}

func (u *UserCenterService)GetClassInfo(c context.Context,in *Id) (*Class, error) {
	log.Printf("GetClassInfo: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetClassInfo> timeout\n")
		return &Class{},errors.New("timeout")
	default:
	}
	class:=Classdb{}
	err:=u.db.Model(&Classdb{}).Where("classid=?",in.I).Find(&class).Error//获取基础信息
	if err!=nil {
		log.Printf("GetClassInfo> %s\n",err.Error())
		return &Class{},err
	}
	err=u.db.Model(&Userdb{}).Select("uid").Where("classid=?",in.I).Find(&class.Students).Error//获取UID
	if err!=nil {
		log.Printf("GetClassInfo> %s\n",err.Error())
	}
	return &Class{Classid: class.Classid,Teacher: class.Teacher,Name: class.Name,Students: class.Students},err
}

func (u *UserCenterService)GetClassTeacher(c context.Context,in *Id) (*Id, error){
	log.Printf("GetClassTeacher: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetClassTeacher> timeout\n")
		return &Id{},errors.New("timeout")
	default:
	}
	var teacher int32
	err:=u.db.Model(&Classdb{}).Select("teacher").Where("classid",in.I).Find(&teacher).Error
	if err!=nil {
		log.Printf("GetClassTeacher> %s\n",err.Error())
	}
	return &Id{I:teacher},err
}

func (u *UserCenterService)GetClassName(c context.Context,in *Id) (*S, error){
	log.Printf("GetClassName: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetClassName> timeout\n")
		return &S{},errors.New("timeout")
	default:
	}
	var name string
	err:=u.db.Model(&Classdb{}).Where("classid=?",in.I).Select("name").Find(&name).Error
	if err!=nil {
		log.Printf("GetClassName> %s\n",err.Error())
	}
	return &S{S:name},err
}

func (u *UserCenterService)DissolveClass(c context.Context,in *Id) (*Empty, error) {
	log.Printf("DissolveClass: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("DissolveClass> timeout\n")
		return &Empty{},errors.New("timeout")
	default:
	}
	err:=u.db.Transaction(func(tx *gorm.DB) error {
		//for _,i:=range in.GetStudents(){
			err:=tx.Model(&Userdb{}).Where("Classid=?",in.I).Update("Classid",-1).Error
			if err!=nil {
				return err
			}
		//}
		err=tx.Model(Classdb{}).Delete(Classdb{Classid: in.I}).Error
		return err
	})
	if err!=nil {
		log.Printf("DissolveClass> %s\n",err.Error())
	}
	return &Empty{},err
}

func (u *UserCenterService)RefreshingClassData(c context.Context,in *Class) (*Class, error) {
	log.Printf("RefreshingClassData: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("RefreshingClassData> timeout\n")
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
		log.Printf("RefreshingClassData> %s\n",err.Error())
	}
	return in,err
}

func (u *UserCenterService)FireStudent(c context.Context,in *Id) (*Class, error){
	log.Printf("FireStudent: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("FireStudent> timeout\n")
		return &Class{},errors.New("timeout")
	default:
	}
	var classid int32
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
		log.Printf("FireStudent> %s\n",err.Error())
	}
	return re,err
}

func(u *UserCenterService)mustEmbedUnimplementedUserCenterServer(){}
