package websocket

import (
	"csidealer/internal/usecase"
	"fmt"
	ws "github.com/gorilla/websocket"
	"net/http"
)

type WebsocketServer struct {
	csiUc       usecase.CsiUC
	port        string
	upgrader    ws.Upgrader
	connections []Connection
}

func NewWebsocketServer(uc usecase.CsiUC, port int) *WebsocketServer {
	return &WebsocketServer{
		csiUc: uc,
		port:  "localhost:" + fmt.Sprint(port),
		upgrader: ws.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Пропускаем любой запрос
			},
		},
	}
}

func (s *WebsocketServer) Run() {
	http.HandleFunc("/", s.startConn)
	fmt.Println("WebSocket-сервер ожидает подключение на", s.port, "порту")
	http.ListenAndServe(s.port, nil)
}

func (s *WebsocketServer) startConn(w http.ResponseWriter, r *http.Request) {
	conn, _ := s.upgrader.Upgrade(w, r, nil)
	s.connections = append(s.connections, *NewConnection(conn))
}
