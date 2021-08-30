package NetworkDisk

import (
	"context"
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"io"
	"log"
	"os"
)


type NetworkDiskService struct{
	db *gorm.DB
}

func newNetworkDiskService() *NetworkDiskService {
	InitGorm()
	return &NetworkDiskService{sql}
}

func (n *NetworkDiskService)Download(in *Fileid,stream NetworkDisk_DownloadServer) error{
	log.Printf("Download: %+v\n",in)

	tmp :=fileinfodb{}
	err:=n.db.Model(fileinfodb{}).Where("id=?",in.Id).Find(&tmp).Error
	if err!=nil {
		log.Printf("Download> %s\n",err.Error())
		return err
	}

	info:=Fileinfo{
		Id: tmp.Id,
		Uploader:tmp.Uploader,
		Classid:tmp.Classid,
		Name:tmp.Name,
		Size:tmp.Size,
		Time:tmp.Time,
	}
	b,err:=json.Marshal(info)
	if err!=nil {
		log.Printf("Download> %s\n",err.Error())
		return err
	}
	stream.Send(&Filestream{Content: b})//先把文件信息法过去
	file,err:=os.Open(tmp.Location)
	if err!=nil {
		log.Printf("Download> %s\n",err.Error())
		return err
	}
	defer file.Close()
	uints:=make([]byte,TransmissionUnit)
	for{
		_,err=file.Read(uints)
		if err==io.EOF{
			break
		}
		if err!=nil {
			log.Printf("Download> %s\n",err.Error())
			return err
		}
		err=stream.Send(&Filestream{Content: uints})
		if err!=nil {
			log.Printf("Download> %s\n",err.Error())
			return err
		}
	}
	return nil
}

func (n *NetworkDiskService)Upload(stream NetworkDisk_UploadServer) error {
	in := Fileinfo{}
	b, err := stream.Recv()
	if err != nil {
		return err
	}
	err = json.Unmarshal(b.Content, &in)
	if err != nil {
		return err
	}
	tmp := fileinfodb{Uploader: in.Uploader, Classid: in.Classid, Name: in.Name, Size: in.Size, Time: in.Time}
	err = n.db.Transaction(func(tx *gorm.DB) error {
		err = tx.Model(fileinfodb{}).Create(&tmp).Error
		if err != nil {
			return err
		}
		tmp.Location = "./files/" + tmp.Name
		err = tx.Model(fileinfodb{}).Where("id=?", tmp.Id).Update("location", tmp.Location).Error
		if err != nil {
			return err
		}
		in.Id = tmp.Id
		file, err := os.Create(tmp.Location)
		if err != nil {
			return err
		}
		defer file.Close()
		for {
			b, err = stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
			file.Write(b.Content)
		}
		return nil
	})
	if err != nil {
		log.Printf("Upload> %s\n", err.Error())
	}
	return err
}

func (n *NetworkDiskService)GetContents(c context.Context,in *Classid) (*Contents, error){
	log.Printf("Upload: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("Upload> timeout\n")
		return &Contents{},errors.New("timeout")
	default:
	}
	var tmp []fileinfodb
	err:=n.db.Model(fileinfodb{}).Where("classid=?",in.Id).Find(&tmp).Error
	if err!=nil {
		log.Printf("GetContents> %s\n",err.Error())
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