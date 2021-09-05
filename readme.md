# ？

> 其中邮件以及mq相关设置因为涉及鄙人个人信息所以上传的版本不可用，
>
> 虽然用到了grpc但为了写代码方便用协程来代替机群，耦合度较低可以直接将server.go的协程服务拆成单独的程序。(除了logs部分…)
>
> 云地址：121.4.76.240:8080 也许在将来某一天会不可用

目录：

> 如果没有把服务单独拆出来：
>
> - exercise
> - files
> - logs
>     - info
>     - error
> - submit
> - test

## 功能

> 账号注册登录注销？
>
> 创建课堂,布置作业
>
> 上传-下载课件资料
>
> 发布话题-话题讨论
>
> 上课签到
>
> 课中提问(学生or老师)和回答（抢答or抽答）
>
> 成绩管理

特殊说明

> 参数有可选必选之分，但由于历史原因，有的时候必选参数不给也能得到结果，只是结果不尽人意，所以该类型参数可以认为‘建议必选’

## 用到的第三方库

```http
github.com/gin-gonic/gin
github.com/gomodule/redigo/redis
github.com/streadway/amqp
gorm.io/gorm
```



## 类型说明

```go
package Exercise
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


package TestBank
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

package UserCenter
//用户类型
const (
	DefaultType = iota
	Teacher
	Student
	Administrator
	TypeNum
)
```



## 路由结构

```go
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
```



## 操作说明

### 1.注册

> 拥有一个账号是一切的开始
>
> 如果成功将会返回UID等重要信息

```http
GET ws://host/register
```

> 建立websocket连接后用户需要以JSON格式传回以下数据

| KEY   | 类型 | 描述   |
| ----- | ---- | ------ |
| name  | 必选 | 用户名 |
| pwd   | 必选 | 密码   |
| email | 必选 | 邮箱   |

```json
{
    "name":"",
    "email":"",
    "pwd":""
}
```



> 如果一切顺利的话会收到含有验证码的邮件，将验证码以纯文本的形式传回即可，最后在返回相关信息后服务器会主动断开连接

### 2.登录

> 拿到上一步的UID和颅内密码后就可以进行下一步登陆了
>
> 如果成功将会返回JWT

```http
POST host/login
```

fohrm-data

| KEY  | 类型 | 描述 |
| ---- | ---- | ---- |
| uid  | 必选 | UID  |
| pwd  | 必选 | 密码 |

> **到这里就可以使用了，后续操作都需要登陆后的JWT**
>
> 将JWT放在header里面，key为token

### 3.注销

> 指退出登录
>
> 自动退出

```http
POST host/logout
```

### 4.修改信息

> 想改名怎么办？想加入班级怎么办？想改密码怎么办？想换个角色(老师/学生)又该怎么办？
>
> 快来setting路由下试试吧…

```http
POST host/setting
```

```http
GET ws://host/setting
```

> 这里提供了两种方法，POST返回当前的账户信息，GET会修改相关信息 -为什么GET修改？因为websocket是GET

> POST不需要参数，以下为GET需要的内容

form-data

| KEY     | 类型 | 描述                     |
| ------- | ---- | ------------------------ |
| name    | 可选 | 名字                     |
| pwd     | 可选 | 密码                     |
| uid     | 必选 | 由前端返回，不对用户开放 |
| classid | 可选 | 班级号                   |
| type    | 可选 | 目前对用户开放           |

```json
{
    "name":"",
    "pwd":"",
    "uid":,
    "classid":,
    "type":
}
```



### 5.题库操作

> 只对管理员开放，不提供注册管理员的方法

```http
POST host:/testbank_upload
```

form-data

| KEY        | 类型 | 描述                     |
| ---------- | ---- | ------------------------ |
| typ        | 必选 | 题目类型                 |
| content    | 必选 | 题目主体                 |
| ans        | 可选 | 答案                     |
| name       | 必选 | 可以是方便记忆的任何内容 |
| discipline | 必选 | 科目                     |

### 6.是时候创建自己的班级了

> 前提是你的用户类型为老师

```http
POST host/create_class
```

body

```json
{
    "name":"",
    "students":[]
}
```

> 班级名以及初始学生列表，切片里面放用户ID

### 7.上传群文件

```http
POST host/class/file/upload
```

form-data

