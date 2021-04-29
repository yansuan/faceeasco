package faceeasco

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Cache struct {
	data sync.Map
}

func (this *Cache) Add(sn string, conn *websocket.Conn) (err error) {
	this.data.Store(sn, conn)
	return
}

func (this *Cache) Remove(sn string) (err error) {
	this.data.Delete(sn)
	return
}

func (this *Cache) Load(sn string) (result *websocket.Conn, err error) {
	if v, ok := this.data.Load(sn); ok {
		if conn, ok1 := v.(*websocket.Conn); ok1 {
			result = conn
			return
		}
	}

	return
}
