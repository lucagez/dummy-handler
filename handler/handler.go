package handler

import (
    "bufio"
    "bytes"
    "fmt"
    "net/http"
    "os/exec"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
    cmd := exec.Command("deno", "eval", "console.log('Ciaooo')")

    var buf bytes.Buffer
    out := bufio.NewWriter(&buf)
    cmd.Stdout = out
    cmd.Stderr = out
    cmd.Run()

    fmt.Fprintf(w, buf.String())
}

