package rabbitmqLowLvlHandler

import (
    "errors"
    "github.com/sirupsen/logrus"
    "github.com/streadway/amqp"
)

func AcknowledgementSend( channel *amqp.Channel, DeliveryTag uint64 ) ( err error ){
    if channel == nil{
        err = errors.New( ERROR_NO_CHANNEL )
        logrus.Errorf( ERROR_NO_CHANNEL )
        return err
    }

    err = channel.Ack( DeliveryTag, false )
    if err != nil{
        logrus.Errorf( ERROR_ACK_RESPONSE, err )
    }

    return err
}