package rabbitmqLowLvlHandler

import "os"

// DetectIdentity returns the host identity. This is usually the host name, but can be overriden by setting the environment
// variable IDENTITY. This value is used to separate individual developer's RabbitMQ message queues, while still allowing
// us to explicitly connect to other devs' queues.
func getIdentity() string {
    if val, ok := os.LookupEnv( "IDENTITY" ); ok {
        return val
    }
    hostName, err := os.Hostname()
    if err != nil {
        return "unknown"
    }
    return hostName
}
