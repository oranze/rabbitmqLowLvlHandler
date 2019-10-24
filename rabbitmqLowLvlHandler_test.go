package rabbitmqLowLvlHandler

import (
    "bitbucket.org/accendere/ms-go-core/logger"
    "fmt"
    "github.com/streadway/amqp"
    "os"
    "testing"
    "time"
)

func TestMain( m *testing.M ){
    // Clean old queue's
    var (
        err     error
        i       int64   = 0
        conn    *amqp.Connection
        channel *amqp.Channel
        queue   string
    )

    err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
    if err != nil{
        fmt.Println( "Can't init handler" )
        os.Exit( 999 )
    }

    conn, err = GetNewConnection()
    if err == nil{

        channel, err = GetNewChannel( conn )
        if err == nil{
            Settings.i = 20
            for i <= Settings.i{
                queue = QUEUE_NAME + fmt.Sprintf( "%d", i )
                str :=  "Try to remove QUEUE: "+queue+" ..."
                err = DeleteQueue( channel, &queue )
                if err == nil{
                    str = str +" OK"
                    fmt.Println( str )
                } else{
                    str = str +" Fail. error: %v"
                    fmt.Println( fmt.Sprintf( str, err ) )
                }
                i++
            }
            Settings.i = 0
        } else{
            fmt.Println( fmt.Sprintf( "Can't remove QUEUE'S. Channel problem. Err: %v", err ) )
        }


    } else{
        fmt.Println( fmt.Sprintf( "Can't remove QUEUE'S. Connector problem. Err: %v", err ) )
    }

    // start testing
    exitCode := m.Run()
    if exitCode == 0 {
        err := logger.DeleteLogFolder()
        if err != nil{
            fmt.Println( fmt.Sprintf( "ERROR: %v", err ) )
            exitCode = 999
        }
    }

    if exitCode != 0{
        fmt.Println( "!!!! The log file saved !!!!" )
    }

    // Queue's cleaning
    if exitCode == 0{
        var (
            err     error
            i       int64   = 0
            conn    *amqp.Connection
            channel *amqp.Channel
            queue   string
        )

        err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
        if err != nil{
            // Prevent loop refresh in Convey
            time.Sleep( 2 * time.Second )
            os.Exit( exitCode )
        }

        conn, err = GetNewConnection()
        if err == nil{

            channel, err = GetNewChannel( conn )
            if err == nil{
                for i <= Settings.i{
                    queue = QUEUE_NAME + fmt.Sprintf( "%d", i )
                    str :=  "Try to remove QUEUE: "+queue+" ..."
                    err = DeleteQueue( channel, &queue )
                    if err == nil{
                        str = str +" OK"
                        fmt.Println( str )
                    } else{
                        str = str +" Fail. error: %v"
                        fmt.Println( fmt.Sprintf( str, err ) )
                    }
                    i++
                }
            } else{
                fmt.Println( fmt.Sprintf( "Can't remove QUEUE'S. Channel problem. Err: %v", err ) )
            }


        } else{
            fmt.Println( fmt.Sprintf( "Can't remove QUEUE'S. Connector problem. Err: %v", err ) )
        }
    }

    // Prevent loop refresh in Convey
    time.Sleep( 2 * time.Second )
    os.Exit( exitCode )
}