package handler

import (
    "bufio"
    "bytes"
    "fmt"
    "net/http"
    "os"
    "os/exec"
)

func init() {
    os.Setenv("DENO_INSTALL", ".")

    cmd := exec.Command("sh", "install.sh")
    cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
    cmd.Run()
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
    cmd := exec.Command("bin/deno", "eval", "console.log('Ciaooo')")

    var buf bytes.Buffer
    out := bufio.NewWriter(&buf)
    cmd.Stdout = out
    cmd.Stderr = out
    err := cmd.Run()

    if err != nil {
        fmt.Fprintf(w, err.Error())
        fmt.Fprintf(w, "\n")
    }

    fmt.Fprintf(w, buf.String())
}

