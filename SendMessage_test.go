package rabbitmqLowLvlHandler

import (
    "fmt"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/streadway/amqp"
    "testing"
)

func Test_SendMessage( t *testing.T ){
    Convey("Checking SendMessage", t, func(){
        var err error

        err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
        So( err, ShouldBeNil )

        Convey("Fail test's 1", func() {
            var (
                err     error
                conn    *amqp.Connection
                channel *amqp.Channel
                queue   string
            )

            _, _, err = SendMessage( conn, channel,"test message", queue, false )
            if err != nil{
                fmt.Println( err )
            }
            So( err, ShouldNotBeNil )
        })

        Convey("Fail test's 2", func() {
            var (
                err     error
                conn    *amqp.Connection
                channel *amqp.Channel
                queue   string
            )

            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            conn, err = GetNewConnection()
            So( err, ShouldBeNil )
            if conn == nil{
                t.Fatal( "Connector point to nil" )
            }
            So( conn.IsClosed(), ShouldBeFalse )

            _, _, err = SendMessage( conn, channel,"test message", queue, false )
            if err != nil{
                fmt.Println( err )
            }
            So( err, ShouldNotBeNil )

            err = conn.Close()
            So( err, ShouldBeNil )
        })

        Convey("One message send", func() {
            var (
                err     error
                conn    *amqp.Connection
                channel *amqp.Channel
            )

            conn, err = GetNewConnection()
            So( err, ShouldBeNil )
            if conn == nil{
                t.Fatal( "Connector point to nil" )
            }
            So( conn.IsClosed(), ShouldBeFalse )

            channel, err = GetNewChannel( conn )
            if channel == nil{
                t.Fatal( "Channel point to nil" )
            }
            So( err, ShouldBeNil )

            Settings.Lock()
            queue := QUEUE_NAME + fmt.Sprintf( "%d", Settings.i )
            Settings.i++
            Settings.Unlock()

            err = CreateQueue( channel, &queue )
            So( err, ShouldBeNil )

            _, _, err = SendMessage( conn, channel,"test message", queue, false )
            if err != nil{
               fmt.Println( err )
            }
            So( err, ShouldBeNil )

            err = CloseChannel( channel )
            channel = nil
            So( err, ShouldBeNil )

            err = conn.Close()
            So( err, ShouldBeNil )
        })
    })
}

