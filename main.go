package main

import (
    "log"
)

func main() {
    log.Fatal(Listen("127.0.0.1", 8080))
}
