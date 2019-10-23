package rabbitmqLowLvlHandler

import (
    "errors"
    "github.com/sirupsen/logrus"
    "github.com/streadway/amqp"
)

func SendMessage( conn *amqp.Connection, channel *amqp.Channel, message string, queueNameOrig string, AutoReconnect bool ) ( connOut *amqp.Connection, channelOut *amqp.Channel, err error ){
    queueNameS := QueueNameGenerate( &queueNameOrig )

    if ( conn == nil || conn.IsClosed() ) && AutoReconnect{
        // Connection broken. Need auto reconnect and send message
        // ToDo implement auto reconnect
        return conn, channel, err
    } else if ( conn == nil || conn.IsClosed() ) && !AutoReconnect{
        // Connection broken... no auto reconnection condition.
        // Just return with error
        logrus.Errorf( ERROR_NO_CONNECTION_DUMP, err )
        return conn, channel, errors.New( ERROR_NO_CONNECTION )
    } else{
        //// is connector is set?
        if conn == nil{
           logrus.Errorf( ERROR_NO_CONNECTION )
           return conn, channel, errors.New( ERROR_NO_CONNECTION )
        }

        // is Channel alive?
        if channel == nil{
            logrus.Errorf( ERROR_NO_CHANNEL )
            return conn, channel, errors.New( ERROR_NO_CHANNEL )
        }

        // Try to create queue (if not exist).
        // If channel broken we also get error.
        // If queue exist channel will be broken.
        if CreateQueue( channel, &queueNameOrig ) != nil {
            // Channel is broken now
            return conn, channel, err
        }

        // Connection and Channel is fine, try to send message
        err = channelPublish( channel, &queueNameS, &message )
        if err != nil{
           logrus.Errorf( ERROR_FAILED_PUBLISH_MESSAGE, err, queueNameS, message )
           return conn, channel, err
        }

        return conn, channel, err
    }
}
