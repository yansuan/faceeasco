package main

import (
	"github.com/yansuan/faceeasco"
	"log"
	"net/http"
)

func PingCall(message []byte) (err error) {
	log.Println("PingCall", string(message))
	return
}
func main() {
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		faceeasco.RegisterWebsocketFunc(faceeasco.WEBSOCKET_API_PING, PingCall)
		faceeasco.ServeWebsocket(w, r, nil)
	})
	err := http.ListenAndServe(":10001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
