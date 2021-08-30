package Exercise

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type ExerciseService struct{
	db *gorm.DB//MySQL连接
}

func newExerciseService()*ExerciseService{
	InitGorm()
	return &ExerciseService{db:sql}
}

//根据考试号获取考试详情-不含题目内容
func(e *ExerciseService)GetExercise(c context.Context,in *I) (*ExerciseData, error){
	log.Printf("GetExercise: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetExercise> timeout\n")
		return &ExerciseData{},errors.New("timeout")
	default:
	}
	//获取考试信息
	re:=Exercisedb{}
	err:=e.db.Model(Exercisedb{}).Where("id=?",in.I).Find(&re).Error
	if err!=nil{
		log.Printf("GetExercise> %s\n",err.Error())
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
	log.Printf("GetExercisec: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetExercisec> timeout\n")
		return &ExerciseData{},errors.New("timeout")
	default:
	}
	////获取题目列表储存路径
	var location string
	err:=e.db.Model(Exercisedb{}).Where("id=?",in.I).Select("location").Find(&location).Error
	if err!=nil{
		log.Printf("GetExercisec> %s\n",err.Error())
		return &ExerciseData{},err
	}
	//获取考试信息
	re:=Exercisedb{}
	err=e.db.Model(Exercisedb{}).Where("id=?",in.I).Find(&re).Error
	if err!=nil{
		log.Printf("GetExercisec> %s\n",err.Error())
		return &ExerciseData{},err
	}
	file,err:=os.Open(location)
	if err!=nil{
		log.Printf("GetExercisec> %s\n",err.Error())
		return &ExerciseData{},err
	}
	b,err:=ioutil.ReadAll(file)
	if err!=nil{
		log.Printf("GetExercisec> %s\n",err.Error())
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
	log.Printf("GetExercises: %+v\n",in)
	var A []int32//考试ID列表
	err:=e.db.Model(Exercisedb{}).Where("classid=?",in.I).Select("id").Find(&A).Error
	if err!=nil {
		log.Printf("GetExercises> %s\n",err.Error())
		return err
	}
	for _,i:=range A{
		re,err:=e.GetExercise(context.Background(),&I{I:i})
		if err!=nil {
			log.Printf("GetExercises> %s\n",err.Error())
			return err
		}
		err=stream.Send(re)
		if err!=nil {
			log.Printf("GetExercises> %s\n",err.Error())
			return err
		}
	}
	return nil
}
//添加一次考试
func(e *ExerciseService)AddExercise(c context.Context,in *ExerciseData) (*ExerciseData, error){
	log.Printf("AddExercise: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetExercise> timeout\n")
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
		err=tx.Model(Exercisedb{}).Where("id=?",in.Id).Update("location",location).Error//存入基本信息获取ID
		if err!=nil {
			return err
		}
		file,err:=os.Create(location)
		if err!=nil {
			return err
		}
		defer file.Close()
		_,err=file.WriteString(in.Content)
		if err!=nil {
			return err
		}
		return nil
	})
	if err!=nil {
		log.Printf("AddExercise> %s\n",err.Error())
	}
	return in,err
}
//学生提交一次考试记录
func(e *ExerciseService)SubmitAns(c context.Context,in *Submit) (*I, error){
	log.Printf("SubmitAns: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("SubmitAns> timeout\n")
		return &I{},errors.New("timeout")
	default:
	}
	//之前提交过了
	var tmp string
	err:=e.db.Model(Submitdb{}).Where(Submitdb{Exerciseid:in.Exerciseid,Uploaderid: in.Uploaderid}).Select("location").Find(&tmp).Error
	if 	tmp!=""{
		file,err:=os.Create(tmp)
		if err!=nil {
			log.Printf("SubmitAns> %s\n",err.Error())
			return &I{},err
		}
		defer file.Close()
		file.WriteString(in.Contents)
		return &I{},nil
	}
	//之前没有提交
	err=e.db.Transaction(func(tx *gorm.DB)error{
		tmp:=Submitdb{Uploaderid:in.Uploaderid, Exerciseid:in.Exerciseid}
		err=tx.Model(Submitdb{}).Save(&tmp).Error
		if err!=nil {
			log.Printf("SubmitAns> %s\n",err.Error())
			return err
		}
		in.Id=tmp.Id
		var location="./submit/"+strconv.Itoa(int(in.Id))+".sub"
		err=tx.Model(Submitdb{}).Where("id=?",in.Id).Update("location",location).Error
		if err!=nil {
			log.Printf("SubmitAns> %s\n",err.Error())
			return err
		}
		file,err:=os.Create(location)
		if err!=nil {
			log.Printf("SubmitAns> %s\n",err.Error())
			return err
		}
		defer file.Close()
		_,err=file.WriteString(in.Contents)
		if err!=nil {
			log.Printf("SubmitAns> %s\n",err.Error())
			return err
		}
		return nil
	})
	if err!=nil {
		log.Printf("SubmitAns> %s\n",err.Error())
	}
	return &I{},err
}
//根据考试ID获取答案
func(e *ExerciseService)GetKey(c context.Context,in *I) (*Submit, error){
	log.Printf("GetKey: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetKey> timeout\n")
		return &Submit{},errors.New("timeout")
	default:
	}
	var own int32
	err:=e.db.Model(Exercisedb{}).Where("id=?",in.I).Select("Ownerid").Find(&own).Error
	if err!=nil {
		log.Printf("GetKey> %s\n",err.Error())
		return &Submit{},err
	}
	var location string
	err=e.db.Model(Submitdb{}).Where(Submitdb{Uploaderid: own,Exerciseid: in.I}).Select("location").Find(&location).Error
	if err!=nil {
		log.Printf("GetKey> %s\n",err.Error())
		return &Submit{},err
	}
	if location == ""{//还没有上传
		return &Submit{},nil
	}
	file,err:=os.Open(location)
	if err!=nil {
		log.Printf("GetKey> %s\n",err.Error())
		return &Submit{},err
	}
	defer file.Close()
	b,err:=ioutil.ReadAll(file)
	if err!=nil {
		log.Printf("GetKey> %s\n",err.Error())
		return &Submit{},err
	}
	re:=Submit{}
	err=e.db.Model(Submitdb{}).Where("Uploaderid=?",own).Find(&re).Error
	if err!=nil {
		log.Printf("GetKey> %s\n",err.Error())
		return &Submit{},err
	}
	re.Contents=string(b)
	return &re,err
}
//给学生打分
func(e *ExerciseService)SetScore(c context.Context,in *Score) (*Empty1, error){
	log.Printf("SetScore: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("SetScore> timeout\n")
		return &Empty1{},errors.New("timeout")
	default:
	}
	err:=e.db.Model(Submitdb{}).Where("id=?",in.Submitid).Update("value",in.Value).Error
	if err!=nil {
		log.Printf("SetScore> %s\n",err.Error())
	}
	return &Empty1{},err
}
//学生根据提交记录获取本次得分
func(e *ExerciseService)GetScore(c context.Context,in *I) (*Score, error){
	log.Printf("GetScore: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("GetScore> timeout\n")
		return &Score{},errors.New("timeout")
	default:
	}
	var v int32
	err:=e.db.Model(Submitdb{}).Where("id=?",in.I).Select("value").Find(&v).Error
	if err!=nil {
		log.Printf("GetScore> %s\n",err.Error())
	}
	return &Score{Value: v},err
}
//学生根据自己的ID获取自己的所有提交记录
func(e *ExerciseService)GetScores(in *I,stream Exercise_GetScoresServer) error{
	log.Printf("GetScores: %+v\n",in)
	var re []Submitdb
	err:=e.db.Model(Submitdb{}).Where("Uploaderid=?",in.I).Find(&re).Error
	if err!=nil {
		log.Printf("GetScores> %s\n",err.Error())
		return err
	}
	for _,i:=range re{
		file,err:=os.Open(i.Location)
		if err!=nil {
			log.Printf("GetScores> %s\n",err.Error())
			return err
		}
		b,err:=ioutil.ReadAll(file)
		if err!=nil {
			log.Printf("GetScores> %s\n",err.Error())
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
			log.Printf("GetScores> %s\n",err.Error())
			return err
		}
	}
	return nil
}
//老师根据考试ID获取本次班级得分情况
func(e *ExerciseService)GetClassScores(in *I,stream Exercise_GetClassScoresServer) error{
	log.Printf("GetClassScores: %+v\n",in)
	var ids []int32
	err:=e.db.Model(Submitdb{}).Where("Exerciseid=?",in.I).Select("id").Find(&ids).Error
	if err!=nil {
		log.Printf("GetClassScores> %s\n",err.Error())
		return err
	}
	var v int32
	for _,i:=range ids{
		err:=e.db.Model(Submitdb{}).Where("id=?",i).Select("value").Find(&v).Error
		if err!=nil {
			log.Printf("GetClassScores> %s\n",err.Error())
			return err
		}
		err=stream.Send(&Score{Submitid:i,Value: v})
		if err!=nil {
			log.Printf("GetClassScores> %s\n",err.Error())
			return err
		}
	}
	return nil
}
//老师根据考试ID获取本次提交情况
func(e *ExerciseService)GetClassSubmit(in *I,stream Exercise_GetClassSubmitServer) error{
	var s []Submitdb
	err:=e.db.Model(Submitdb{}).Where("Exerciseid=?",in.I).Find(&s).Error
	if err!=nil {
		return err
	}
	for _,i:=range s{
		file,err:=os.Open(i.Location)
		if err!=nil {
			return err
		}
		b,err:=ioutil.ReadAll(file)
		if err!=nil {
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
	log.Printf("DelExercise: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("DelExercise> timeout\n")
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
	return &Empty1{},err
}
//根据班级ID删除该班级所有记录
func(e *ExerciseService)DelExercises(c context.Context,in *I) (*Empty1, error){
	log.Printf("DelExercise: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("DelExercise> timeout\n")
		return &Empty1{},errors.New("timeout")
	default:
	}
	//获取所有考试列表
	err:=e.db.Transaction(func(tx *gorm.DB)error{
		var A []int32//考试ID列表
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
	return &Empty1{},err
}

func(e *ExerciseService)mustEmbedUnimplementedExerciseServer(){}