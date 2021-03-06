package main
 
import (
        "github.com/gorilla/websocket"
        "net/http"
        "fmt"
)
 
var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
 
func print_binary(s []byte) {
  fmt.Printf("Received b:");
  for n := 0;n < len(s);n++ {
    fmt.Printf("%d,",s[n]);
  }
  fmt.Printf("\n");
}
 
func echoHandler(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        //log.Println(err)
        return
    }
 
    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            return
        }
 
        print_binary(p)
 
        err = conn.WriteMessage(messageType, p);
        if  err != nil {
            return
        }
    }
}
 
func main() {
  http.HandleFunc("/echo", echoHandler)
  http.Handle("/", http.FileServer(http.Dir(".")))
  err := http.ListenAndServe(":8085", nil)
  if err != nil {
    panic("Error: " + err.Error())
  }
}