| KEY  | 类型 | 描述       |
| ---- | ---- | ---------- |
| file | 必选 | 上传的文件 |

### 8.下载文件

```http
GET host/class/file/download?fileid
```

> fileid: 文件ID

### 9.查看文件列表

```http
GET host/class/file/contents
```

> 不需要特殊参数

### 10.查看群聊

```http
GET host/class/forum
```

> 获取记录

```http
POST host/class/forum
```

> 发言

form-data

| KEY     | 类型 | 描述         |
| ------- | ---- | ------------ |
| content | 必选 | 内容         |
| tosb    | 可选 | 指定发给某人 |

> Tosb是to somebody的简写，即使明确该字段所有人也能看见，只是somebody可以通过下面的方式快速查看

### 11.查看留言

```http
GET host/class/forum/messages
```

> 不需要参数

> **老师的特权**

### 12.从自己的班级里面除名某些人

```http
POST host/class/fire
```

body

```json
[]
```

> 直接传一个数组，里面放出名的UID

### 13.解散班级

```http
POST host/class/dissolve
```

> 不需要别的参数，自动解散

### 14.是时候给同学们布置任务了

```http
POST host/class/assign_homework?auto&subjective&objective
```

> 有两种模式，以auto的值做区分

> 如果auto为1则内容自动生成，此时subjective、objective的取值代表主观题与客观题的数量。

>  其他取值认为用户自己上传内容

> **上面两种只区分内容，其他信息由form决定**

| KEY        | 类型 | 描述                   |
| ---------- | ---- | ---------------------- |
| typ        | 必选 | 类型(限时/限日期/不限) |
| begin      | 可选 | 开始日期               |
| end        | 可选 | 结束日期               |
| duration   | 可选 | 持续时间               |
| name       | 必选 | 任务名                 |
| content    | 可选 | 内容                   |
| discipline | 必选 | 科目                   |

> 和时间有关的统一使用Unix时间戳

### 15.让我看看他们做得怎么样了

```http
GET host/class/check_test_status?id
```

> (id指任务id)如果id已指定就会返回对应任务的提交情况，否则列出所有历史任务

### 16.打分时间

```http
POST host/class/mark?id&v
```

> id:提交记录的ID	v:分值

> 来看看学生这边**后面部分为学生专享**

### 17.查看班级任务

```http
POST host/class/assignment
```

> 不需要别的参数，自动返回所属班级的任务记录

### 18.查看分值

```http
GET host/class/grade?id
```

> id:提交记录ID 如果不指定的话返回所有提交记录

### 19.请诚信考试…

```http
GET ws://host/class/examination_room?eid
```

> eid:任务ID 进行指定的任务，由于后台阻塞接收，所以要等到前边提交后才能判断时间

### 20.课堂互动(实时聊天)

```http
GET ws://localhost:8080/class/chatroom
```

> 前提是需要登录且已经加入班级

```json
{
    "Content":"",
    "To":[]
}
```

> Content:内容	To:可以看见这条消息的人(默认包含自己)

### 21.找回账号信息

```http
POST host/retrieve?email
```

> email为注册时预留的邮箱，稍后会以邮件的方式发送

## services



> 由于一些未知原因，鄙人目前还没有搞清楚import的用法，为了避免命名冲突，出现了不同文件中定义了命名看似不同字段内容却又雷同的情况。
>
> 服务与服务之间不能直接访问

### Email

> 邮件服务

```protobuf
syntax="proto3";
option go_package="./Email";
package Email;

service Email{
  rpc send(mail)returns(empty){};//发邮件
}

message mail{
  string subject=1;//主题
  string content=2;//内容
  string to=3;//收件人
}

message empty{}
```



### Exercise

> 处理布置的任务

