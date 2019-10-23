package rabbitmqLowLvlHandler

const(
    ERROR_CONNECTION_FAILED                                     = "Rabbit connection failed, err: %v"
    ERROR_NO_CONNECTION_DUMP                                    = "Rabbit no connection, err: %v"
    ERROR_NO_CONNECTION                                         = "Rabbit no connection"
    ERROR_NO_CHANNEL                                            = "Rabbit no channel"
    ERROR_WRONT_HOSTNAME                                        = "Rabbit. Wrong host name `%s`. Can't connect."
    ERROR_CONNECTION_CLOSE                                      = "Rabbit connection close failed, err: %v"
    ERROR_ACK_RESPONSE                                          = "Acknowledgement send error: %v"
    ERROR_CHANNEL_CLOSE_FAILED                                  = "Rabbit channel close failed, err: %v"
    ERROR_CHANNEL_OPEN_FAILED                                   = "Rabbit channel open failed, err: %v"
    ERROR_FAILED_PUBLISH_MESSAGE                                = "Rabbit Failed to publish a message error: %v,\nqueue: %s,\nmessage: %s"
    ERROR_WRONG_INIT_DATA                                       = "Rabbit. Wrong Init() params:\nenvironment: %s\nhost: %s\nuser: %s\npassword: ******\nport: %d\ninitConnectMaxTry: %d"
    ERROR_OPTIONS_NOT_INITIALIZED                               = "Rabbit. Options not initialized. Please call rabbitmqLowLvlHandler.Init() first"
)
