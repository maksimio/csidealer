package websocket

import (
	"csidealer/internal/models"
	"csidealer/internal/services/processor"
	"fmt"
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
)

type WebsocketController struct {
	port             string
	upgrader         ws.Upgrader
	connections      []Connection
	in               <-chan models.Package
	processorService *processor.ProcessorService
}

func (s *WebsocketController) send() {
	for {
		pack := <-s.in

		apiPack := models.ApiPackageAbsPhase{
			Timestamp: pack.Timestamp,
			Id:        pack.Uuid,
			Info:      pack.Info,
			Number:    pack.Number,
			Abs:       s.processorService.CsiMap(pack.Data, processor.AbsHandler),
			Phase:     s.processorService.CsiMap(pack.Data, processor.PhaseHandler),
		}

		for _, c := range s.connections {
			c.Write(apiPack)
		}
	}
}

func NewWebsocketController(in <-chan models.Package, processorService *processor.ProcessorService, port int) *WebsocketController {
	return &WebsocketController{
		in:               in,
		processorService: processorService,
		port:             "localhost:" + fmt.Sprint(port),
		upgrader: ws.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Пропускаем любой запрос
			},
		},
	}
}

func (s *WebsocketController) Run() {
	go s.send()

	http.HandleFunc("/", s.startConn)
	log.Print("WebSocket-сервер ожидает подключение на", s.port, "порту")
	http.ListenAndServe(s.port, nil)
}

func (s *WebsocketController) startConn(w http.ResponseWriter, r *http.Request) {
	conn, _ := s.upgrader.Upgrade(w, r, nil)
	s.connections = append(s.connections, *NewConnection(conn))
}
