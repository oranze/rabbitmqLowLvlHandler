package rabbitmqLowLvlHandler

import (
    _ "bitbucket.org/accendere/ms-go-core/configDistributor"
    . "github.com/smartystreets/goconvey/convey"
    "github.com/streadway/amqp"
    "testing"
)

func Test_CloseChannel( t *testing.T ) {
    Convey("Checking CloseChannel", t, func() {
        Convey("Fail 1", func() {
            var (
                err     error
                channel *amqp.Channel
            )

            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            err = CloseChannel( channel )
            So( err, ShouldNotBeNil )
        })

        Convey("Fail 2", func() {
            var (
                err     error
                channel *amqp.Channel
            )

            err = Init( "unit", TEST_HOST_OK, "test", "test", TEST_PORT_OK, 5 )
            So( err, ShouldBeNil )

            Settings.connectMaxTry = 0

            err = CloseChannel( channel )
            So( err, ShouldNotBeNil )
        })
    })
}