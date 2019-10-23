package rabbitmqLowLvlHandler

import (
    "errors"
    "github.com/sirupsen/logrus"
    "github.com/streadway/amqp"
)

/*
    GetNewConnection() - try to establish connection to RabbitMQ server
*/
func GetNewConnection() ( conn *amqp.Connection, err error ){
    if len( Settings.host )< 1{
        err = errors.New( ERROR_WRONT_HOSTNAME )
        logrus.Errorf( ERROR_WRONT_HOSTNAME, Settings.host )
        return conn, err
    }

    var i int64 = 0
    for i < Settings.connectMaxTry{
        conn, err = amqp.Dial( Settings.host )
        if err != nil{
            logrus.Errorf( ERROR_CONNECTION_FAILED, err )
            i++
        } else{
            logrus.Infof( INFO_CONNECTION_ESTABLISHED, Settings.host )
            break
        }
    }

    return conn, err
}