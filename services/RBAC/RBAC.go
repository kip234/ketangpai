package RBAC

import (
	"KeTangPai/Models/Redis"
	"KeTangPai/services/Log"
	"context"
	"errors"
	"gorm.io/gorm"
	"log"
)

type RBACService struct{
	db *gorm.DB
	pool *Redis.RedisPool
}

func newRBACService()*RBACService{//创建一个默认的服务
	InitGorm()//初始化MySQL连接
	re:=RBACService{}
	re.db=DB
	re.pool=&DefaultRedis
	re.pool.Init()
	return &re
}

//获取用户的角色
func (r *RBACService)GetRole(c context.Context,in *Uids) (*Roles, error){
	Log.Send("RBAC.GetRole.info",in)
	select {
	case<-c.Done():
		Log.Send("RBAC.GetRole.error","timeout")
		return &Roles{},errors.New("timeout")
	default:
	}
	paths:=make([]string,len(in.Uids))
	for i,_:=range in.Uids{
		var tmp string
		err:=r.db.Model(UserRoledb{}).Where("uid=?",in.Uids[i]).Select("role").Find(&tmp).Error
		if err!=nil {
			Log.Send("RBAC.GetRole.error",err.Error())
			return &Roles{},err
		}
		paths[i]=tmp
	}
	return &Roles{Roles: paths},nil
}

//获取子节点
func (r *RBACService)GetPaths(c context.Context,in *Path) (*Paths, error){
	Log.Send("RBAC.GetPaths.info",in)
	select {
	case<-c.Done():
		Log.Send("RBAC.GetPaths.error","timeout")
		return &Paths{},errors.New("timeout")
	default:
	}
	id:=Pathdb{Path: in.Path}
	err:=r.db.Model(Pathdb{}).Where(id).Find(&id).Error
	if err!=nil {
		Log.Send("RBAC.GetPaths.error",err.Error())
		return &Paths{},err
	}
	var paths []string
	err=r.db.Model(Pathdb{}).Where("fid=?",id.Id).Find(&paths).Error
	if err!=nil {
		Log.Send("RBAC.GetPaths.error",err.Error())
		return &Paths{},err
	}
	return &Paths{Paths: paths},nil
}

//添加角色
func (r *RBACService)AddRole(c context.Context,in *Roles) (*Empty, error){
	Log.Send("RBAC.AddRole.info",in)
	select {
	case<-c.Done():
		Log.Send("RBAC.AddRole.error","timeout")
		return &Empty{},errors.New("timeout")
	default:
	}

	err:=r.db.Transaction(func(tc *gorm.DB)error{
		for _,i:=range in.Roles{
			var role string
			err:=r.db.Model(Roledb{}).Select("role=?",i).Find(&role).Error
			if err!=nil {
				return err
			}
			log.Println("role:",role)
			if role == "0" {
				err=r.db.Model(Roledb{}).Create(Roledb{Role: i}).Error
				if err!=nil {
					return err
				}
			}
		}
		return nil
	})
	if err!=nil {
		Log.Send("RBAC.AddRole.error",err.Error())
	}
	return &Empty{},err
}

//添加路由分支 - 第一个元素为父节点
func (r *RBACService)RefreshPaths(c context.Context,in *Paths) (*Empty, error){
	Log.Send("RBAC.RefreshPaths.info",in)
	select {
	case<-c.Done():
		Log.Send("RBAC.RefreshPaths.error","timeout")
		return &Empty{},errors.New("timeout")
	default:
	}
	lenth:=len(in.Paths)
	if lenth<1 {
		return &Empty{},errors.New("nil")
	}
	id:=Pathdb{Path: in.Paths[0]}
	if id.Path==""{
		id.Id=0
	}else{
		err:=r.db.Model(Pathdb{}).Where(id).Find(&id).Error
		if err!=nil {
			Log.Send("RBAC.RefreshPaths.error",err.Error())
			return &Empty{},err
		}

		err=r.db.Model(Pathdb{}).Where(id).Find(&id).Error
		if err!=nil {
			Log.Send("RBAC.RefreshPaths.error",err.Error())
			return &Empty{},err
		}
	}

	lenth--
	for lenth>0 {
		tmp:=Pathdb{Fid: id.Id,Path: in.Paths[lenth]}
		err:=r.db.Model(Pathdb{}).Save(&tmp).Error
		if err!=nil {
			Log.Send("RBAC.RefreshPaths.error",err.Error())
			return &Empty{},err
		}
		lenth--
	}
	return &Empty{},nil
}

