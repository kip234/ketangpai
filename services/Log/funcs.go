package Log

import (
	"log"
	"os"
	"time"
)

//func Init(){
//	conn, err := amqp.Dial(mqAddr)
//	failOnError(err, "Failed to connect to RabbitMQ")
//	defer conn.Close()
//
//	send, err = conn.Channel()
//	failOnError(err, "Failed to open a channel")
//	//defer send.Close()
//
//	err = send.ExchangeDeclare(
//		exchangeName, // name
//		exchangeType,      // type
//		true,         // durable
//		false,        // auto-deleted
//		false,        // internal
//		false,        // no-wait
//		nil,          // arguments
//	)
//	failOnError(err, "Failed to declare an exchange")
//}

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
		log.Fatalf("%s: %s", msg, err)
	}
}

func Run(){
	go recieveInfo()
	go recieveError()
}