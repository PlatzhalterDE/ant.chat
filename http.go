package main

import (
    "strconv"
    "net/http"

    "github.com/gorilla/mux"

    //"github.com/PlatzhalterDE/ant.chat/dht"
    "github.com/PlatzhalterDE/ant.chat/lib"
    "github.com/PlatzhalterDE/ant.chat/api/v1"
)

func Listen(host string, port int) error {
    listenString := host + ":" + strconv.Itoa(port)

    err := lib.WriteToFileIfNotExists(lib.GenerateKeyPair(host))
    if err != nil {
        return err
    }

    return http.ListenAndServeTLS(listenString, "./certs/cert.pem", "./certs/key.pem", NewRouter(api.GetRoutes()))
}

func NewRouter(routes api.Routes) *mux.Router {
    //dht.Bootstrap()

    router := mux.NewRouter().StrictSlash(true)

    for _, route := range routes {
        var handler http.Handler
        handler = route.HandlerFunc

        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}
