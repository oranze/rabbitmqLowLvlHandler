package rabbitmqLowLvlHandler

import (
    "errors"
    "github.com/sirupsen/logrus"
    "github.com/streadway/amqp"
)

// DeleteQueue remove a queue with the name queueName on the AMQP server and binds it to the exchange rabbitKey.exchange.
// The binding uses the queue name as the routing key.
func DeleteQueue ( channel *amqp.Channel, queueName *string ) ( err error ){
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

    queueNameS := QueueNameGenerate( queueName )

    _,err = channel.QueueDelete( queueNameS, false, false, false )
    return err
}
