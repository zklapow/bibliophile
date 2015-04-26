package main

import (
    "flag"
    "github.com/zklapow/bibliophile/api"
)

var (
    redisServer = flag.String("redis", ":6379", "Address of redis server.")
    port = flag.Int("port", 8080, "Port to start API on.")
)

func main() {
    api.Start(*port, *redisServer)
}


