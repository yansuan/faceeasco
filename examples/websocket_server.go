package main

import (
	"encoding/json"
	"github.com/yansuan/faceeasco"
	"log"
	"net/http"
	"time"
)

var (
	isNotAliveCount = 0
)

var (
	client *faceeasco.Client
)

func init() {
	client = faceeasco.New()

	go func() {
		for {
			time.Sleep(1 * time.Second)
			alive()
		}
	}()
}

func alive() {
	isAlive := client.GetAliveClient("RLM-100111973")
	if isAlive == nil {
		isNotAliveCount++
		log.Println("============================================================Alive check", isAlive, isNotAliveCount)
	}
}
func main() {
	faceeasco.Debug = true
	http.HandleFunc("/device/face1", func(w http.ResponseWriter, r *http.Request) {
		client.ServeWebsocket(w, r, nil)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("============================================================Alive check", isNotAliveCount)
		http.ServeFile(w, r, "home.html")
	})
	http.HandleFunc("/face/call", func(w http.ResponseWriter, r *http.Request) {
		cmd := r.URL.Query().Get("cmd")
		sn := r.URL.Query().Get("sn")
		if cmd == "" {
			return
		}

		requestId := faceeasco.NewRequestId()
		log.Println("call requestId", requestId, cmd)

		if cmd == faceeasco.WEBSOCKET_API_SETTINGS_GET {
			form := faceeasco.WebsocketApiSettingsGetRequest{}
			form.Cmd = "to_device"
			form.From = requestId
			form.To = sn

			form.Data = faceeasco.WebsocketApiSettingsGetRequestData{}
			form.Data.Cmd = faceeasco.WEBSOCKET_API_SETTINGS_GET
			form.Data.Settings = []string{"all"}

			resp, err := client.SendWebsocketMessage(requestId, sn, form)
			if err != nil {
				log.Println(cmd, err)
				return
			}
			w.Header().Set("content-type", "text/json")
			w.Write(resp)

		} else if cmd == "client" {
			//log.Println(faceeasco.GetClientList())
			clients := client.GetClientList()
			for _, client := range clients {
				log.Println(client.GetSN(), client.GetConnTime())
			}
		} else if cmd == faceeasco.WEBSOCKET_API_SETTINGS_SET_Door {
			form := faceeasco.WebsocketApiSettingsSetDoorRequest{}
			form.Cmd = "to_device"
			form.From = requestId
			form.To = sn

			form.Data = faceeasco.WebsocketApiSettingsSetDoorRequestData{}
			form.Data.Cmd = faceeasco.WEBSOCKET_API_SETTINGS_SET_Door
			form.Data.Value = "on"

			resp, err := client.SendWebsocketMessage(requestId, sn, form)
			if err != nil {
				log.Println(cmd, err)
				return
			}
			w.Header().Set("content-type", "text/json")
			w.Write(resp)
		} else if cmd == faceeasco.WEBSOCKET_API_USER_ADD {
			for i := 0; i < 10; i++ {
				form := faceeasco.WebsocketApiUserAddRequest{}
				form.Cmd = "to_device"
				userAddRequestId := faceeasco.NewRequestId()
				form.From = userAddRequestId
				form.To = sn

				form.Data.Cmd = faceeasco.WEBSOCKET_API_USER_ADD
				form.Data.UserId = "ST-001"
				form.Data.Name = "ST-NAME"
				//form.Data.FaceTemplate = "https://up.enterdesk.com/edpic/70/0e/33/700e3312f74e378fbcc2fb3819421e73.jpg"
				photo, err1 := faceeasco.FileRead("z:\\tmp\\photo.jpg")
				if err1 != nil {
					log.Println(err1)
					return
				}
				base, err2 := faceeasco.ImageToBase64(photo)
				if err2 != nil {
					log.Println(err2)
					return
				}
				form.Data.FaceTemplate = base

				form.Data.IdValid = ""
				form.Data.Mode = 0

				resp, err := client.SendWebsocketMessage(userAddRequestId, sn, form)
				if err != nil {
					log.Println("ERROR:==================", cmd, err, resp)
					return
				}
				w.Header().Set("content-type", "text/json")
				w.Write(resp)
			}

		} else if cmd == faceeasco.WEBSOCKET_API_OnlineAuthorization {
			form := faceeasco.WebsocketApiOnlineAuthorizationRequest{}
			form.Cmd = "to_device"
			form.From = requestId
			form.To = sn

			form.Data = faceeasco.WebsocketApiOnlineAuthorizationRequestData{}
			form.Data.Cmd = faceeasco.WEBSOCKET_API_OnlineAuthorization

			resp, err := client.SendWebsocketMessage(requestId, sn, form)
			if err != nil {
				log.Println(cmd, err)
				return
			}
			respData := &faceeasco.WebsocketApiOnlineAuthorizationResponse{}
			err = json.Unmarshal(resp, respData)

			w.Header().Set("content-type", "text/html")
			src, err1 := faceeasco.Base64Unescape(respData.Data.VlFaceTemplate)
			if err != nil {
				log.Println(err1)
				return
			}
			w.Write([]byte(`<img src="` + faceeasco.Base64ImageHeader + string(src) + `"`))
			if err != nil {
				log.Println(err)
			}
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
