package rabbitmqLowLvlHandler

import (
    _ "bitbucket.org/accendere/ms-go-core/configDistributor"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/streadway/amqp"
    "testing"
)

func Test_CloseConnection( t *testing.T ){
    Convey("Checking CloseConnection", t, func(){
        Convey("Fail 1", func() {
            var (
                err     error
                conn *amqp.Connection
            )

            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            Settings.connectMaxTry = 0

            err = CloseConnection( conn )
            So( err, ShouldNotBeNil )
        })
        Convey("Normal test", func() {
            var (
                err     error
                conn *amqp.Connection
            )

            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            conn, err = GetNewConnection()
            So( err, ShouldBeNil )
            if conn == nil{
                t.Fatal( "Connector point to nil" )
            }
            So( conn.IsClosed(), ShouldBeFalse )

            _ = CloseConnection( conn )

            err = CloseConnection( conn )
            So( err, ShouldNotBeNil )
        })
    })
}

