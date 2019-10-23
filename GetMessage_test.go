package rabbitmqLowLvlHandler

import (
    "fmt"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/streadway/amqp"
    "testing"
    "time"
)

func Test_GetMessage( t *testing.T ) {
    Convey("Checking GetMessage", t, func() {

        Convey("Fail test 1", func() {
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

        Convey("Fail test 2", func() {
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

            Settings.connectMaxTry = 0

            _, _, err = SendMessage( conn, channel,"test message", queue, false )
            if err != nil{
                fmt.Println( err )
            }
            So( err, ShouldNotBeNil )

            err = conn.Close()
            So( err, ShouldBeNil )
        })

        Convey("Normal config", func() {
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

            var(
                MessageByte     []byte
                //consumerTag     string
                //messageO        amqp.Delivery
                DeliveryTag     uint64
                messageStr      string          = time.Now().String()
            )

            channel, err = GetNewChannel( conn )
            if channel == nil{
                t.Fatal( "Channel point to nil" )
            }
            So( err, ShouldBeNil )

            Settings.Lock()
            queue = QUEUE_NAME + fmt.Sprintf( "%d", Settings.i )
            Settings.i++
            Settings.Unlock()

            //Send message
            _, _, err = SendMessage( conn, channel, messageStr, queue, false )
            if err != nil{
                fmt.Println( err )
            }

            MessageByte, _, _, DeliveryTag, err = GetMessage( channel, queue, false )
            if err != nil{
                fmt.Println( err )
            }
            So( err, ShouldBeNil )

            fmt.Println( "DeliveryTag: "+fmt.Sprintf( "%v", DeliveryTag ) )

            fmt.Println( fmt.Sprintf( "Send to RabbitMQ: %s\nReceived: %s", messageStr, string( MessageByte ) ) )
            So( string( MessageByte ), ShouldEqual, messageStr )

            err = AcknowledgementSend( channel, DeliveryTag )
            So( err, ShouldBeNil )

            err = CloseChannel( channel )
            So( err, ShouldBeNil )

            err = conn.Close()
            So( err, ShouldBeNil )
        })
    })
}