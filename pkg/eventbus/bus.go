package eventbus

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
)

type Event interface {
	ToMessage() string
}

func Publish(e Event) {
	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(error); ok {
				log.Fatal(err.Error())
			}
		}
	}()

	conn, err := amqp.Dial(viper.GetString("amqpbus.host"))
	throwOnError(err)
	defer conn.Close()

	ch, err := conn.Channel()
	throwOnError(err)
	defer ch.Close()

	queue, err := ch.QueueDeclare(
		viper.GetString("amqpbus.queueName"),
		true,
		false,
		false,
		false,
		nil,
	)
	throwOnError(err)

	body := e.ToMessage()
	err = ch.Publish(
		"",
		queue.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         []byte(body),
		},
	)
	throwOnError(err)

	log.WithFields(log.Fields{
		"queue": queue.Name,
		"body":  body,
	}).Info("Message published")
}

func throwOnError(e error) {
	if e != nil {
		panic(e)
	}
}
