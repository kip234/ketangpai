package Exercise

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const Addr		= ":8082"

var sql *gorm.DB

//type
const (
	DefaultTyp = iota
	Unlimited 	//不限时
	LimitedDate			//有日期限制
	TimeLimit			//没有日期限制
	TypNum
)

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

//sql
const (
	SqlName="ketangpai"
	SqlUserName="root"
	SqlUserPwd="root"
	SqlAddr="127.0.0.1:3306"
)

type Exercisedb struct{
	Id       int32  `form:"classid" json:"id" gorm:"primaryKey"` //数据自身ID
	Classid  int32  `form:"classid" json:"classid" gorm:"not null"`	//所属班级
	Ownerid  int32  `form:"ownerid" json:"ownerid" gorm:"not null"`	//发布人
	Content  string `binding:"required" form:"content" json:"content" gorm:"-"`			//题目列表
	Typ      uint32  `binding:"required" form:"typ" json:"typ"gorm:"not null"`	//类型起始日期与截止日期、无时间限制、单次限时
	Begin    int64	`form:"begin" json:"begin"`					//起始日期
	End      int64	`form:"end" json:"end"`					//截止日期
	Duration int64	`form:"duration" json:"duration"`					//持续时长
	Name     string `binding:"required" form:"name" json:"name" gorm:"not null"`	//考试名
	Location string `gorm:"not null" form:"location" json:"location"`   //储存位置
	Discipline uint32 `form:"discipline" json:"discipline" gorm:"not null"`//学科
}

type Submitdb struct{
	Id         int32  `form:"id" json:"id" gorm:"primaryKey"`	//数据本身的ID-固定！！！
	Uploaderid int32  `binding:"required" form:"uploaderid" json:"Uploaderid" gorm:"not null"`		//上传者ID
	Exerciseid int32  `binding:"required" form:"exerciseid" json:"exerciseid" gorm:"not null"`		//考试ID
	Contents   string `binding:"required" form:"content" json:"contents" gorm:"-"`			//提交内容
	Value      int32  `form:"value" json:"value"`						//得分
	Location string	  `form:"location" json:"location" gorm:"not null"`		//储存位置
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
	err=sql.AutoMigrate(Exercisedb{},Submitdb{})
	if err != nil {
		panic(err)
	}
}