```protobuf
syntax="proto3";
option go_package="./exercise";
package exercise;

service exercise{
  rpc get_exercise(i)returns(exerciseData){};//根据考试号获取考试详情-不含题目内容
  rpc get_exercisec(i)returns(exerciseData){};//根据考试号获取考试详情-含题目内容
  rpc get_exercises(i)returns(stream exerciseData){};//根据班级号获取考试列表
  rpc add_exercise(exerciseData)returns(exerciseData){};//添加一次考试
  rpc submit_ans(submit)returns(i){};//提交一次考试记录-这里不做时间检测
  rpc get_key(i)returns(submit){};//根据考试ID获取答案
  rpc set_score(score)returns(empty1){};//给学生打分
  rpc get_score(i)returns(score){};//学生根据提交记录获取本次得分
  rpc get_scores(i)returns(stream submit){};//学生根据自己的ID获取所有提交记录
  rpc get_class_scores(i)returns(stream score){};//老师根据考试ID获取本次班级得分情况
  rpc get_class_submit(i)returns(stream submit){};//老师根据考试ID获取本次提交情况
  rpc del_exercise(i)returns(empty1){};//根据考试ID删除考试记录-试题，提交记录等
  rpc del_exercises(i)returns(empty1){};//根据班级ID删除该班级所有记录
}

message empty1{}

message i{
  int32 i=1;
}

message exerciseData{
  int32 id=1;//数据自身ID
  int32 classid=2;//所属班级
  int32 ownerid=3;//发布人
  string content=4;//内容
  int32 typ=5;//类型起始日期与截止日期、无时间限制、单次限时
  int64 begin=6;//起始日期
  int64 end=7;//截止日期
  int64 duration=8;//持续时长
  string name=9;//考试名
}

message submit{//只记录最近的一次-在时限内可以多次提交
  int32 id=1;//数据本身的ID-固定！！！
  int32 uploaderid=2;//上传者ID
  int32 exerciseid=3;//考试ID
  string contents=4;//提交内容
  int32 value=5;//得分
}

message score{
  int32 submitid=1;//提交记录
  int32 judge=2;//打分人
  int32 value=3;//分值
}
```

### Filter

> 敏感词句过滤

```protobuf
syntax="proto3";
option go_package="./Filter";
package Filter;

service Filter{
  rpc Process(FilterData)returns(FilterData){};
  rpc Add(FilterData)returns(FilterData){};

}

message right{
  bool right=1;
}

message FilterData{
  bytes data=1;
}
```

### Forum

> 留言等…

```protobuf
//留言什么的
syntax="proto3";
option go_package="./Forum";
package Forum;

service Forum{
  rpc speak(message)returns(message){};//发言
  rpc get_message(uid)returns(messages){};//查看别人给自己的留言
  rpc get_history(classid)returns(messages){};//查看班级历史
}

message classid{
  int32 id=1;
}

message uid{
  int32 id=1;
}

message message{
  int32 id=1;//自身ID
  int32 owner=2;//发起者
  int32 tosb=3;//对某人说的
  string content=4;//说了啥？
  int64 time=5;//说的时间
  int32 classid=6;//所属班级
}

message messages{//一段话
  repeated message m=1;
}
```

### JWT

> 生成和检测JWT

```protobuf
syntax="proto3";
option go_package="./JWT";
package JWT;

service JWT{
  rpc refresh_token(juser) returns(token){};//刷新token
  rpc check_token(token)returns(juser){};//检查token
  rpc del_jwt(juser)returns(token){};//删除token
}

message token{
  string content=1;
}

message juser{
  string name=1;
  int32 uid=2;
  string pwd=3;
}
```

> 主要任务是token操作，用Redis进行token的缓存

### NetworkDisk

> 班级文件管理

```protobuf
syntax="proto3";
option go_package="./NetworkDisk";
package NetworkDisk;

service NetworkDisk{
  rpc download(fileid)returns(stream filestream);//下载文件
  rpc upload(stream filestream)returns(fileinfo);//上传文件
  rpc get_contents(classid)returns(contents);//获取文件目录
}

message fileid{
  int32 id=1;
}

message classid{
  int32 id=1;
}

message contents{
  repeated string name=1;
  repeated int32 id=2;
}

message filestream{
  bytes content=1;
}

message fileinfo{
  int32 id=1;
  int32 uploader=2;
  int32 classid=3;
  string name=4;
  int64 size=5;
  int64 time=6;
}
```

### RankingList

> 排行榜

