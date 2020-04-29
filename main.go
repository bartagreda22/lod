package main

import (
  "fmt"
  "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>Kubernetes is awesome!</h1>")
}

func main() {
  http.HandleFunc("/", homeHandler)
  http.ListenAndServe(":8000", nil)
}
