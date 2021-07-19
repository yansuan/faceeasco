package queue

import (
	"log"
	"sync"
	"time"
)

const TIMEOUT = 10

var q *Queue

func init() {
	q = newQueue()
	go q.clean()
}

type Message struct {
	RequestId string
	Time      int64
	SN        string
	Body      []byte
}

type Client struct {
	RequestId string
	Message   chan []byte
}

type Queue struct {
	sync.RWMutex
	data    map[string]*Message //
	clients map[*Client]int64   //
}

func newQueue() *Queue {
	o := &Queue{}
	o.data = make(map[string]*Message)
	o.clients = make(map[*Client]int64)

	return o
}

func (this *Queue) cleanData() {
	this.Lock()
	defer this.Unlock()
	now := time.Now().Unix()

	//删除过期消息
	for requestId, msg := range this.data {

		if msg.Time+TIMEOUT > now {
			delete(this.data, requestId)
		}
	}

	//删除过期client
	for client, ts := range this.clients {
		if ts+TIMEOUT > now {
			delete(this.clients, client)
		}
	}
}

func (this *Queue) clean() {
	for {
		time.Sleep(TIMEOUT * time.Second)
		this.cleanData()
	}
}

func Connect(requestId string) *Client {
	q.Lock()
	defer q.Unlock()
	c := &Client{RequestId: requestId}
	c.Message = make(chan []byte)
	q.clients[c] = time.Now().Unix()
	return c
}

func Push(msg *Message) {
	q.Lock()
	defer q.Unlock()
	//q.data[msg.RequestId] = msg

	for client := range q.clients {
		log.Println("Push", msg.RequestId, client.RequestId)
		if client.RequestId == msg.RequestId {
			client.Message <- msg.Body
		}
	}
}

//func Pull(requestId string) *Message {
//	q.RLock()
//	defer q.RUnlock()
//
//	now := time.Now().Unix()
//	for id, msg := range q.data {
//		if msg.Time+TIMEOUT <= now && requestId == id {
//			return msg
//		}
//	}
//
//	return nil
//}

//func Remove(requestId string) {
//	q.Lock()
//	defer q.Unlock()
//	delete(q.data, requestId)
//}