//刷新 用户-角色 关系
func (r *RBACService)RefreshUserRole(c context.Context,in *UsersRoles) (*Empty, error){
	Log.Send("RBAC.RefreshUserRole.info",in)
	select {
	case<-c.Done():
		Log.Send("RBAC.RefreshUserRole.error","timeout")
		return &Empty{},errors.New("timeout")
	default:
	}

	err:=r.db.Transaction(func(tx *gorm.DB)error{
		for i,_:=range in.Roles{
			if len(in.Roles[i])<1{
				return errors.New("the length of the role name is invalid")
			}
			var role string
			err:=r.db.Model(Roledb{}).Where("role=?",in.Roles[i]).Find(&role).Error
			if err!=nil {
				return err
			}
			if role=="" {
				return errors.New("known role name")
			}
			role=""
			err=r.db.Model(UserRoledb{}).Where("uid=?",in.Uids[i]).Select("role").Find(&role).Error
			if role==""{
				err=r.db.Model(UserRoledb{}).Create(UserRoledb{Uid: in.Uids[i],Role: in.Roles[i]}).Error
				if err!=nil {
					return err
				}
			}else{
				err=r.db.Model(UserRoledb{}).Where("uid=?",in.Uids[i]).Update("role",in.Roles[i]).Error
				if err!=nil{
					return err
				}
			}
		}
		return nil
	})

	if err!=nil {
		Log.Send("RBAC.RefreshUserRole.error",err.Error())
	}
	return &Empty{},err
}

//验证角色与路由的关系
func (r *RBACService)CheakRolePath(c context.Context,in *RolesPaths) (*Bools, error){
	Log.Send("RBAC.CheakRolePath.info",in)
	select {
	case<-c.Done():
		Log.Send("RBAC.CheakRolePath.error","timeout")
		return &Bools{},errors.New("timeout")
	default:
	}
	res:=make([]bool,len(in.Roles))
	for i,_:=range in.Roles{
		re,err:=r.pool.SISMEMBER(in.Roles[i],in.Paths[i])
		if err!=nil {
			Log.Send("RBAC.CheakRolePath.error",err.Error())
			return &Bools{},err
		}
		res[i]=re
	}

	return &Bools{Bools: res},nil
}

func (r *RBACService)Cache(c context.Context,in *Empty) (*Empty, error){
	Log.Send("RBAC.Cache.info",in)
	select {
	case<-c.Done():
		Log.Send("RBAC.Cache.error","timeout")
		return &Empty{},errors.New("timeout")
	default:
	}
	err:=r.cache()
	if err!=nil {
		Log.Send("RBAC.Cache.error",err.Error())
	}
	return &Empty{},err
}

func (r *RBACService)mustEmbedUnimplementedRBACServer(){}

//对MySQL数据进行缓存 角色与路由之间的继承
func (r *RBACService)cache()error{
	var roles []string
	err:=r.db.Model(Roledb{}).Find(&roles).Error
	if err!=nil {
		return err
	}
	for _,i:=range roles{
		err=r.pool.DEL(i)//删除所有记录
		if err!=nil {
			return err
		}
		err=r.pool.SADD(i,i)
		if err!=nil {
			return err
		}
		fid:=0
		err=r.db.Model(Pathdb{}).Where("path=?",i).Select("fid").Find(&fid).Error
		if err!=nil {
			return err
		}
		for fid>0{
			var path string
			err=r.db.Model(Pathdb{}).Where("id=?",fid).Select("path").Find(&path).Error
			if err!=nil {
				return err
			}
			err=r.pool.SADD(i,path)
			if err!=nil {
				return err
			}
			err=r.db.Model(Pathdb{}).Where("id=?",fid).Select("fid").Find(&fid).Error
			if err!=nil {
				return err
			}
		}
	}
	return nil
}