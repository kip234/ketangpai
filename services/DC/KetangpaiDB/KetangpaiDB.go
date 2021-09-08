package KetangpaiDB

import (
	"KeTangPai/services/Log"
	"context"
	"errors"
	"gorm.io/gorm"
)

type KetangpaiDBService struct{
	db *gorm.DB//MySQL连接
}

func newKetangpaiDBService() *KetangpaiDBService {
	InitGorm()
	return &KetangpaiDBService{sql}
}

//创建用户
func(u *KetangpaiDBService)CreateUser(c context.Context,in *Uid) (*Uid, error){
	Log.Send("KetangpaiDB.CreateUser.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.CreateUser.error","timeout")
		return &Uid{},errors.New("timeout")
	default:
	}
	//tmp:=Memberdb{Uid: in.Uid}
	err:=u.db.Model(Memberdb{}).Create(in).Error
	if err!=nil{
		Log.Send("KetangpaiDB.CreateUser.error",err.Error())
	}
	return in,err
}

//设置用户类型
func(u *KetangpaiDBService)SetType(c context.Context,in *Member) (*Member, error){
	Log.Send("KetangpaiDB.SetType.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.SetType.error","timeout")
		return &Member{},errors.New("timeout")
	default:
	}
	err:=u.db.Model(&Memberdb{}).Where("uid=?",in.Uid).Update("type",in.Type).Error
	if err!=nil{
		Log.Send("KetangpaiDB.SetType.error",err.Error())
	}
	return in,err
}

//凭借UID获取用户所处班级
func(u *KetangpaiDBService)GetUserClass(c context.Context,in *Uid) (*Classid, error){
	Log.Send("KetangpaiDB.GetUserClass.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.GetUserClass.error","timeout")
		return &Classid{Classid:0},errors.New("timeout")
	default:
	}
	var class uint32
	err:=u.db.Model(&Memberdb{}).Where("uid=?",in.Uid).Select("classid").Find(&class).Error
	if err!=nil {
		Log.Send("KetangpaiDB.GetUserClass.error",err.Error())
	}
	return &Classid{Classid:class},err
}

//凭借UID获取用户类型
func(u *KetangpaiDBService)GetUserType(c context.Context,in *Uid) (*Typecode, error){
	Log.Send("KetangpaiDB.GetUserType.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.GetUserType.error","timeout")
		return &Typecode{Typecode:0},errors.New("timeout")
	default:
	}
	var typ uint32
	err:=u.db.Model(&Memberdb{}).Where("uid=?",in.Uid).Select("type").Find(&typ).Error
	if err!=nil {
		Log.Send("KetangpaiDB.GetUserType.error",err.Error())
	}
	return &Typecode{Typecode:typ},err
}

///创建一个班级
func (u *KetangpaiDBService)CreateClass(c context.Context,class *Class) (*Class, error) {
	Log.Send("KetangpaiDB.CreateClass.info",class)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.CreateClass.error","timeout")
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
		_,err=u.DissolveClass(context.Background(),&Classid{Classid: classid})
		return err
	})
	if err!=nil {
		Log.Send("KetangpaiDB.CreateClass.error",err.Error())
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
			err=tx.Model(&Memberdb{}).Where("uid=?",i).Update("classid",class.GetClassid()).Error
			if err!=nil {
				return err
			}
		}
		//更改老师记录
		err=tx.Model(&Memberdb{}).Where("uid=?",class.Teacher).Update("classid",class.GetClassid()).Error
		return err
	})
	if err!=nil {
		Log.Send("KetangpaiDB.CreateClass.error",err.Error())
	}
	return class,err
}

//凭借classID获取班级信息
func (u *KetangpaiDBService)GetClassInfo(c context.Context,in *Classid) (*Class, error) {
	Log.Send("KetangpaiDB.GetClassInfo.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.GetClassInfo.error","timeout")
		return &Class{},errors.New("timeout")
	default:
	}
	class:=Classdb{}
	err:=u.db.Model(&Classdb{}).Where("classid=?",in.Classid).Find(&class).Error//获取基础信息
	if err!=nil {
		Log.Send("KetangpaiDB.GetClassInfo.error",err.Error())
		return &Class{},err
	}
	err=u.db.Model(&Memberdb{}).Select("uid").Where("classid=?",in.Classid).Find(&class.Students).Error//获取UID
	if err!=nil {
		Log.Send("KetangpaiDB.GetClassInfo.error",err.Error())
	}
	return &Class{Classid: class.Classid,Teacher: class.Teacher,Name: class.Name,Students: class.Students},err
}

