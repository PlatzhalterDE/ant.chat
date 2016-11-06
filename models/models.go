package models

type Channel    struct {
    Name        string      `json:"name"`
    Clients     Clients     `json:"clients"`
}

type Client struct {
    Publickey   string      `json:"publickey"`
}

type Relay      struct {
    Host        string      `json:"host"`
    Channels    Channels    `json:"channels"`
}

type Message    struct {
    Message     string      `json:"message"`
}

type Channels   []Channel
type Clients    []Client
type Relays     []Relay
