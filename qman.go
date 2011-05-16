package main

import (
  "http"
  "io"
  "os"
  "flag"
  "fmt"
)

var addr = flag.String("addr", ":1718", "http service address") // Q=17, R=18

func enqueue(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "success! \n")
}

func main() {
  flag.Parse()
  port := os.Getenv("PORT")
  http.HandleFunc("/enqueue", enqueue)

  fmt.Println("listening on", *addr)
  http.ListenAndServe(*addr, nil)
}
