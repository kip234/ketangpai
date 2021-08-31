package Log

import (
	"github.com/streadway/amqp"
	"log"
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
		exchangeName, // name
		exchangeType,      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err,"exchange error")
	q, err := ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err,"queue error")

	err = ch.QueueBind(
		q.Name, // queue name
		errorRoutingKey,     // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	failOnError(err, "bind error")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "set QoS error")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
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

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}