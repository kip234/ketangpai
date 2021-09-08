package TestBank

import (
	"KeTangPai/services/Log"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type TestBankService struct{
	db *gorm.DB
}

func newTestBankService() *TestBankService {
	InitGorm()
	return &TestBankService{sql}
}

//上传测试题目
func (t *TestBankService)Upload(c context.Context,in *Test) (*Test, error){
	Log.Send("TestBank.Upload.info",in)
	select {
	case <-c.Done():
		Log.Send("TestBank.Upload.error","timeout")
		return &Test{}, errors.New("timeout")
	default:
	}
	tmp:=Testdb{Typ: in.Typ,Content: in.Content,Ans: in.Ans,Name:in.Name,Uploader: in.Uploader,Discipline: in.Discipline}
	err:=t.db.Transaction(func(tx *gorm.DB)error{
		err:=tx.Model(Testdb{}).Create(&tmp).Error//储存基础信息
		if err!=nil {
			return err
		}
		tmp.Location="./test/"+strconv.Itoa(int(tmp.Id))+".tst"
		//更新储存位置
		err=tx.Model(Testdb{}).Where("id=?",tmp.Id).Update("location",tmp.Location).Error
		if err!=nil {
			return err
		}
		file,err:=os.Create(tmp.Location)
		if err!=nil {
			return err
		}
		defer file.Close()
		b,err:=json.Marshal(tmp)//以JSON格式储存测试题目信息
		if err!=nil {
			return err
		}
		file.Write(b)
		return nil
	})
	if err!=nil {
		Log.Send("TestBank.Upload.error",err.Error())
		return &Test{},err
	}
	in.Id=tmp.Id
	return in,nil
}

//下载题目内容-可下载多个题目
func (t *TestBankService)Download(stream TestBank_DownloadServer) error{
	for{
		id,err:=stream.Recv()//接收测试题的ID
		if err==io.EOF{
			return nil
		}
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
		var location string
		//寻找测试题储存位置
		err=t.db.Model(Testdb{}).Where("id=?",id.Id).Select("location").Find(&location).Error
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
		//开始读取记录
		file,err:=os.Open(location)
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
		c,err:=ioutil.ReadAll(file)
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
		T:=Test{}
		err=json.Unmarshal(c,&T)
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
		err=stream.Send(&T)
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
	}
}

//自动生成一套试卷
func (t *TestBankService)GenerateTest(in *Testconf,stream TestBank_GenerateTestServer) error{
	Log.Send("TestBank.GenerateTest.info",in)
	var  subjective []uint32
	var  objective []uint32
	err:=t.db.Model(Testdb{}).Where(Testdb{Typ: Subjective,Discipline: in.Discipline}).Select("id").Find(&subjective).Error
	if err!=nil {
		Log.Send("TestBank.GenerateTest.error",err.Error())
		return err
	}
	err=t.db.Model(Testdb{}).Where(Testdb{Typ: Objective,Discipline: in.Discipline}).Select("id").Find(&objective).Error
	if err!=nil {
		Log.Send("TestBank.GenerateTest.error",err.Error())
		return err
	}
	//开始‘随机’选取题目

	step:=len(objective)/int(in.ObjectiveItem)
	for in.ObjectiveItem > 0{
		rand.Seed(time.Now().UnixNano())
		err=stream.Send(&Testid{Id: objective[step*(int(in.ObjectiveItem)-1)+rand.Intn(step)]})
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return err
		}
		in.ObjectiveItem-=1
	}

	step=len(subjective)/int(in.SubjectiveItem)
	for in.SubjectiveItem > 0{
		rand.Seed(time.Now().UnixNano())
		err=stream.Send(&Testid{Id: subjective[step*(int(in.SubjectiveItem)-1)+rand.Intn(step)]})
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return err
		}
		in.SubjectiveItem-=1
	}
	return nil
}

func (t *TestBankService)mustEmbedUnimplementedTestBankServer(){}
