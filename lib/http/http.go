package http

import (
    "net/http"

    "github.com/gorilla/mux"

    //"github.com/PlatzhalterDE/ant.chat/dht"
    "github.com/PlatzhalterDE/ant.chat/api/v1"
    "github.com/PlatzhalterDE/ant.chat/lib/log"
)

const TAG string = "HTTP"

func Listen(listen string) error {
    log.I(TAG, "Listening on " + listen)
    return http.ListenAndServeTLS(listen, "./certs/cert.pem", "./certs/key.pem", NewRouter(routes.GetRoutes()))
}

func NewRouter(r routes.Routes) *mux.Router {
    //dht.Bootstrap()

    router := mux.NewRouter().StrictSlash(true)

    for _, route := range r {
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
