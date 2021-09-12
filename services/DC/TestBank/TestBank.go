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
	tmp:=Testdb{
		Typ: in.Typ,
		Content: in.Content,
		Ans: in.Ans,
		Name:in.Name,
		Uploader: in.Uploader,
		Discipline: in.Discipline,
		Withans: in.Withans,
	}
	err:=t.db.Transaction(func(tx *gorm.DB)error{
		err:=tx.Model(Testdb{}).Create(&tmp).Error//储存基础信息
		if err!=nil {
			return err
		}
		//储存内容
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
		_,err=file.WriteString(tmp.Content)
		file.Close()
		if err!=nil {
			return err
		}

		//储存答案
		tmp.AnsLocation="./test/"+strconv.Itoa(int(tmp.Id))+".ans"
		//更新储存位置
		err=tx.Model(Testdb{}).Where("id=?",tmp.Id).Update("ans_location",tmp.AnsLocation).Error
		if err!=nil {
			return err
		}
		file,err=os.Create(tmp.AnsLocation)
		if err!=nil {
			return err
		}
		_,err=file.WriteString(tmp.Ans)
		file.Close()
		if err!=nil {
			return err
		}
		return nil
	})
	if err!=nil {
		Log.Send("TestBank.Upload.error",err.Error())
		return &Test{},err
	}
	in.Id=tmp.Id
	//Log.Send("TestBank.Upload.error",err.Error())
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

		//寻找测试题储存位置
		T:=Testdb{}
		err=t.db.Model(Testdb{}).Where("id=?",id.Id).Find(&T).Error
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
		//开始读取记录
		file,err:=os.Open(T.Location)
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
		c,err:=ioutil.ReadAll(file)
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
		T.Content=string(c)

		err=stream.Send(&Test{
			Id: T.Id,
			Typ: T.Typ,
			Content: T.Content,
			Ans: T.Ans,
			Name: T.Name,
			Uploader: T.Uploader,
			Discipline: T.Discipline,
			Withans: T.Withans,
		})
		if err!=nil{
			Log.Send("TestBank.Download.error",err.Error())
			return err
		}
	}
}

//自动生成一套试卷
func (t *TestBankService)GenerateTest(c context.Context,in *Testconf) (*Tests, error){
	Log.Send("TestBank.GenerateTest.info",in)
	var  subjective []uint32
	var  objective []uint32
	err:=t.db.Model(Testdb{}).Where(Testdb{Typ: Subjective,Discipline: in.Discipline}).Select("id").Find(&subjective).Error
	if err!=nil {
		Log.Send("TestBank.GenerateTest.error",err.Error())
		return &Tests{},err
	}
	err=t.db.Model(Testdb{}).Where(Testdb{Typ: Objective,Discipline: in.Discipline}).Select("id").Find(&objective).Error
	if err!=nil {
		Log.Send("TestBank.GenerateTest.error",err.Error())
		return &Tests{},err
	}
	//开始‘随机’选取题目

	re:=Tests{Tests: make([]string,in.SubjectiveItem+in.ObjectiveItem)}
	ans:=make(map[uint32]string)

	index:=0
	step:=len(objective)/int(in.ObjectiveItem)
	for in.ObjectiveItem > 0{
		rand.Seed(time.Now().UnixNano())
		id:=objective[step*(int(in.ObjectiveItem)-1)+rand.Intn(step)]
		tmp:=Testdb{}
		err=t.db.Model(Testdb{}).Where("id=?",id).Find(&tmp).Error
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		//读取题目
		file,err:=os.Open(tmp.Location)
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		b,err:=ioutil.ReadAll(file)
		file.Close()
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		tmp.Content=string(b)
		//获取答案
		file,err=os.Open(tmp.AnsLocation)
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		b,err=ioutil.ReadAll(file)
		file.Close()
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		ans[id]=string(b)

		b,err=json.Marshal(tmp)
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}

		re.Tests[index]=string(b)
		index++
		in.ObjectiveItem-=1
	}

	step=len(subjective)/int(in.SubjectiveItem)
	for in.SubjectiveItem > 0{
		rand.Seed(time.Now().UnixNano())
		id:=subjective[step*(int(in.SubjectiveItem)-1)+rand.Intn(step)]

		tmp:=Testdb{}
		err=t.db.Model(Testdb{}).Where("id=?",id).Find(&tmp).Error
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		//读取内容
		file,err:=os.Open(tmp.Location)
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		b,err:=ioutil.ReadAll(file)
		file.Close()
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		tmp.Content=string(b)
		//读取答案
		file,err=os.Open(tmp.AnsLocation)
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		b,err=ioutil.ReadAll(file)
		file.Close()
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}
		ans[id]=string(b)

		b,err=json.Marshal(tmp)
		if err!=nil {
			Log.Send("TestBank.GenerateTest.error",err.Error())
			return &Tests{},err
		}

		re.Tests[index]=string(b)
		index++
		in.SubjectiveItem-=1
	}
	b,err:=json.Marshal(ans)
	if err!=nil {
		Log.Send("TestBank.GenerateTest.error",err.Error())
		return &Tests{},err
	}
	re.Ans=b
	return &re,nil
}

func (t *TestBankService)GetAns(c context.Context,in *Testids) (*Anss, error){
	Log.Send("TestBank.GetAns.info",in)
	select {
	case <-c.Done():
		Log.Send("TestBank.GetAns.error","timeout")
		return &Anss{}, errors.New("timeout")
	default:
	}
	anss:=make([]string,len(in.Testids))
	for i,_:=range in.Testids{
		var ans string
		err:=t.db.Model(Testdb{}).Where("id=?",in.Testids[i]).Select("ans_location").Find(&ans).Error
		if err!=nil {
			Log.Send("TestBank.GetAns.error",err.Error())
			return &Anss{}, err
		}
		file,err:=os.Open(ans)
		if err!=nil {
			Log.Send("TestBank.GetAns.error",err.Error())
			return &Anss{}, err
		}
		b,err:=ioutil.ReadAll(file)
		if err!=nil {
			Log.Send("TestBank.GetAns.error",err.Error())
			file.Close()
			return &Anss{}, err
		}
		anss[i]=string(b)
	}
	return &Anss{Anss: anss},nil
}

func (t *TestBankService)mustEmbedUnimplementedTestBankServer(){}
