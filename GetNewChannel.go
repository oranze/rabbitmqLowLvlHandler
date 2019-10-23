package rabbitmqLowLvlHandler

import (
    "errors"
    "github.com/sirupsen/logrus"
    "github.com/streadway/amqp"
)

func GetNewChannel( conn *amqp.Connection ) ( channel *amqp.Channel, err error ){
    if Settings.connectMaxTry < 1{
        err = errors.New( ERROR_NO_CONNECTION )
        logrus.Errorf( ERROR_NO_CONNECTION )
        return channel, err
    }

    if conn == nil || conn.IsClosed(){
        err = errors.New( ERROR_NO_CONNECTION )
        logrus.Errorf( ERROR_NO_CONNECTION )
        return channel, err
    }

    channel, err = conn.Channel()
    if err != nil || channel == nil{
        logrus.Errorf( ERROR_CHANNEL_OPEN_FAILED, err )
    } else{
        logrus.Infof( INFO_CHANNEL_OPEN )
    }

    return channel, err
}