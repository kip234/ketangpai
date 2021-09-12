//建立路由结构

package main

import (
	"KeTangPai/Handlers"
	"KeTangPai/Middlewares"
	"KeTangPai/Models"
	"github.com/gin-gonic/gin"
)

func BuildRouter(s Services,rooms map[uint32]*Models.Room) *gin.Engine {
	server:=gin.Default()

	server.GET("/register", Handlers.Register(s.User,s.Email))
	server.POST("/login", Middlewares.CheakUserInfo(s.User),Handlers.Login(s.JWT))
	server.POST("/retrieve",Handlers.Retrieve(s.User,s.Email))//找回密码

	user:=server.Group("/user",Middlewares.CheakJWT(s.JWT,s.User,s.KetangpaiDB,s.RBAC))//已登录
	{
		user.POST("/logout", Handlers.Logout(s.JWT))//注销
		user.GET("/setting", Handlers.Setting())  //获取信息
		user.GET("/chpwd", Handlers.Chpwd(s.User,s.Email)) //修改密码 password - websocket
		user.POST("/chname",Handlers.Chname(s.KetangpaiDB))//改名
		user.POST("/create_class", Handlers.Create_class(s.KetangpaiDB,s.RBAC)) //创建班级

		admin:=user.Group("/admin",Middlewares.CheakRole(s.RBAC,"/admin"))//管理员？
		{
			admin.POST("/testbank_upload",Handlers.Testbank_upload(s.TestBank))//上传

			rbac:=admin.Group("/rbac")
			{
				rbac.POST("/add/role",Handlers.Add_role(s.RBAC))//添加角色
				rbac.POST("/add/user_role",Handlers.Add_user_role(s.RBAC))//添加 用户-角色
				rbac.POST("/add/routes",Handlers.Add_routes(s.RBAC))//添加 路由
				rbac.POST("/cache",Handlers.Cache(s.RBAC))//刷新缓存
			}
		}

		classmember:=user.Group("/class", Middlewares.CheakRole(s.RBAC,"/class"))//班级成员
		{
			file:=classmember.Group("/file")//班级文件操作
			{
				file.POST("upload", Handlers.File_upload(s.NetworkDisk))    //上传文件
				file.GET("download", Handlers.File_download(s.NetworkDisk)) //下载文件
				file.GET("contents", Handlers.File_contents(s.NetworkDisk)) //查看目录
			}

			forum:=classmember.Group("/forum")//讨论区操作
			{
				forum.GET("/history", Handlers.History(s.Forum))                     //查看记录
				forum.POST("/speak",Handlers.Speak(s.Forum,s.Filter))              //发言
				forum.GET("/messages",Handlers.Messages(s.Forum))            //查看留言
			}

			classmember.GET("/chatroom",Handlers.ChatRoom(s.Filter,rooms,s.RankingList))//进入教室
			classmember.GET("/assignment", Handlers.Assignment(s.Exercise)) //查看任务-限时考试、作业

			teacher:=classmember.Group("/teacher", Middlewares.CheakRole(s.RBAC,"/teacher"))//班级负责人
			{
				teacher.POST("/fire", Handlers.Fire(s.KetangpaiDB))                 			//开除某人
				teacher.POST("/dissolve", Handlers.Dissolve(s.KetangpaiDB,s.Exercise))         //解散班级
				teacher.POST("/assign_homework", Handlers.Assign_homework(s.Exercise,s.TestBank))//布置作业
				teacher.GET("/check_test_status", Handlers.Check_test_status(s.Exercise))//查看考试情况
				teacher.POST("/mark",Handlers.Mark(s.Exercise))//打分
				teacher.POST("/add",Handlers.Add(s.KetangpaiDB))//把某些人添加进班级
			}

			student := classmember.Group("/student", Middlewares.CheakRole(s.RBAC,"/student")) //同学
			{
				student.GET("/grade", Handlers.Grade(s.Exercise))                       //成绩分析
				student.GET("/examination_room",Handlers.Examination_room(s.Exercise,s.TestBank))//开始做题
			}
		}
	}
	return server
}