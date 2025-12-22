package ws

// type WebSocket struct {
// 	ws websocket.Upgrader
// }

// func NewWebSocket(ws websocket.Upgrader) WebSocket {
// 	return WebSocket{
// 		ws: ws,
// 	}
// }

// func (ws *WebSocket) WsHandler(w http.ResponseWriter, r *http.Request) {

// }

type Hub struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}
}
