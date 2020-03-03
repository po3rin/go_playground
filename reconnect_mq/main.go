package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type queue struct {
	url  string
	name string

	errorChannel chan *amqp.Error
	connection   *amqp.Connection
	channel      *amqp.Channel
	closed       bool

	consumers []messageConsumer
}

type messageConsumer func(string)

func (q *queue) Send(message string) error {
	var retryCounter int
	for {
		err := q.channel.Publish(
			"",     // exchange
			q.name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
		if err != nil {
			retryCounter++
			if retryCounter > 10 {
				return errors.New("retry ")
			}
			time.Sleep(1 * time.Second)
			continue
		}
		return nil
	}
}

func (q *queue) Consume(consumer messageConsumer) {
	log.Println("Registering consumer...")
	deliveries, err := q.registerQueueConsumer()
	log.Println("Consumer registered! Processing messages...")
	q.executeMessageConsumer(err, consumer, deliveries, false)
}

func (q *queue) Close() {
	log.Println("Closing connection")
	q.closed = true
	q.channel.Close()
	q.connection.Close()
}

func (q *queue) reconnector() {
	for {
		err := <-q.errorChannel
		if !q.closed {
			logError("Reconnecting after connection closed", err)

			q.connect()
			q.recoverConsumers()
		}
	}
}

func (q *queue) connect() {
	for {
		log.Printf("Connecting to rabbitmq on %s\n", q.url)
		conn, err := amqp.Dial(q.url)
		if err == nil {
			q.connection = conn
			q.errorChannel = make(chan *amqp.Error)
			q.connection.NotifyClose(q.errorChannel)

			log.Println("Connection established!")

			q.openChannel()
			q.declareQueue()

			return
		}

		logError("Connection to rabbitmq failed. Retrying in 1 sec... ", err)
		time.Sleep(1000 * time.Millisecond)
	}
}

func (q *queue) declareQueue() {
	_, err := q.channel.QueueDeclare(
		q.name, // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	logError("Queue declaration failed", err)
}

func (q *queue) openChannel() {
	channel, err := q.connection.Channel()
	logError("Opening channel failed", err)
	q.channel = channel
}

func (q *queue) registerQueueConsumer() (<-chan amqp.Delivery, error) {
	msgs, err := q.channel.Consume(
		q.name, // queue
		"",     // messageConsumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	logError("Consuming messages from queue failed", err)
	return msgs, err
}

func (q *queue) executeMessageConsumer(err error, consumer messageConsumer, deliveries <-chan amqp.Delivery, isRecovery bool) {
	if err == nil {
		if !isRecovery {
			q.consumers = append(q.consumers, consumer)
		}
		go func() {
			for delivery := range deliveries {
				consumer(string(delivery.Body[:]))
			}
		}()
	}
}

func (q *queue) recoverConsumers() {
	for i := range q.consumers {
		var consumer = q.consumers[i]

		log.Println("Recovering consumer...")
		msgs, err := q.registerQueueConsumer()
		log.Println("Consumer recovered! Continuing message processing...")
		q.executeMessageConsumer(err, consumer, msgs, true)
	}
}

func logError(message string, err error) {
	if err != nil {
		log.Printf("%s: %s", message, err)
	}
}

func NewQueue(url string, qName string) *queue {
	q := new(queue)
	q.url = url
	q.name = qName
	q.consumers = make([]messageConsumer, 0)

	q.connect()
	go q.reconnector()

	return q
}

func main() {
	queue := NewQueue("amqp://guest:guest@localhost:5672/", "hello")
	defer queue.Close()

	queue.Consume(func(i string) {
		log.Printf("Received message with second consumer: %s", i)
	})

	queue.Consume(func(i string) {
		log.Printf("Received message with first consumer: %s", i)
	})

	for i := 0; i < 100; i++ {
		log.Println("Sending message...")
		err := queue.Send(fmt.Sprint("dupa", i))
		if err != nil {
			log.Println("shutdown...")
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}
