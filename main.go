package main

import (
    "net/http"
    "github.com/lucagez/dummy-handler/handler"
)

func main() {
    http.HandleFunc("/hello", handler.HelloHandler)

    http.ListenAndServe(":8090", nil)
}
