package main

import (
    "net/http"
    "flag"
    "fmt"
)

var port = flag.Int("port", 3000, "Port to run on")
var out = flag.String("out", "boom", "bam")

func main() {
    flag.Parse()

    http.HandleFunc("/foo", foo)
    http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte(*out))
}
