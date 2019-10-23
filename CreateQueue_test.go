package rabbitmqLowLvlHandler

import (
    _ "bitbucket.org/accendere/ms-go-core/configDistributor"
    "fmt"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/streadway/amqp"
    "testing"
)

func Test_CreateQueue( t *testing.T ){
    Convey("Checking CreateQueue", t, func(){
        Convey("Fail 1", func() {
            var (
                err     error
                channel *amqp.Channel
            )

            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            Settings.connectMaxTry = 0

            Settings.Lock()
            queue := QUEUE_NAME + fmt.Sprintf( "%d", Settings.i )
            Settings.i++
            Settings.Unlock()

            err = CreateQueue( channel, &queue )
            So( err, ShouldNotBeNil )
        })
        Convey("Fail 2", func() {
            var (
                err     error
                channel *amqp.Channel
            )

            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            Settings.Lock()
            queue := QUEUE_NAME + fmt.Sprintf( "%d", Settings.i )
            Settings.i++
            Settings.Unlock()

            err = CreateQueue( channel, &queue )
            So( err, ShouldNotBeNil )
        })
        Convey("Normal config", func() {
            var err error

            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            conn, err := GetNewConnection()
            So( err, ShouldBeNil )
            if conn == nil{
                t.Fatal( "Connector point to nil" )
            }
            So( conn.IsClosed(), ShouldBeFalse )

            channel, err := GetNewChannel( conn )
            So( err, ShouldBeNil )
            if channel == nil{
                t.Fatal( "Channel point to nil" )
            }

            Settings.Lock()
            queue := QUEUE_NAME + fmt.Sprintf( "%d", Settings.i )
            Settings.i++
            Settings.Unlock()

            err = CreateQueue( channel, &queue )
            So( err, ShouldBeNil )

            err = DeleteQueue( channel, &queue )
            So( err, ShouldBeNil )

            err = CloseChannel( channel )
            channel = nil
            So( err, ShouldBeNil )

            err = conn.Close()
            So( err, ShouldBeNil )
        })
    })
}

