package adapter

import (
	"github.com/gorilla/websocket"
	"net/http"
)

type WsAdapter struct {
	Conn *websocket.Conn
}

var defaultUpGrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func NewWsAdapter(w http.ResponseWriter, r *http.Request) (*WsAdapter, error) {
	conn, err := defaultUpGrader.Upgrade(w, r, w.Header())
	if err != nil {
		return nil, err
	}

	return &WsAdapter{Conn: conn}, nil
}

func (w *WsAdapter) Network() string {
	return NetworkWss
}

func (w *WsAdapter) Read() ([]byte, error) {
	_, content, err := w.Conn.ReadMessage()
	return content, err
}

func (w *WsAdapter) Write(bytes []byte) error {
	return w.Conn.WriteMessage(websocket.TextMessage, bytes)
}

func (w *WsAdapter) Close() error {
	return w.Conn.Close()
}

func (w *WsAdapter) SetCloseHandler(fn func(code int, text string) error) {
	w.Conn.SetCloseHandler(fn)
}
