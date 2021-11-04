package main

import (
    "github.com/lucagez/dummy-handler/handler"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/hello", handler.HelloHandler)

    port := os.Getenv("PORT")

    http.ListenAndServe(":" + port, nil)
}
