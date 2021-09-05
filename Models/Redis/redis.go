package Redis

import (
	"github.com/gomodule/redigo/redis"
)

//这里声明了一个带初始化配置的Redis连接结构

type RedisPool struct {
	Read 		string
	Write 		string
	IdLeTimeout	int
	MaxIdle		int
	MaxActive	int
	rpool *redis.Pool//负责读取
	wpool *redis.Pool//负责写入
}