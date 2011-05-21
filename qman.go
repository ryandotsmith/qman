package main

import (
  "http"
  "json"
  "os"
  "io/ioutil"
  "flag"
  "fmt"
)

type Queue struct { QueueId string }

var (
  sickleUrl = "http://sickle:jHTbTrJtBry8WZIl@localhost:9292"
  addr = flag.String("addr", ":1718", "http service address")
)

func sickle_lookup(u ,p ,q_name string, q *Queue) {
  url := sickleUrl + "/lookup?username=" + u + "&password=" + p + "&queue_name=" + q_name
  client := new(http.Client)
  resp, _, _ := client.Get(url)
  body, _ := ioutil.ReadAll(resp.Body)
  if resp.StatusCode == 404 {
    fmt.Println("could not find queue")
  }
  json.Unmarshal(body,q)
  return
}

func write_enqueue_stats(q *Queue, payload []byte) os.Error {
  fmt.Println("writing stats")
  return ioutil.WriteFile("stats.txt", payload , 0600)
}

func write_enqueue_payload(q *Queue, payload []byte) os.Error {
  fmt.Println("writing to queue")
  filename := q.QueueId+".txt"
  return ioutil.WriteFile(filename, payload, 0600)
}

func (q *Queue) enqueue(payload []byte) {
  sickle_lookup("user","pass","default_queue",q)
  go write_enqueue_payload(q, payload)
  go write_enqueue_stats(q, payload)
  fmt.Println("200")
  return
}

func echo(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "success")
}

func main() {
  flag.Parse()
  http.HandleFunc("/enqueue", echo)
  fmt.Println("listening on", *addr)
  http.ListenAndServe(*addr, nil)
  //payload := []byte("Class.method")
  //queue := new(Queue)
  //queue.enqueue(payload)
}
