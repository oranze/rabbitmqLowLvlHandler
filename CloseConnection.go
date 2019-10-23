package rabbitmqLowLvlHandler

import (
    "errors"
    "fmt"
    "github.com/sirupsen/logrus"
    "github.com/streadway/amqp"
)

/*
    Closing open connection
*/
func CloseConnection( conn *amqp.Connection ) ( err error ){
    if conn == nil || conn.IsClosed(){
        err = errors.New( ERROR_NO_CONNECTION )
        logrus.Warnf( ERROR_NO_CONNECTION )
        return err
    }

    err = conn.Close()
    if err != nil{
        logrus.Errorf( ERROR_CONNECTION_CLOSE, err )
        err = errors.New( fmt.Sprintf( ERROR_CONNECTION_CLOSE, err ) )
        return err
    }

    return err
}
