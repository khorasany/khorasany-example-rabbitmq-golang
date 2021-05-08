package main

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

// Get the connection string from the environment variable
var amqpUrl = os.Getenv("AMQP_URL")

func main() {
	publishingAmqp()
	createQueue()
	consumeMessage()
}

func connectRabbitmqServer() *amqp.Connection {
	//If it doesn't exist, use the default connection string.
	if amqpUrl == "" {
		        //Don't do this in production, this is for testing purposes only.
		amqpUrl = "amqp://guest:guest@localhost:5672"
	}

	// Connect to the rabbitMQ instance
	connection,err := amqp.Dial(amqpUrl)
	if err != nil {
		log.Fatalf("error while lunch amqp server: %v",err)
	}

	return connection
}

func makeChannelConnection() *amqp.Channel {
	connection := connectRabbitmqServer()
	// Create a channel from the connection. We'll use channels to access the data in the queue rather than the connection itself.
	channel,err := connection.Channel()
	if err != nil {
		log.Fatalf("error while make channel: %v",err)
	}

    // We create an exchange that will bind to the queue to send and receive messages
	err = channel.ExchangeDeclare("events","topic",true,false,false,false,nil)
	if err != nil {
		log.Fatal(err)
	}
	return channel
}

func publishingAmqp() {
	// We create a message to be sent to the queue. 
    // It has to be an instance of the aqmp publishing struct
	message := amqp.Publishing{
		Body : []byte("hello world"),
	}

	channel := makeChannelConnection()

	// We publish the message to the exahange we created earlier
	err := channel.Publish("events","random-key",false,false,message)
	if err != nil {
		log.Fatalf("error while publish channel: %v",err)
	}
}

func createQueue() {
	channel := makeChannelConnection()

	// We create a queue named Test
	_,err := channel.QueueDeclare("test",true,false,false,false,nil)
	if err != nil {
		log.Fatalf("error while declare queue: %v",err)
	}

	// We bind the queue to the exchange to send and receive data from the queue
	err = channel.QueueBind("test","#","events",false,nil)
	if err != nil {
		log.Fatalf("error while bind queue: %v",err)
	}
}

func consumeMessage() {
	connection := connectRabbitmqServer()
	channel := makeChannelConnection()

	// We consume data in the queue named test using the channel we created in go.
	messages,err := channel.Consume("test","",false,false,false,false,nil)
	if err != nil {
		log.Fatalf("error while consuming delivery channel: %v",err)
	}

	// We loop through the messages in the queue and print them to the console.
    // The msgs will be a go channel, not an amqp channel
	for message := range messages{
		//print the message to the console
		fmt.Println("message recieved: " + string(message.Body))
		// Acknowledge that we have received the message so it can be removed from the queue
		message.Ack(false)
	}

	// We close the connection after the operation has completed.
	defer connection.Close()
}