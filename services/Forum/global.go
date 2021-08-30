package Forum

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Messagedb struct{
	Id      int32  `form:"id" json:"id" gorm:"primaryKey"`  //自身ID
	Owner   int32  `form:"owner" json:"owner" gorm:"not null"`    //发起者
	Tosb    int32  `form:"tosb" json:"tosb" gorm:"not null"`    //对某人说的
	Content string `binding:"required" form:"content" json:"content" gorm:"not null"` 	//说了啥？
	Time    int64  `form:"time" json:"time" gorm:"not null"`    //说的时间
	Classid int32  `form:"classid" json:"classid" gorm:"not null"`	//所属班级
}

var DB *gorm.DB

const Addr="localhost:8089"

//sql
const (
	SqlName="ketangpai"
	SqlUserName="root"
	SqlUserPwd="root"
	SqlAddr="127.0.0.1:3306"
)


func InitGorm() {
	dsn := SqlUserName+":"+SqlUserPwd+"@tcp("+SqlAddr+")/"+SqlName+"?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB,err=gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	err=DB.AutoMigrate(Messagedb{})
	if err != nil {
		panic(err)
	}
}