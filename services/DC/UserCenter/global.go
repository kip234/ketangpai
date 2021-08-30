package UserCenter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

const Addr="localhost:8086"

//sql
const (
	SqlName="ketangpai"
	SqlUserName="root"
	SqlUserPwd="root"
	SqlAddr="127.0.0.1:3306"
)

//type of user
const (
	DefaultType = iota
	Teacher
	Student
	Administrator
	TypeNum
)

//储存对象
type Userdb struct {
	Uid     int32   `form:"uid" json:"uid" gorm:"primaryKey"`
	Name    string	`form:"name" json:"name" gorm:"not null"`
	Pwd     string	`binding:"required" form:"pwd" json:"pwd" gorm:"not null"`
	Type    uint32	`form:"type" json:"type" gorm:"not null;default:0"`
	Classid int32	`form:"classid" json:"Classid"`
	Email string    `form:"email" json:"email" gorm:"not null;unique"`
}

type Classdb struct {
	Classid  int32   `form:"classid" json:"classid" gorm:"primaryKey"`
	Teacher  int32   `form:"teacher" json:"teacher" gorm:"not null;unique"`
	Name     string  `binding:"required" form:"name" json:"name" gorm:"not null;unique"`
	Students []int32 `binding:"required" form:"students" json:"students" gorm:"-"`
}

func InitGorm() {
	dsn := SqlUserName+":"+SqlUserPwd+"@tcp("+SqlAddr+")/"+SqlName+"?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB,err=gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	err=DB.AutoMigrate(Userdb{},Classdb{})
	if err != nil {
		panic(err)
	}
}