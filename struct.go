package rabbitmqLowLvlHandler

import (
    "sync"
)

type options struct {
    sync.RWMutex
    identity        string
    env             string
    host            string
    exchange        string
    contentType     string
    connectMaxTry   int64
    i               int64
}