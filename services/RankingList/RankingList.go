package RankingList

import (
	"KeTangPai/Models/Redis"
	"context"
	"errors"
	"log"
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

func (r *RankingListService)Flushlist(c context.Context,in *Flushin) (f *Flushout,err error){
	defer func(){
		if err!=nil {
			log.Printf("Flushlist> %s\n",err.Error())
		}
	}()
	log.Printf("Flushlist: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("Flushlist> timeout\n")
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

func (r *RankingListService)Getlistinfo(c context.Context,in *Listname) (l *Listinfo,err error){
	defer func(){
		if err!=nil {
			log.Printf("Getlistinfo> %s\n",err.Error())
		}
	}()
	log.Printf("Getlistinfo: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("Getlistinfo> timeout\n")
		return &Listinfo{},errors.New("timeout")
	default:
	}
	re,err:=r.pool.ZREVRANGE(in.Name,0,-1)
	if err!=nil {
		return &Listinfo{},err
	}
	//l.List=re
	//l.Name=in.Name
	return &Listinfo{Name: in.Name,List:re},err
}

func (r *RankingListService)Dellist(c context.Context,in *Listname) (e *Empty,err error){
	defer func(){
		if err!=nil {
			log.Printf("Dellist> %s\n",err.Error())
		}
	}()
	log.Printf("Dellist: %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("Dellist> timeout\n")
		return &Empty{},errors.New("timeout")
	default:
	}
	err=r.pool.ZREMRANGEBYRANK(in.Name,0,-1)//删除所有记录
	return &Empty{},err
}

func (r *RankingListService)Getranking(c context.Context,in *Members) (ranks *Rankings,err error){
	defer func(){
		if err!=nil {
			log.Printf("Getranking: %s\n",err.Error())
		}
	}()
	log.Printf("Getranking> %+v\n",in)
	select {
	case <-c.Done():
		log.Printf("Getranking> timeout\n")
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
