package UserCenter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

const Addr="localhost:8087"

//sql
const (
	SqlName="ketangpai"
	SqlUserName="root"
	SqlUserPwd="root"
	SqlAddr=":3306"
)

//储存对象
type Userdb struct {
	Id     uint32   `form:"uid" json:"uid" gorm:"primaryKey"`
	Pwd     string	`binding:"required" form:"pwd" json:"pwd" gorm:"not null"`
	Email string    `binding:"required" form:"email" json:"email" gorm:"not null;unique"`
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
	err=DB.AutoMigrate(Userdb{})
	if err != nil {
		panic(err)
	}
}