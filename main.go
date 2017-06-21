package main

import (
  "fmt"
  "github.com/thedch/md5_server/sums"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello there! The md5 hash of %s is %x", r.URL.Path[1:],
      calculate_sum.GetMD5Hash(r.URL.Path[1:]))
}

func main() {
  http.HandleFunc("/", handler)
  http.ListenAndServe(":8080", nil)
}
