package RBAC

import (
	"KeTangPai/Models/Redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pathdb struct{
	Id	uint32	`form:"id" json:"id" gorm:"primaryKey"`
	Path string	`binding:"required" form:"path" json:"path" gorm:"not null;unique"`
	Fid uint32	`form:"fid" json:"fid" gorm:"default:0"`
}

type UserRoledb struct{
	Uid uint32	`binding:"required" form:"uid" json:"uid" gorm:"not null;unique"`
	Role string	`binding:"required" form:"role" json:"role" gorm:"not null"`
}

type Roledb struct {
	Role string `gorm:"not null;unique"`
}

var DB *gorm.DB

const Addr="localhost:8092"

//sql
const (
	SqlName="ketangpai"
	SqlUserName="root"
	SqlUserPwd="root"
	SqlAddr=":3306"
)

var DefaultRedis = Redis.RedisPool{
	Read 		:"localhost:6379",
	Write 		:"localhost:6379",
	IdLeTimeout	:5,
	MaxIdle		:20,
	MaxActive	:8,
}

func InitGorm() {
	dsn := SqlUserName+":"+SqlUserPwd+"@tcp("+SqlAddr+")/"+SqlName+"?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB,err=gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	err=DB.AutoMigrate(UserRoledb{},Pathdb{},Roledb{})
	if err != nil {
		panic(err)
	}
}