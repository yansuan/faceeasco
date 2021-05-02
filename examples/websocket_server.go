package main

import (
	"github.com/yansuan/faceeasco"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/face", func(w http.ResponseWriter, r *http.Request) {
		faceeasco.ServeWebsocket(w, r, nil)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "home.html")
	})
	http.HandleFunc("/face/call", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Query().Get("cmd")
		sn := r.URL.Query().Get("sn")
		if cmd == "" {
			return
		}

		requestId := faceeasco.NewUUID()
		log.Println("call requestId", requestId, cmd)

		if cmd == faceeasco.WEBSOCKET_API_SETTINGS_GET {
			form := faceeasco.WebsocketApiSettingsGetRequest{}
			form.Cmd = "to_device"
			form.From = requestId
			form.To = sn

			form.Data = faceeasco.WebsocketApiSettingsGetRequestData{}
			form.Data.Cmd = faceeasco.WEBSOCKET_API_SETTINGS_GET
			form.Data.Settings = []string{"all"}

			resp, err := faceeasco.SendWebsocketMessage(requestId, sn, form)
			if err != nil {
				log.Println(cmd, err)
				return
			}
			w.Header().Set("content-type", "text/json")
			w.Write(resp)

		} else if cmd == "client" {
			log.Println(faceeasco.GetClientList())
		}
	})

	err := http.ListenAndServe(":10001", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
