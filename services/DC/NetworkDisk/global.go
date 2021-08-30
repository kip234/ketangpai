package NetworkDisk

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const Addr		= ":8084"
const TransmissionUnit=1024//bytes

var sql *gorm.DB

//sql
const (
	SqlName="ketangpai"
	SqlUserName="root"
	SqlUserPwd="root"
	SqlAddr="127.0.0.1:3306"
)

type fileinfodb struct{
	Id       int32  `gorm:"primaryKey"`
	Uploader int32  `gorm:"not null"`
	Classid  int32  `gorm:"not null"`
	Name     string `gorm:"not null"`
	Size     int64  `gorm:"not null"`
	Time     int64  `gorm:"not null"`
	Location string `gorm:"not null"`
}

func InitGorm() {
	dsn := SqlUserName+":"+SqlUserPwd+"@tcp("+SqlAddr+")/"+SqlName+"?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	sql,err=gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	err=sql.AutoMigrate(fileinfodb{})
	if err != nil {
		panic(err)
	}
}