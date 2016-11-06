package dht

import (
    "fmt"
    "bytes"
    "strings"
    "net/http"
)

func Bootstrap() {
    fmt.Println("Getting list...")
    res, _ := http.Get("http://www.iplists.com/google.txt")

    
    buf := new(bytes.Buffer)
    buf.ReadFrom(res.Body)
    responseSplit := strings.Split(buf.String(), "\n")
    for _, element := range responseSplit {
        if len(element) > 0 && element[0] != '#' {
            fmt.Println(element)
        }
    }
}
