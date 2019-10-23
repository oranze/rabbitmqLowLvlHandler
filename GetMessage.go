package rabbitmqLowLvlHandler

import (
    "errors"
    uuid "github.com/satori/go.uuid"
    "github.com/sirupsen/logrus"
    "github.com/streadway/amqp"
)

func GetMessage( channel *amqp.Channel, queueNameOrig string, autoAcknowledgement bool ) ( MessageByte []byte, consumerTag string, messageO amqp.Delivery, DeliveryTag uint64, err error ){

    if Settings.connectMaxTry < 1{
        err = errors.New( ERROR_OPTIONS_NOT_INITIALIZED )
        logrus.Errorf( ERROR_OPTIONS_NOT_INITIALIZED )
        return MessageByte, consumerTag, messageO, DeliveryTag, err
    }

    if channel == nil{
        err = errors.New( ERROR_NO_CHANNEL )
        logrus.Errorf( ERROR_NO_CHANNEL )
        return MessageByte, consumerTag, messageO, DeliveryTag, err
    }

    queueNameS := QueueNameGenerate( &queueNameOrig )

    err = channel.Qos( 2, 0, true )
    if err != nil{
        return MessageByte, "", messageO, DeliveryTag, err
    }

    if err = CreateQueue( channel, &queueNameOrig ); err != nil {
        return MessageByte, "", messageO, DeliveryTag, err
    }

    var consumerTagU uuid.UUID
    consumerTagU, err = uuid.NewV4()
    if err != nil {
        return MessageByte, "", messageO, DeliveryTag, err
    }
    consumerTag = consumerTagU.String()

    Messages, err := channel.Consume(
        queueNameS,         // queue
        consumerTag,        // consumer
        autoAcknowledgement,    // auto-ack true/false. if false will be need call .AcknowledgementSend()
        false,    // exclusive
        false,     // no-local
        false,      // no-wait
        nil,          // args
    )

    // read only one message
    messageO = <- Messages
    DeliveryTag = messageO.DeliveryTag
    _ = channel.Cancel( consumerTag, false )

    MessageByte = messageO.Body

    return MessageByte, consumerTag, messageO, DeliveryTag, err
}
