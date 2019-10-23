package rabbitmqLowLvlHandler

import(
    "testing"
    . "github.com/smartystreets/goconvey/convey"
)

var (
    QUEUE_NAME      string = "testQueue_"
    TEST_HOST_OK    string = "127.0.0.1"
    TEST_PORT_OK    int64  = 5672
    TEST_HOST_FAIL  string = "127.0.0.1"
    TEST_PORT_FAIL  int64  = 5999
)

func Test_Init( t *testing.T ) {
    Convey("Checking Init to fail", t, func() {
        err := Init( "", "", "", "", 0, 0)
        So( err, ShouldNotBeNil )
    })
}