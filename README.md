# khorasany-example-rabbitmq-golang
Example rabbitmq service 

this is simple example for initiate and run rabbitmq message broker

RabbitMQ Setup with Docker

We can setup rabbitMQ in our development environment in a couple ways; in this tutorial we'll be using docker.
In your terminal, run docker run --detach --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
This will create a docker container with a rabbitMQ instance. In your browser, navigate to localhost:15672, you should see the RabbitMQ UI; login with guest as the username and password.

Connect and Publish Messages with RabbitMQ
Now that we have our project setup and rabbitMQ running, we'll be connecting to rabbitMQ from our GoLang project. To perform the connection, we'll need the github.com/streadway/amqp library installed. To do that, run `go get -u github.com/streadway/amqp` in your terminal.

In the go_rabbit directory, create a main.go file, this is where we will put all the Golang code.
