package Models

type In struct{//从前端获取信息
	Content string//内容
	To []int32//对某些人说
}

type Out struct{//给前端的消息
	In
	Classid int32//班级ID
	Classname string//班级名字
	Uid int32//用户ID
	Uname string//用户名
	Online int32//在线人数
	Ranks []string//排名信息
}
