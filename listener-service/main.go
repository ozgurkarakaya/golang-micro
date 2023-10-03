package main

import (
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"math"
	"os"
	"time"
)

func main() {

	//try to connect rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	log.Println("Connected to RabbitMQ!")

	//start listening for messages

	//create consumer

	//watch the queue and consume events

}

func connect() (*amqp.Connection, error) {

	var counts int64
	var backoff = 1 * time.Second
	var connection *amqp.Connection

	//wait until rabbitmq is ready
	for {
		c, err := amqp.Dial("amqp://guest:guest@localhost")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			connection = c
			break
		}

		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		backoff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println(fmt.Sprintf("backing of for %d sec", backoff/10^9))
		time.Sleep(backoff)
		continue
	}

	return connection, nil

}
