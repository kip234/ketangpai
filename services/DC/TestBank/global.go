package TestBank

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const Addr		= ":8085"

var sql *gorm.DB

//科目
const (
	DefaultDiscipline =iota
	Mathematics
	English
	Physics
	CLang
	Python
	Java
	Sports
	DisciplineNum
)

//题目类型
const (
	DefaultTyp=iota
	Subjective//主观题
	Objective //客观题
	TypNum
)

//sql
const (
	SqlName="ketangpai"
	SqlUserName="root"
	SqlUserPwd="root"
	SqlAddr="127.0.0.1:3306"
)

type Testdb struct {
	Id      int32  `form:"id" json:"id" gorm:"primaryKey"`          //自身ID
	Typ     uint32  `binding:"required" form:"typ" json:"typ" gorm:"not null"`        //类型-主观题/客观题
	Content string `binding:"required" form:"content" json:"content" gorm:"-"` //内容
	Ans     string `form:"ans" json:"ans" gorm:"-"`         //答案(如果有)
	Name    string `binding:"required" form:"name" json:"name" gorm:"not null"`       //名字-题目描述
	Location string`form:"location" json:"location" gorm:"not null"`		//储存路径
	Uploader int32 `form:"uploader" json:"uploader" gorm:"not null"` //上传者
	Discipline uint32 `form:"discipline" json:"discipline" gorm:"not null"`//学科
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
	err=sql.AutoMigrate(Testdb{})
	if err != nil {
		panic(err)
	}
}