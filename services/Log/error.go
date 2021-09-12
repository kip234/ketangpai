package Log

import (
	"github.com/streadway/amqp"
	"time"
)

func recieveError(){
	conn, err := amqp.Dial(mqAddr)
	failOnError(err,"mq connect error")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err,"channel error")
	defer ch.Close()

	err = ch.ExchangeDeclare(
		exchangeName,
		exchangeType,
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err,"exchange error")
	q, err := ch.QueueDeclare(
		"",
		false,
		false,
		true,
		false,
		nil,
	)
	failOnError(err,"queue error")

	err = ch.QueueBind(
		q.Name,
		errorRoutingKey,
		exchangeName,
		false,
		nil,
	)
	failOnError(err, "bind error")

	err = ch.Qos(
		1,
		0,
		false,
	)
	failOnError(err, "set QoS error")

	msgs, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err,"Consume error")

	forever := make(chan bool)
	file,err:=flushFile(errorDir,nil)
	failOnError(err,"")
	defer file.Close()
	go func() {
		for d := range msgs {
			_,err=file.WriteString(time.Now().Format("15:04:05")+" "+d.RoutingKey+" "+string(d.Body))
			failOnError(err,"Write error")
			file,err=flushFile(errorDir,file)
			failOnError(err,"flushFile error")
			err=d.Ack(false)
			failOnError(err,"Ack error")
		}
	}()

	<-forever
}