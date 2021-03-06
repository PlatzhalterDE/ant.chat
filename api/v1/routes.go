package routes

import (
    "net/http"
    "encoding/json"

    "github.com/PlatzhalterDE/ant.chat/models"
    "github.com/PlatzhalterDE/ant.chat/lib/log"
)

const TAG = "ROUTES"

type Route struct {
    Name        string
    Method      string
    Pattern     string
    HandlerFunc http.HandlerFunc
}
type Routes []Route

func GetRoutes() Routes {
    return Routes {
        Route {
            "DHT",
            "GET",
            "/api/v1/dht",
            DhtIndexGet,
        },
    }
}

func DhtIndexGet(writer http.ResponseWriter, request *http.Request) {
    log.I(TAG, "/api/v1/dht [GET]")

    clients := models.Clients {
        models.Client { "asdf" },
        models.Client { "test" },
    }

    channels := models.Channels {
        models.Channel { "3f94c3d2-3b86-4c2a-ba6f-ca660ff24e84", clients },
        models.Channel { "3f94c3d2-3b86-4c2a-ba6f-ca660ff24e84", clients },
    }

    relays := models.Relays {
        models.Relay { "127.0.0.1:8080", channels },
    }

    json.NewEncoder(writer).Encode(relays)
}
