package KetangpaiDB

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const Addr		= ":8084"

//type of user
const (
	DefaultType = iota
	Teacher
	Student
	Administrator
	TypeNum
)

//sql
const (
	SqlName="ketangpai"
	SqlUserName="root"
	SqlUserPwd="123456"
	SqlAddr=":3306"
)

var sql *gorm.DB

type Memberdb struct{
	Uid 	uint32	`form:"uid" json:"uid" gorm:"not null;unique"`
	Classid uint32	`form:"classid" json:"classid" gorm:"not null;default:0;"`
	Type 	uint32	`form:"type" json:"type" gorm:"not null;default:0;"`
}

type Classdb struct {
	Classid  uint32   `form:"classid" json:"classid" gorm:"primaryKey"`
	Teacher  uint32   `form:"teacher" json:"teacher" gorm:"not null;unique"`
	Name     string  `binding:"required" form:"name" json:"name" gorm:"not null;unique"`
	Students []uint32 `binding:"required" form:"students" json:"students" gorm:"-"`
}

func InitGorm() {
	dsn := SqlUserName+":"+SqlUserPwd+"@tcp("+SqlAddr+")/"+SqlName+"?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	sql,err=gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	err=sql.AutoMigrate(Memberdb{},Classdb{})
	if err != nil {
		panic(err)
	}
}