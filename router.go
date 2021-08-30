//建立路由结构

package main

import (
	"KeTangPai/Handlers"
	"KeTangPai/Middlewares"
	"KeTangPai/Models"
	"github.com/gin-gonic/gin"
)

func BuildRouter(s Services,rooms map[int32]*Models.Room) *gin.Engine {
	server:=gin.Default()

	server.GET("/register", Handlers.Register(s.User,s.Email))
	server.POST("/login", Middlewares.CheakUserInfo(s.User),Handlers.Login(s.JWT))
	server.POST("/retrieve",Handlers.Retrieve(s.User,s.Email))//找回密码

	user:=server.Group("/",Middlewares.CheakJWT(s.JWT,s.User))//已登录
	{
		user.POST("/logout", Handlers.Logout(s.JWT))//注销
		user.POST("/setting", Handlers.Setting(s.User,s.Email))  //获取信息
		user.GET("/setting", Handlers.Setting(s.User,s.Email)) //修改信息 - websocket

		admin:=user.Group("/",Middlewares.IsAdmin(s.User))//管理员？
		{
			admin.POST("/testbank_upload",Handlers.Testbank_upload(s.TestBank))//上传
		}

		classmember:=user.Group("/class", Middlewares.HaveClass(s.User))//班级成员
		{
			classmember.POST("/file/upload", Handlers.File_upload(s.NetworkDisk))    //上传文件
			classmember.GET("/file/download", Handlers.File_download(s.NetworkDisk)) //下载文件
			classmember.GET("/file/contents", Handlers.File_contents(s.NetworkDisk)) //查看目录
			classmember.GET("/forum", Handlers.History(s.Forum))                     //查看记录
			classmember.POST("/forum",Handlers.Speak(s.Forum,s.Filter))              //发言
			classmember.GET("/forum/messages",Handlers.Messages(s.Forum))            //查看留言
			classmember.GET("/chatroom",Handlers.ChatRoom(s.Filter,rooms,s.RankingList))//进入教室
		}

		teacher := user.Group("/", Middlewares.IsTeacher(s.User))//老师
		{
			teacher.POST("/create_class", Handlers.Create_class(s.User)) //创建班级
			monitor:=teacher.Group("/class", Middlewares.HaveClass(s.User))//班级负责人
			{
				monitor.POST("/assign_homework", Handlers.Assign_homework(s.Exercise,s.TestBank))//布置作业
				monitor.GET("/check_test_status", Handlers.Check_test_status(s.Exercise))//查看考试情况
				monitor.POST("/dissolve", Handlers.Dissolve(s.User,s.Exercise))         //解散班级
				monitor.POST("/fire", Handlers.Fire(s.User))                 			//开除某人
				monitor.POST("/mark",Handlers.Mark(s.Exercise))//打分
			}
		}

		student := user.Group("/", Middlewares.IsStudent(s.User))//普通学生
		{
			classmate := student.Group("/class", Middlewares.HaveClass(s.User)) //同学
			{
				classmate.GET("/assignment", Handlers.Assignment(s.Exercise)) //查看任务-限时考试、作业
				classmate.GET("/grade", Handlers.Grade(s.Exercise))                       //成绩分析
				classmate.GET("/examination_room",Handlers.Examination_room(s.Exercise,s.Filter))//开始做题
			}
		}
	}
	return server
}