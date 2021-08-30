package JWT

import (
	"KeTangPai/Models/Redis"
)

const Secret 	= "I'mTooDishes"//JWT秘钥
const Addr		= ":8083"
var DefaultJwt = Jwt{
	Header: Header{
		Alg: "HS256",
		Typ: "JWT",
	},
	Payload: Payload{
		Iss: "kip",
		Sub: "KeTangPai",
	},
	Secret: Secret,
}

var DefaultRedis = Redis.RedisPool{
	Read 		:"localhost:6379",
	Write 		:"localhost:6379",
	IdLeTimeout	:5,
	MaxIdle		:20,
	MaxActive	:8,
}