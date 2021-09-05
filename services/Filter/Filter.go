package Filter

import (
	"KeTangPai/services/Filter/prefix_tree"
	"KeTangPai/services/Log"
	"context"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

type Filter struct{
	tree    *prefix_tree.Prefix_tree
	replace byte
}

//临时添加敏感词
func (f *Filter)Add(c context.Context,arg *FilterData)(*FilterData, error){
	Log.Send("Filter.Add.info",arg)
	select {
	case <-c.Done():
		Log.Send("Filter.Add.error","timeout")
		return &FilterData{},errors.New("timeout")
	default:
	}
	f.tree.Add(arg.Data)
	return &FilterData{},nil
}

//处理敏感词
func (f *Filter)Process(c context.Context,arg *FilterData) (*FilterData, error) {
	Log.Send("Filter.Process.info",arg)
	select {
	case <-c.Done():
		Log.Send("Filter.Process.error","timeout")
		//log.Printf("Process> timeout\n")
		return &FilterData{},errors.New("timeout")
	default:
	}
	content:=arg.Data
	var re []byte
	location:=f.tree.Find(content)
	index1:=0//content下标
	index2:=0//location下标
	length:=len(content)
	nums:=len(location)
	if nums == 0{//没有出现
		return &FilterData{Data: content},nil
	}

	for index1<length{
		if index2<nums && index1==location[index2][0]{//开始出现违规文字
			for index1<location[index2][1]{
				re=append(re,f.replace)
				if content[index1]>127{
					index1+=3//中文
				}else{
					index1+=1
				}
			}
			index2+=1
		}else{
			re=append(re,content[index1])
			index1+=1
		}
	}
	return &FilterData{Data:re},nil
}

//依靠SensitiveWords指向的文件内容创建一个过滤器
func newFilter() (error,*Filter) {
	file,err:=os.Open(SensitiveWords)
	if err!=nil{
		return err,nil
	}
	b,err:=ioutil.ReadAll(file)
	if err!=nil{
		return err,nil
	}
	s:=string(b)
	words:=strings.Split(s,"\r\n")

	f:= Filter{
		tree: prefix_tree.NewPrefix_tree(),
	}
	f.replace=[]byte(words[0])[0]
	i:=1
	nums:=len(words)
	for i<nums{
		f.tree.Add([]byte(words[i]))
		//f.Add([]byte(words[i]))
		i+=1
	}
	return nil,&f
}

