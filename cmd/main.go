// +build appengine

package main

import (
    "fmt"
    "net/http"
    "github.com/austin1237/gaeDeployTest/message"
)

func init() {
    http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
    output := message.GetMessage()
    fmt.Fprint(w, output)
}
