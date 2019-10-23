package rabbitmqLowLvlHandler

import (
    "errors"
    "fmt"
    "github.com/sirupsen/logrus"
    "strings"
)

var (
    Settings options
)

/*
    Init() - initialization function
    After Init() we can call GetNewConnect
*/
func Init( environment string, host string, user string, password string, port int64, initConnectMaxTry int64 ) ( err error ){
    environment = strings.TrimSpace( environment )
    host = strings.ToLower( host )
    host = strings.TrimSpace( host )
    host = strings.TrimPrefix( host, "http://" )
    host = strings.TrimPrefix( host, "https://" )
    host = strings.TrimPrefix( host, "amqp://" )
    host = strings.TrimSuffix( host, "/" )
    host = strings.TrimSpace( host )

    user = strings.TrimSpace( user )
    password = strings.TrimSpace( password )

    if len( environment ) < 1 || len( host ) < 2 || len( user ) < 1 || len( password ) < 1 || port < 1 || initConnectMaxTry < 1{
        err = errors.New( fmt.Sprintf( ERROR_WRONG_INIT_DATA, environment, host, user, port, initConnectMaxTry ) )
        logrus.Errorf( ERROR_WRONG_INIT_DATA, environment, host, user, port, initConnectMaxTry )
        return err
    }

    // all fine
    // start initialization
    Settings.identity       = getIdentity()
    Settings.env            = strings.ToLower( environment )
    Settings.host           = "amqp://"+ user +":"+ password +"@"+ host +":"+ fmt.Sprintf( "%d", port )
    Settings.exchange       = ""
    Settings.contentType    = "text/plain"
    Settings.connectMaxTry  = initConnectMaxTry

    return err
}