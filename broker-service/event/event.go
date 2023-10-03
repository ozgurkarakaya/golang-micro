package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topic", //name of the exchange
		"topic",      //type
		true,         //isDurable
		false,        //is auto-delete enabled
		false,        //is it internal, in our case it's between microservices
		false,        //no-wait?
		nil,          //arguments
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    //name --> randomly generated in our case
		false, //durable?
		false, //delete when unused?
		true,  //exclusive? --> don't share
		false, //no wait
		nil,   //arguments?
	)
}
