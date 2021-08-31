package Log

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func Send(routingKey string,body interface{}) error {
	message:=fmt.Sprintf("%+v\n",body)



	conn, err := amqp.Dial(mqAddr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	send, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	//defer send.Close()

	err = send.ExchangeDeclare(
		exchangeName, // name
		exchangeType,      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")



	err = send.Publish(
		exchangeName,          // exchange
		routingKey, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err!=nil {
		log.Printf("Send> routingKey: %s  error:%s\n",routingKey,err.Error())
	}
	return err
}