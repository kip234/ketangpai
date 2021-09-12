package RankingList

import (
	"KeTangPai/Models/Redis"
	"KeTangPai/services/Log"
	"context"
	"errors"
)

type RankingListService struct {
	pool *Redis.RedisPool
}

func newRankingListService() (re *RankingListService) {
	re = &RankingListService{
		pool: &DefaultRedis,
	}
	re.pool.Init()
	return
}

//刷新记录
func (r *RankingListService)Flushlist(c context.Context,in *Flushin) (f *Flushout,err error){
	Log.Send("RankingList.Flushlist.info",in)
	defer func(){
		if err!=nil {
			Log.Send("RankingList.Flushlist.error",err.Error())
		}
	}()
	select {
	case <-c.Done():
		Log.Send("RankingList.Flushlist.error","timeout")
		return &Flushout{},errors.New("timeout")
	default:
	}
	err=r.pool.ZINCRBY(in.Key,in.Increment,in.Member)
	if err!=nil {
		return &Flushout{},err
	}
	rank,err:=r.pool.ZREVRANK(in.Key,in.Member)
	if err!=nil {
		return &Flushout{},err
	}
	return &Flushout{Member: in.Member,Ranking: rank},err
}

//获取榜单信息
func (r *RankingListService)Getlistinfo(c context.Context,in *Listname) (l *Listinfo,err error){
	Log.Send("RankingList.Getlistinfo.info",in)
	defer func(){
		if err!=nil {
			Log.Send("RankingList.Getlistinfo.error",err.Error())
		}
	}()
	select {
	case <-c.Done():
		Log.Send("RankingList.Getlistinfo.error","timeout")
		return &Listinfo{},errors.New("timeout")
	default:
	}
	re,err:=r.pool.ZREVRANGE(in.Name,0,-1)
	if err!=nil {
		return &Listinfo{},err
	}
	return &Listinfo{Name: in.Name,List:re},err
}

//删除榜单
func (r *RankingListService)Dellist(c context.Context,in *Listname) (e *Empty,err error){
	defer func(){
		if err!=nil {
			Log.Send("RankingList.Dellist.error",err.Error())
		}
	}()
	Log.Send("RankingList.Dellist.info",in)
	select {
	case <-c.Done():
		Log.Send("RankingList.Dellist.error","timeout")
		return &Empty{},errors.New("timeout")
	default:
	}
	err=r.pool.ZREMRANGEBYRANK(in.Name,0,-1)//删除所有记录
	return &Empty{},err
}

//获取排名
func (r *RankingListService)Getranking(c context.Context,in *Members) (ranks *Rankings,err error){
	defer func(){
		if err!=nil {
			Log.Send("RankingList.Getranking.error",err.Error())
		}
	}()
	Log.Send("RankingList.Getranking.info",in)
	select {
	case <-c.Done():
		Log.Send("RankingList.Getranking.error","timeout")
		return &Rankings{},errors.New("timeout")
	default:
	}
	ranks.Rank=make([]int64,len(in.Members))
	for index,i:=range in.Members{
		rank,err:=r.pool.ZREVRANK(in.Name,i)
		if err!=nil {
			return &Rankings{},err
		}
		ranks.Rank[index]=rank
	}
	return
}

func (r *RankingListService)mustEmbedUnimplementedRankingListServer(){}
