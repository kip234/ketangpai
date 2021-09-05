package Handlers

import (
	"KeTangPai/services/DC/NetworkDisk"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"time"
)

const TransmissionUnit=1024//bytes-每次传输的单位大小

func File_upload(n NetworkDisk.NetworkDiskClient)gin.HandlerFunc{
	return func(c *gin.Context){
		uid,err:= getInt("uid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		classid,err:= getInt("classid",c)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		tmp:=NetworkDisk.Fileinfo{
			Uploader: uint32(uid),
			Classid: uint32(classid),
			Time: time.Now().Unix(),
		}
		file,err:=c.FormFile("file")//获取上传的文件
		if err!=nil{
			c.JSON(http.StatusBadRequest,gin.H{
				"error": err.Error(),
			})
			return
		}
		tmp.Name=file.Filename
		tmp.Size=uint64(file.Size)
		tmp.Unit=TransmissionUnit

		b,err:=json.Marshal(tmp)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		stream,err:=n.Upload(ctx)
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error": err.Error(),
			})
			return
		}
		defer stream.CloseSend()
		err=stream.Send(&NetworkDisk.Filestream{Content: b})//文件信息发过去
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error": err.Error(),
			})
			return
		}

		src,err:=file.Open()
		if err!=nil{
			c.JSON(http.StatusInternalServerError,gin.H{
				"error": err.Error(),
			})
			return
		}

		units:=make([]byte,tmp.Unit)
		for {
			_,err=src.Read(units)
			if err==io.EOF {
				break
			}
			if err!=nil{
				c.JSON(http.StatusInternalServerError,gin.H{
					"error": err.Error(),
				})
				return
			}

			err=stream.Send(&NetworkDisk.Filestream{Content: units})//文件信息发过去
			if err!=nil{
				c.JSON(http.StatusInternalServerError,gin.H{
					"error": err.Error(),
				})
				return
			}
		}
	}
}
