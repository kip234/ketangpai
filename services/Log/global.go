package Log

//package.method.type....
//type : info/err...

const mqAddr="amqp://kip:kip@121.4.76.240:5672/"

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
