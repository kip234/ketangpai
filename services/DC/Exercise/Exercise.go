package Exercise

import (
	"KeTangPai/services/Log"
	"context"
	"errors"
	"gorm.io/gorm"
	"io/ioutil"
	"os"
	"strconv"
)

type ExerciseService struct{
	db *gorm.DB//MySQL连接
}

//管理考试/作业信息

func newExerciseService()*ExerciseService{
	InitGorm()
	return &ExerciseService{db:sql}
}

//根据考试号获取考试详情-不含题目内容
func(e *ExerciseService)GetExercise(c context.Context,in *I) (*ExerciseData, error){
	Log.Send("Exercise.GetExercise.info",in)
	select {
	case <-c.Done():
		Log.Send("Exercise.GetExercise.error","GetExercise> timeout")
		return &ExerciseData{},errors.New("timeout")
	default:
	}
	//获取考试信息
	re:=Exercisedb{}
	err:=e.db.Model(Exercisedb{}).Where("id=?",in.I).Find(&re).Error
	if err!=nil{
		Log.Send("Exercise.GetExercise.error",err.Error())
		return &ExerciseData{},err
	}
	return &ExerciseData{
		Id:re.Id,
		Typ: re.Typ,
		Classid: re.Classid,
		Ownerid: re.Ownerid,
		Begin: re.Begin,
		End:re.End,
		Duration:re.Duration,
		Name: re.Name,
	},err
}
//根据考试号获取考试详情-含题目内容
func(e *ExerciseService)GetExercisec(c context.Context,in *I) (*ExerciseData, error){
	Log.Send("Exercise.GetExercisec.info",in)
	select {
	case <-c.Done():
		Log.Send("Exercise.GetExercisec.error","timeout")
		return &ExerciseData{},errors.New("timeout")
	default:
	}
	//获取题目列表储存路径
	var location string
	err:=e.db.Model(Exercisedb{}).Where("id=?",in.I).Select("location").Find(&location).Error
	if err!=nil{
		Log.Send("Exercise.GetExercisec.error",err.Error())
		return &ExerciseData{},err
	}
	//获取考试信息
	re:=Exercisedb{}
	err=e.db.Model(Exercisedb{}).Where("id=?",in.I).Find(&re).Error
	if err!=nil{
		Log.Send("Exercise.GetExercisec.error",err.Error())
		return &ExerciseData{},err
	}
	file,err:=os.Open(location)
	if err!=nil{
		Log.Send("Exercise.GetExercisec.error",err.Error())
		return &ExerciseData{},err
	}
	b,err:=ioutil.ReadAll(file)
	if err!=nil{
		Log.Send("Exercise.GetExercisec.error",err.Error())
		return &ExerciseData{},err
	}
	return &ExerciseData{
		Id:re.Id,
		Typ: re.Typ,
		Classid: re.Classid,
		Ownerid: re.Ownerid,
		Begin: re.Begin,
		End:re.End,
		Duration:re.Duration,
		Name: re.Name,
		Content: string(b),
	},err
}
//根据班级号获取考试列表
func(e *ExerciseService)GetExercises(in *I,stream Exercise_GetExercisesServer) error{
	Log.Send("Exercise.GetExercises.info",in)
	var A []uint32//考试ID列表
	err:=e.db.Model(Exercisedb{}).Where("classid=?",in.I).Select("id").Find(&A).Error
	if err!=nil {
		Log.Send("Exercise.GetExercises.error",err.Error())
		return err
	}
	for _,i:=range A{
		re,err:=e.GetExercise(context.Background(),&I{I:i})
		if err!=nil {
			Log.Send("Exercise.GetExercises.error",err.Error())
			return err
		}
		err=stream.Send(re)
		if err!=nil {
			Log.Send("Exercise.GetExercises.error",err.Error())
			return err
		}
	}
	return nil
}
//添加一次考试
func(e *ExerciseService)AddExercise(c context.Context,in *ExerciseData) (*ExerciseData, error){
	Log.Send("Exercise.AddExercise.info",in)
	select {
	case <-c.Done():
		Log.Send("Exercise.AddExercise.error","timeout")
		return &ExerciseData{},errors.New("timeout")
	default:
	}
	err:=e.db.Transaction(func(tx *gorm.DB)error{
		tmp:=Exercisedb{Classid: in.Classid,Ownerid: in.Ownerid,Typ: in.Typ,Begin: in.Begin,End: in.End,Duration: in.Duration,Name: in.Name}
		err:=tx.Model(Exercisedb{}).Create(&tmp).Error//存入基本信息获取ID
		if err!=nil {
			return err
		}
		in.Id=tmp.Id
		location:="./exercise/"+strconv.Itoa(int(in.Id))+".exc"
		err=tx.Model(Exercisedb{}).Where("id=?",in.Id).Update("location",location).Error//更新存储路径
		if err!=nil {
			return err
		}
		file,err:=os.Create(location)//创建本地文件
		if err!=nil {
			return err
		}
		defer file.Close()
		_,err=file.WriteString(in.Content)//存入内容
		if err!=nil {
			return err
		}
		return nil
	})
	if err!=nil {
		Log.Send("Exercise.AddExercise.error",err.Error())
	}
	return in,err
}
//学生提交一次考试记录
func(e *ExerciseService)SubmitAns(c context.Context,in *Submit) (*I, error){
	Log.Send("Exercise.SubmitAns.info",in)
	select {
	case <-c.Done():
		Log.Send("Exercise.SubmitAns.error","timeout")
		return &I{},errors.New("timeout")
	default:
	}

	var tmp string
	err:=e.db.Model(Submitdb{}).Where(Submitdb{Exerciseid:in.Exerciseid,Uploaderid: in.Uploaderid}).Select("location").Find(&tmp).Error
	if 	tmp!=""{//之前提交过了
		file,err:=os.Create(tmp)//覆盖原来的文件
		if err!=nil {
			Log.Send("Exercise.SubmitAns.error",err.Error())
			return &I{},err
		}
		defer file.Close()
		file.WriteString(in.Contents)
		return &I{},nil
	}
	//之前没有提交
	err=e.db.Transaction(func(tx *gorm.DB)error{
		tmp:=Submitdb{Uploaderid:in.Uploaderid, Exerciseid:in.Exerciseid}
		err=tx.Model(Submitdb{}).Save(&tmp).Error//存入数据
		if err!=nil {
			Log.Send("Exercise.SubmitAns.error",err.Error())
			return err
		}
		in.Id=tmp.Id
		var location="./submit/"+strconv.Itoa(int(in.Id))+".sub"
		err=tx.Model(Submitdb{}).Where("id=?",in.Id).Update("location",location).Error//更新存储路径
		if err!=nil {
			Log.Send("Exercise.SubmitAns.error",err.Error())
			return err
		}
		//开始写入内容
		file,err:=os.Create(location)
		if err!=nil {
			Log.Send("Exercise.SubmitAns.error",err.Error())
			return err
		}
		defer file.Close()
		_,err=file.WriteString(in.Contents)
		if err!=nil {
			Log.Send("Exercise.SubmitAns.error",err.Error())
			return err
		}
		return nil
	})
	if err!=nil {
		Log.Send("Exercise.SubmitAns.error",err.Error())
	}
	return &I{},err
}
//根据考试ID获取答案
func(e *ExerciseService)GetKey(c context.Context,in *I) (*Submit, error){
	Log.Send("Exercise.GetKey.info",in)
	select {
	case <-c.Done():
		Log.Send("Exercise.GetKey.error","timeout")
		return &Submit{},errors.New("timeout")
	default:
	}
	var own uint32
	err:=e.db.Model(Exercisedb{}).Where("id=?",in.I).Select("Ownerid").Find(&own).Error//找到考试的发起人
	if err!=nil {
		Log.Send("Exercise.GetKey.error",err.Error())
		return &Submit{},err
	}
	var location string//发起人提交内容的路径
	err=e.db.Model(Submitdb{}).Where(Submitdb{Uploaderid: own,Exerciseid: in.I}).Select("location").Find(&location).Error
	if err!=nil {
		Log.Send("Exercise.GetKey.error",err.Error())
		return &Submit{},err
	}
	if location == ""{//还没有上传
		return &Submit{},nil
	}
	file,err:=os.Open(location)
	if err!=nil {
		Log.Send("Exercise.GetKey.error",err.Error())
		return &Submit{},err
	}
	defer file.Close()
	b,err:=ioutil.ReadAll(file)
	if err!=nil {
		Log.Send("Exercise.GetKey.error",err.Error())
		return &Submit{},err
	}
	re:=Submit{}//用来返回的数据
	err=e.db.Model(Submitdb{}).Where("Uploaderid=?",own).Find(&re).Error//获取基础信息
	if err!=nil {
		Log.Send("Exercise.GetKey.error",err.Error())
		return &Submit{},err
	}
	re.Contents=string(b)//添加读取到的内容
	return &re,err
}
//给学生打分
func(e *ExerciseService)SetScore(c context.Context,in *Score) (*Empty1, error){
	Log.Send("Exercise.SetScore.info",in)
	select {
	case <-c.Done():
		Log.Send("Exercise.SetScore.error","timeout")
		return &Empty1{},errors.New("timeout")
	default:
	}
	//更新分值
	err:=e.db.Model(Submitdb{}).Where("id=?",in.Submitid).Update("value",in.Value).Error
	if err!=nil {
		Log.Send("Exercise.SetScore.error",err.Error())
	}
	return &Empty1{},err
}
//学生根据提交记录获取本次得分
func(e *ExerciseService)GetScore(c context.Context,in *I) (*Score, error){
	Log.Send("Exercise.GetScore.info",in)
	select {
	case <-c.Done():
		Log.Send("Exercise.GetScore.error","timeout")
		return &Score{},errors.New("timeout")
	default:
	}
	var v int32
	//寻找分值记录
	err:=e.db.Model(Submitdb{}).Where("id=?",in.I).Select("value").Find(&v).Error
	if err!=nil {
		Log.Send("Exercise.GetScore.error",err.Error())
	}
	return &Score{Value: v},err
}
//学生根据自己的ID获取自己的所有提交记录
func(e *ExerciseService)GetScores(in *I,stream Exercise_GetScoresServer) error{
	Log.Send("Exercise.GetScores.info",in)
	var re []Submitdb
	//从数据库中查找记录
	err:=e.db.Model(Submitdb{}).Where("Uploaderid=?",in.I).Find(&re).Error
	if err!=nil {
		Log.Send("Exercise.GetScores.error","timeout")
		return err
	}
	for _,i:=range re{
		file,err:=os.Open(i.Location)
		if err!=nil {
			Log.Send("Exercise.GetScores.error",err.Error())
			return err
		}
		b,err:=ioutil.ReadAll(file)
		if err!=nil {
			Log.Send("Exercise.GetScores.error",err.Error())
			return err
		}
		err=stream.Send(&Submit{
			Id:i.Id,
			Uploaderid:i.Uploaderid,
			Exerciseid:i.Exerciseid,
			Contents:string(b),
			Value:i.Value,
		})
		if err!=nil {
			Log.Send("Exercise.GetScores.error",err.Error())
			return err
		}
	}
	return nil
}
//老师根据考试ID获取本次班级得分情况
func(e *ExerciseService)GetClassScores(in *I,stream Exercise_GetClassScoresServer) error{
	Log.Send("Exercise.GetClassScores.info",in)
	var ids []uint32
	//获取该考试的所有提交记录
	err:=e.db.Model(Submitdb{}).Where("Exerciseid=?",in.I).Select("id").Find(&ids).Error
	if err!=nil {
		Log.Send("Exercise.GetClassScores.error",err.Error())
		return err
	}
	var v int32
	for _,i:=range ids{
		//获取分值
		err:=e.db.Model(Submitdb{}).Where("id=?",i).Select("value").Find(&v).Error
		if err!=nil {
			Log.Send("Exercise.GetClassScores.error",err.Error())
			return err
		}
		err=stream.Send(&Score{Submitid:i,Value: v})
		if err!=nil {
			Log.Send("Exercise.GetClassScores.error",err.Error())
			return err
		}
	}
	return nil
}
//老师根据考试ID获取本次提交情况
func(e *ExerciseService)GetClassSubmit(in *I,stream Exercise_GetClassSubmitServer) error{
	Log.Send("Exercise.GetClassSubmit.info",in)
	var s []Submitdb
	//获取所有提交记录
	err:=e.db.Model(Submitdb{}).Where("Exerciseid=?",in.I).Find(&s).Error
	if err!=nil {
		Log.Send("Exercise.GetClassSubmit.error",err.Error())
		return err
	}
	//提取提交内容并返回
	for _,i:=range s{
		file,err:=os.Open(i.Location)
		if err!=nil {
			Log.Send("Exercise.GetClassSubmit.error",err.Error())
			return err
		}
		b,err:=ioutil.ReadAll(file)
		if err!=nil {
			Log.Send("Exercise.GetClassSubmit.error",err.Error())
			return err
		}
		file.Close()
		stream.Send(&Submit{
			Id: i.Id,
			Uploaderid: i.Uploaderid,
			Exerciseid: i.Exerciseid,
			Contents: string(b),
			Value:i.Value,
		})
	}
	return nil
}
//根据考试ID删除考试记录-试题，提交记录等
func(e *ExerciseService)DelExercise(c context.Context,in *I) (*Empty1, error){
	Log.Send("Exercise.DelExercise.info",in)
	select {
	case <-c.Done():
		Log.Send("Exercise.DelExercise.error","timeout")
		return &Empty1{},errors.New("timeout")
	default:
	}
	tmp:=Exercisedb{}
	err:=e.db.Transaction(func(tx *gorm.DB)error{
		err:=tx.Model(Exercisedb{}).Where("id=?",in.I).Find(&tmp).Error
		if err!=nil{
			return err
		}
		//删除记录
		err=tx.Model(Exercisedb{}).Delete(tmp).Error
		if err!=nil{
			return err
		}
		//删除题目
		if tmp.Location!=""{
			err=os.Remove(tmp.Location)
			if err!=nil{
				return err
			}
		}
		//删除提交记录
		var location []string//提交的文件
		err=tx.Model(Submitdb{}).Where("Exerciseid=?",tmp.Id).Select("location").Find(&location).Error
		if err!=nil{
			return err
		}
		//删除提交内容
		for _,i:=range location{
			err=os.Remove(i)
			if err!=nil{
				return err
			}
		}
		return nil
	})
	if err!=nil {
		Log.Send("Exercise.DelExercise.error",err.Error())
	}
	return &Empty1{},err
}
//根据班级ID删除该班级所有记录
func(e *ExerciseService)DelExercises(c context.Context,in *I) (*Empty1, error){
	Log.Send("Exercise.DelExercises.info",in)
	select {
	case <-c.Done():
		Log.Send("Exercise.DelExercises.error","timeout")
		return &Empty1{},errors.New("timeout")
	default:
	}
	err:=e.db.Transaction(func(tx *gorm.DB)error{
		var A []uint32//考试ID列表
		//获取所有考试列表
		err:=e.db.Model(Exercisedb{}).Where("classid=?",in.I).Select("id").Find(&A).Error
		if err!=nil {
			return err
		}
		for _,i:=range A{
			_,err=e.DelExercise(context.Background(),&I{I: i})
			if err!=nil {
				return err
			}
		}
		return nil
	})
	if err!=nil {
		Log.Send("Exercise.DelExercises.error",err.Error())
	}
	return &Empty1{},err
}

func(e *ExerciseService)mustEmbedUnimplementedExerciseServer(){}