package main

import (
    "strconv"

    "github.com/PlatzhalterDE/ant.chat/lib/cert"
    "github.com/PlatzhalterDE/ant.chat/lib/log"
    "github.com/PlatzhalterDE/ant.chat/lib/http"
)

var (
    host string = "127.0.0.1"
    port int = 8080
)

const TAG string = "MAIN"

func main() {
    log.R("             _         _           _")
    log.R("  __ _ _ __ | |_   ___| |__   __ _| |_")
    log.R(" / _` | '_ \\| __| / __| '_ \\ / _` | __|")
    log.R("| (_| | | | | |_ | (__| | | | (_| | |_")
    log.R(" \\__,_|_| |_|\\__(_)___|_| |_|\\__,_|\\__|")
    log.R("")

    listen := host + ":" + strconv.Itoa(port)

    if(!cert.FilesExist()) {
        log.I(TAG, "No certificates exist! Generating...")
        err := cert.WriteToFile(cert.GenerateKeyPair(listen))
        if err != nil {
            log.F(TAG, err.Error())
        }
        log.I(TAG, "Certificates generated!")
    }

    log.E(TAG, http.Listen(listen).Error())
}
