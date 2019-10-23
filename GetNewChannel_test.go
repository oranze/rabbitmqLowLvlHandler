package rabbitmqLowLvlHandler

import (
    _ "bitbucket.org/accendere/ms-go-core/configDistributor"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/streadway/amqp"
    "testing"
)

func Test_GetNewChannel( t *testing.T ){
    Convey("Checking GetNewChannel", t, func(){
        Convey("Fail 1", func() {
            var (
                err     error
                conn    *amqp.Connection
            )
            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            _, err = GetNewChannel( conn )
            So( err, ShouldNotBeNil )
        })

        Convey("Fail 2", func() {
            var err error

            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            conn, err := GetNewConnection()
            So( err, ShouldBeNil )
            if conn == nil{
                t.Fatal( "Connector point to nil" )
            }
            So( conn.IsClosed(), ShouldBeFalse )

            Settings.connectMaxTry = 0

            _, err = GetNewChannel( conn )
            So( err, ShouldNotBeNil )

            err = conn.Close()
            So( err, ShouldBeNil )
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

            // Close channel test
            err = CloseChannel( channel )
            So( err, ShouldBeNil )

            err = conn.Close()
            So( err, ShouldBeNil )
        })
    })
}

