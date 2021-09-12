package Log

import (
	"os"
	"time"
)

//刷新文件-按日期记录
func flushFile(dir string,file *os.File) (*os.File,error) {
	str:=dir+time.Now().Format("2006-01-02")
	if file == nil {
		re,err:=os.Create(str)
		if err!=nil {
			return file,err
		}
		return re,nil
	}else if file.Name()!=str {
		file.Close()
		re,err:=os.Create(str)
		if err!=nil {
			return file,err
		}
		return re,nil
	}else{
		return file,nil
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		panic(err)
		//log.Fatalf("%s: %s", msg, err)
	}
}

func Run(){
	go recieveInfo()
	go recieveError()
}