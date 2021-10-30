package handler

import (
    "fmt"
    "net/http"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "This wa definitely updated \n")
}

