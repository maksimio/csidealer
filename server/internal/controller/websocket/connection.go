package websocket

import (
	"github.com/google/uuid"
	ws "github.com/gorilla/websocket"
	"time"
)

type Connection struct {
	conn      *ws.Conn
	Uuid      string
	Timestamp int64
}

func NewConnection(conn *ws.Conn) *Connection {
	return &Connection{
		conn:      conn,
		Uuid:      uuid.New().String(),
		Timestamp: time.Now().UnixMilli(),
	}
}

func (c *Connection) Write(data interface{}) {
	c.conn.WriteJSON(data)
}
