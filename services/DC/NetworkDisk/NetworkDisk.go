package NetworkDisk

import (
	"KeTangPai/services/Log"
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io"
	"os"
)


type NetworkDiskService struct{
	db *gorm.DB
}

func newNetworkDiskService() *NetworkDiskService {
	InitGorm()
	return &NetworkDiskService{sql}
}

//凭fileID下载文件
func (n *NetworkDiskService)Download(in *Fileid,stream NetworkDisk_DownloadServer) error{
	Log.Send("NetworkDisk.Download.info",in)

	tmp :=fileinfodb{}
	//凭fileID获取文件信息
	err:=n.db.Model(fileinfodb{}).Where("id=?",in.Id).Find(&tmp).Error
	if err!=nil {
		Log.Send("NetworkDisk.Download.error","timeout")
		return err
	}

	info:=Fileinfo{
		Id: tmp.Id,
		Uploader:tmp.Uploader,
		Classid:tmp.Classid,
		Name:tmp.Name,
		Size:tmp.Size,
		Time:tmp.Time,
		Unit: TransmissionUnit,
	}
	b,err:=json.Marshal(info)
	if err!=nil {
		Log.Send("NetworkDisk.Download.error",err.Error())
		return err
	}
	//发送文件信息-(应该把每次发送的大小一并发出去、但我这里并没有)
	stream.Send(&Filestream{Content: b})//先把文件信息法过去
	file,err:=os.Open(tmp.Location)
	if err!=nil {
		Log.Send("NetworkDisk.Download.error",err.Error())
		return err
	}
	defer file.Close()
	//按照规定一点一点发
	uints:=make([]byte,info.Unit)
	for{
		_,err=file.Read(uints)
		if err==io.EOF{
			break
		}
		if err!=nil {
			Log.Send("NetworkDisk.Download.error",err.Error())
			return err
		}
		err=stream.Send(&Filestream{Content: uints})
		if err!=nil {
			Log.Send("NetworkDisk.Download.error",err.Error())
			return err
		}
	}
	return nil
}

//上传文件
func (n *NetworkDiskService)Upload(stream NetworkDisk_UploadServer) error {
	in := Fileinfo{}
	//先接收文件的信息
	b, err := stream.Recv()
	if err != nil {
		Log.Send("NetworkDisk.Upload.error",err.Error())
		return err
	}
	err = json.Unmarshal(b.Content, &in)
	if err != nil {
		Log.Send("NetworkDisk.Upload.error",err.Error())
		return err
	}

	tmp := fileinfodb{Uploader: in.Uploader, Classid: in.Classid, Name: in.Name, Size: in.Size, Time: in.Time}
	err = n.db.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(fileinfodb{}).Create(&tmp).Error
		if err != nil {
			Log.Send("NetworkDisk.Upload.error",err.Error())
			return err
		}
		tmp.Location = "./files/" + tmp.Name
		//更新储存位置
		err = tx.Model(fileinfodb{}).Where("id=?", tmp.Id).Update("location", tmp.Location).Error
		if err != nil {
			Log.Send("NetworkDisk.Upload.error",err.Error())
			return err
		}
		in.Id = tmp.Id
		file, err := os.Create(tmp.Location)
		if err != nil {
			Log.Send("NetworkDisk.Upload.error",err.Error())
			return err
		}
		defer file.Close()
		for {
			b, err = stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				Log.Send("NetworkDisk.Upload.error",err.Error())
				return err
			}
			file.Write(b.Content)
		}
		return nil
	})
	if err != nil {
		Log.Send("NetworkDisk.Upload.error",err.Error())
	}
	return err
}

//获取文件目录
func (n *NetworkDiskService)GetContents(c context.Context,in *Classid) (*Contents, error){
	Log.Send("NetworkDisk.GetContents.info",in)
	select {
	case <-c.Done():
		Log.Send("NetworkDisk.GetContents.error","timeout")
		return &Contents{},errors.New("timeout")
	default:
	}
	var tmp []fileinfodb
	//寻找对应班级的文件记录
	err:=n.db.Model(fileinfodb{}).Where("classid=?",in.Id).Find(&tmp).Error
	if err!=nil {
		Log.Send("NetworkDisk.GetContents.error",err.Error())
		return &Contents{},err
	}
	var re Contents
	for _,i:=range tmp{
		re.Id=append(re.Id,i.Id)
		re.Name=append(re.Name,i.Name)
	}
	return &re,nil
}

func (n *NetworkDiskService)mustEmbedUnimplementedNetworkDiskServer(){}