//凭借classID获取班级教师
func (u *KetangpaiDBService)GetClassTeacher(c context.Context,in *Classid) (*Uid, error){
	Log.Send("KetangpaiDB.GetClassTeacher.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.GetClassTeacher.error","timeout")
		return &Uid{},errors.New("timeout")
	default:
	}
	var teacher uint32
	err:=u.db.Model(&Classdb{}).Select("teacher").Where("classid",in.Classid).Find(&teacher).Error
	if err!=nil {
		Log.Send("KetangpaiDB.GetClassTeacher.error",err.Error())
	}
	return &Uid{Uid:teacher},err
}

//凭借classID获取班级名
func (u *KetangpaiDBService)GetClassName(c context.Context,in *Classid) (*Classname, error){
	Log.Send("KetangpaiDB.GetClassName.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.GetClassName.error","timeout")
		return &Classname{},errors.New("timeout")
	default:
	}
	var name string
	err:=u.db.Model(&Classdb{}).Where("classid=?",in.Classid).Select("name").Find(&name).Error
	if err!=nil {
		Log.Send("KetangpaiDB.GetClassName.error",err.Error())
	}
	return &Classname{Name:name},err
}

//解散班级
func (u *KetangpaiDBService)DissolveClass(c context.Context,in *Classid) (*Empty, error) {
	Log.Send("KetangpaiDB.DissolveClass.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.DissolveClass.error","timeout")
		return &Empty{},errors.New("timeout")
	default:
	}
	err:=u.db.Transaction(func(tx *gorm.DB) error {
		//用-1表示班级已解散
		err:=tx.Model(&Memberdb{}).Where("Classid=?",in.Classid).Update("Classid",0).Error
		if err!=nil {
			return err
		}
		err=tx.Model(Classdb{}).Delete(Classdb{Classid: in.Classid}).Error
		return err
	})
	if err!=nil {
		Log.Send("KetangpaiDB.DissolveClass.error",err.Error())
	}
	return &Empty{},err
}

//依靠传入的数据刷新班级数据
func (u *KetangpaiDBService)RefreshingClassData(c context.Context,in *Class) (*Class, error) {
	Log.Send("KetangpaiDB.RefreshingClassData.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.RefreshingClassData.error","timeout")
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
		Log.Send("KetangpaiDB.RefreshingClassData.error",err.Error())
	}
	return in,err
}

//凭借UID将该用户从其所属班级中移除
func (u *KetangpaiDBService)FireStudent(c context.Context,in *Uid) (*Class, error){
	Log.Send("KetangpaiDB.FireStudent.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.FireStudent.error","timeout")
		return &Class{},errors.New("timeout")
	default:
	}
	var classid uint32
	var re *Class
	err:=u.db.Transaction(func(tx *gorm.DB)error{
		err:=tx.Model(&Memberdb{}).Where("uid=?",in.Uid).Select("classid").Find(&classid).Error//获取classID
		if err!=nil {
			return err
		}
		err=tx.Model(&Memberdb{}).Where("uid=?",in.Uid).Update("classid",0).Error//退出班级
		if err!=nil {
			return err
		}
		re,err=u.GetClassInfo(context.Background(),&Classid{Classid: classid})
		return err
	})
	if err!=nil {
		Log.Send("KetangpaiDB.FireStudent.error",err.Error())
	}
	return re,err
}

//添加某个学生到班级
func (u *KetangpaiDBService)AddStudent(c context.Context,in *Member) (*Class, error){
	Log.Send("KetangpaiDB.AddStudent.info",in)
	select {
	case <-c.Done():
		Log.Send("KetangpaiDB.AddStudent.error","timeout")
		return &Class{},errors.New("timeout")
	default:
	}
	err:=u.db.Model(Memberdb{}).Where("uid=?",in.Uid).Update("classid",in.Classid).Error
	if err!=nil{
		Log.Send("KetangpaiDB.AddStudent.error",err.Error())
		return &Class{},err
	}
	re,err:=u.GetClassInfo(context.Background(),&Classid{Classid: in.Classid})
	if err!=nil{
		Log.Send("KetangpaiDB.AddStudent.error",err.Error())
	}
	return re,err
}

func (u *KetangpaiDBService)mustEmbedUnimplementedKetangpaiDBServer(){}