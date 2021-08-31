package Log

//package.method.type....
//type : info/err...

const mqAddr="amqp://#$%:*^&@:5672/"

//交换机信息
const (
	exchangeType = "topic"
	exchangeName="logs"
)

//
const (
	infoRoutingKey = "#.info.#"
	errorRoutingKey = "#.error.#"
)

//
const (
	infoDir = "./logs/info/"
	errorDir="./logs/error/"
)
