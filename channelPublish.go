package rabbitmqLowLvlHandler

import(
    "github.com/streadway/amqp"
)

func channelPublish( ch *amqp.Channel, queueName *string, message *string ) ( err error ){
    return ch.Publish(
        Settings.exchange,                     // exchange, this is the empty exchange "" by default.
        *queueName,                            // routing key
        true,                       // mandatory
        false,
        amqp.Publishing{
            DeliveryMode: amqp.Persistent,
            ContentType:  Settings.contentType,
            Body:         []byte( *message ),
        })
}