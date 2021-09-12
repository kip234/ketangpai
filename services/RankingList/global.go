package RankingList

import "KeTangPai/Models/Redis"

const Addr=":8091"

var DefaultRedis = Redis.RedisPool{
	Read 		:"localhost:6379",
	Write 		:"localhost:6379",
	IdLeTimeout	:5,
	MaxIdle		:20,
	MaxActive	:8,
}