```protobuf
syntax="proto3";
option go_package="./RankingList";
package RankingList;

//成员与排名：一对一
//排名与成员：多对多

service RankingList{
  rpc flushlist(flushin)returns(flushout){};//刷新榜单-用传入的数据刷新 榜单、成员名、刷新量
  rpc getlistinfo(listname)returns(listinfo){};//获取排行榜完整信息
  rpc dellist(listname)returns(empty){};//删除榜单
  rpc getranking(members)returns(rankings){};//获取排名
}

message empty{}

message listname{
  string name=1;
}

message members{
  string name=1;//榜单名
  repeated string members=2;
}

message rankings{
  repeated int64 rank=1;
}

message flushin{
  string key=1;//刷新的表
  int32 increment=2;//刷新量
  string member=3;//刷新对象

}

message flushout{
  string member=1;//刷新对象
  int64 ranking=2;//刷新后排名
}

message listinfo{
  string name=1;//榜名
  repeated string list=2;//榜单-仅对象/对象+分值、、、
}

```



### TestBank

> 试题库

```protobuf
syntax="proto3";
option go_package="./TestBank";
package TestBank;

service TestBank{
  rpc upload(test)returns(test){};//上传一道题
  rpc download(stream testid)returns(stream test){};//下载题目
  rpc generate_test(testconf)returns(stream testid){};//自动生成一套试卷
}

message ans{//答案
  string ans=1;
}

message testid{//试题ID
  int32 id=1;
}

message testconf{
  int32 subjective_item=1;//主观题的数量
  int32 objective_item=2;//客观题的数量
  int32 discipline=3;//学科
}

message test{
  int32 id=1;//自身ID
  int32 typ=2;//类型-主观题/客观题
  string content=3;//内容
  string ans=4;//答案(如果有)
  string name=5;//名字-题目描述
  int32 uploader=6;//上传者
  int32 discipline=7;//学科
}
```

### UserCenter

> 储存用户信息

```protobuf
syntax="proto3";
option go_package="./UserCenter";
package UserCenter;

service UserCenter{
  rpc creat_user(uuser) returns(uuser){};//创建用户
  rpc get_user_info(id) returns(uuser){};//获取用户所有信息
  rpc get_user_info_by_email(s)returns(uuser){};//参照邮箱返回用户信息
  rpc get_user_pwd(id)returns(s){};//获取密码
  rpc get_user_name(id)returns(s){};//获取名字
  rpc get_user_email(id)returns(s){};//获取联系方式
  rpc get_user_class(id)returns(id){};//获取班级
  rpc get_user_type(id)returns(id){};//获取用户类型
  rpc user_is_Exist(s)returns(right){};//判断用户是否存在:以邮箱为标准
  rpc refreshing_user_data(uuser)returns(uuser){};//刷新用户数据

  rpc create_class(class)returns(class){};//创建一个班级
  rpc get_class_info(id)returns(class){};//获取班级所有信息
  rpc get_class_teacher(id)returns(id){};//获取班级负责人ID
  rpc get_class_name(id)returns(s){};//获取班级名字
  rpc dissolve_class(id)returns(empty){};//解散班级
  rpc refreshing_class_data(class)returns(class){};//更新班级数据-名字、老师
  rpc fire_student(id)returns(class){};//将某学生从班级里面删除
}

message empty{
}

message id{
  int32 i=1;
}

message s{
  string s=1;
}

message right{
  bool right=1;
}

message uuser{
  int32 uid=2;
  string name=1;
  string pwd=3;
  int32 type=4;
  int32 classid=5;
  string email=6;
}

message class{
  int32 classid=1;
  int32 teacher=2;
  string name=3;
  repeated int32 students=4;
}

```

> 使用MySQL储存，用uusers和class两张表

### TestBank

```protobuf
syntax="proto3";
option go_package="TestBank";
service TestBank{
  rpc download_questions(stream questionid)returns(stream question){};//下载试题-可以多个
  rpc upload_questions(stream question)returns(stream question){};//上传试题
}
message questionid{
  int64 id=1;
}
message question{
  int64 id=2;//试题ID
  string content=3;//试题内容
}
```

> 储存和发放试题

### Log

> 收集日志，使用rabbitmq

> 关于该服务：这部分写得并不好，因为error等级没有明确的划分，不管是因为用户的数据有问题还是程序执行过程中自己产生的问题都一视同仁，这样并不利于后面的维护以及及时报警(打算用邮件通知)，问什么我没有继续优化？那当然是因为要抓住假期仅剩的余额啦。

## 新知识？

> 使用gorm操作时，只要用Model()指明查找的表，就可以用另一种结构接收数据，要求是该结构的所有开放字段被包含于Model()所指明类型的字段中，于是用于传输数据的结构可以与面向储存的结构分离