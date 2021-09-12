package Handlers

import (
	"KeTangPai/services/DC/NetworkDisk"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"strconv"
	"time"
)

func File_download(n NetworkDisk.NetworkDiskClient)gin.HandlerFunc{
	return func(c *gin.Context){
		s,ok:=c.GetQuery("fileid")
		if !ok{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"missing necessary parameter",
			})
			return
		}
		if s==""{
			c.JSON(http.StatusBadRequest,gin.H{
				"error":"fileid == nil",
			})
			return
		}
		fileid,err:=strconv.Atoi(s)
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}
		ctx,_:=context.WithTimeout(context.Background(),serviceTimeLimit*time.Second)
		stream,err:=n.Download(ctx,&NetworkDisk.Fileid{Id: uint32(fileid)})
		b,err:=stream.Recv()
		if err!=nil {
			c.JSON(http.StatusBadRequest,gin.H{
				"error":err.Error(),
			})
			return
		}

		info:=NetworkDisk.Fileinfo{}
		err=json.Unmarshal(b.Content,&info)
		if err!=nil {
			c.JSON(http.StatusInternalServerError,gin.H{
				"error":err.Error(),
			})
			return
		}
		content:=make([]byte,info.Size)
		var index uint64
		for  {
			f,err:=stream.Recv()
			if err==io.EOF{
				break
			}
			if err!=nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}
			copy(content[index:],f.Content)
			index+= info.Unit
		}
		c.Header("Content-Disposition","attachment;filename="+info.Name)
		_,err=c.Writer.Write(content)
		if err!=nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
	}
}
