package rabbitmqLowLvlHandler

import (
    _ "bitbucket.org/accendere/ms-go-core/configDistributor"
    . "github.com/smartystreets/goconvey/convey"
    "testing"
)

func Test_GetNewConnection( t *testing.T ){
    Convey("Checking connector", t, func() {
        Convey("Wrong config", func() {
            var err error

            err = Init( "unit", TEST_HOST_FAIL, "test", "test", TEST_PORT_FAIL, 5 )
            So( err, ShouldBeNil )

            _, err = GetNewConnection()
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
            err = conn.Close()
            So( err, ShouldBeNil )
        })
    })
}

