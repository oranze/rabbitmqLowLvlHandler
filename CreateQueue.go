package rabbitmqLowLvlHandler

import (
    "errors"
    "github.com/sirupsen/logrus"
    "github.com/streadway/amqp"
)

// createQueue declares a queue with the name queueName on the AMQP server and binds it to the exchange rabbitKey.exchange.
// The binding uses the queue name as the routing key.
//
// This function requires rabbitKey to at least be read-locked.
func CreateQueue( channel *amqp.Channel, queueName *string ) ( err error ){
    if Settings.connectMaxTry < 1{
        err = errors.New( ERROR_OPTIONS_NOT_INITIALIZED )
        logrus.Errorf( ERROR_OPTIONS_NOT_INITIALIZED )
        return err
    }

    if channel == nil{
        err = errors.New( ERROR_NO_CHANNEL )
        logrus.Errorf( ERROR_NO_CHANNEL )
        return err
    }

    _, err = channel.QueueDeclare(
        QueueNameGenerate( queueName ),
        true,        // durable
        false,     // delete when unused
        false,      // exclusive
        false,        // no-wait
        nil,            // arguments
    )

    return err
}
