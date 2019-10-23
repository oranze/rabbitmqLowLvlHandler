package rabbitmqLowLvlHandler

import (
    "errors"
    "github.com/sirupsen/logrus"
    "github.com/streadway/amqp"
)

/*
    Try to close current open channel
*/

func CloseChannel ( channel *amqp.Channel ) ( err error ){
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

    err = channel.Close()
    if err != nil{
        logrus.Errorf( ERROR_CHANNEL_CLOSE_FAILED, err )
    } else{
        logrus.Infof( INFO_CHANNEL_CLOSED )
    }
    return err
}