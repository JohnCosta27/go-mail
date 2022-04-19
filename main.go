package main;

import (
  "fmt"
  "net"
  "encoding/json"
  "bufio"
)

func main() {
  fmt.Println("Go Mail")

  dstream, err := net.Listen("tcp", ":8000")
  if err != nil {
    fmt.Println(err)
    return
  }
  defer dstream.Close()

  for {
    con, err := dstream.Accept()
    if err != nil {
      fmt.Println(err)
      return
    }

    go handle(con)
  }

}

type Message struct {
  From     string  `json:"from"`
  Message  string  `json:"message"`
}

// Handles each connection in its own thread.
func handle(con net.Conn) {

  for {
    data, err := bufio.NewReader(con).ReadString('\n')
    fmt.Println(string(data))
    if err != nil {
      fmt.Println(err)
      return
    }
    
    var message Message
    err = json.Unmarshal([]byte(data), &message)
    if err != nil {
      fmt.Println("Invalid json")
      return
    }

    fmt.Println(message)
    con.Write([]byte("Hi There\n"))

  }
  con.Close()
}
