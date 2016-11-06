package log

import (
    "os"
    "log"
    "strings"
)

var (
    raw  *log.Logger
    info *log.Logger
    err  *log.Logger
)

func I(tag, str string) { // Information stream
    if info == nil {
        info = log.New(os.Stdout, "info  | ", log.Ldate | log.Ltime)
    }

    info.Print(strings.ToUpper(tag) + " :: " + str)
}

func E(tag, str string) { // Error stream
    if err == nil {
        err = log.New(os.Stderr, "error | ", log.Llongfile)
    }

    err.Print(strings.ToUpper(tag) + " :: " + str)
}

func F(tag, str string) { // Fatal error stream
    if err == nil {
        err = log.New(os.Stderr, "error | ", log.Llongfile)
    }

    err.Fatal(strings.ToUpper(tag) + " :: " + str)
}

func R(str string) { // Raw stream (no tag)
    if raw == nil {
        raw = log.New(os.Stdout, "raw   | ", log.Ldate | log.Ltime)
    }

    raw.Print(str)
}
