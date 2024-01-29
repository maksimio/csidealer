package websocket

import (
	"csidealer/internal/entity"
	"csidealer/internal/usecase"
	"fmt"
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
)

type WebsocketServer struct {
	csiUc       usecase.CsiUC
	port        string
	upgrader    ws.Upgrader
	connections []Connection
}

func (s *WebsocketServer) send(pack entity.ApiPackageAbsPhase) {
	for _, c := range s.connections {
		c.Write(pack)
	}
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
	s.csiUc.OnPushPacket(s.send)

	http.HandleFunc("/", s.startConn)
	log.Print("WebSocket-сервер ожидает подключение на", s.port, "порту")
	http.ListenAndServe(s.port, nil)

}

func (s *WebsocketServer) startConn(w http.ResponseWriter, r *http.Request) {
	conn, _ := s.upgrader.Upgrade(w, r, nil)
	s.connections = append(s.connections, *NewConnection(conn))
}
