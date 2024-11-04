package adapter

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"voo.su/pkg/socket/adapter/encoding"
)

type TcpAdapter struct {
	conn      net.Conn
	reader    *bufio.Reader
	hookClose func(code int, text string) error
}

func NewTcpAdapter(conn net.Conn) (*TcpAdapter, error) {
	return &TcpAdapter{conn: conn, reader: bufio.NewReader(conn)}, nil
}

func (t *TcpAdapter) Network() string {
	return NetworkTcp
}

func (t *TcpAdapter) Read() ([]byte, error) {
	msg, err := encoding.NewDecode(t.reader)
	if err == io.EOF {
		if t.hookClose != nil {
			if err := t.hookClose(1000, "Клиент закрыт"); err != nil {
				return nil, err
			}
		}

		return nil, fmt.Errorf("соединение разорвано")
	}
	if err != nil {
		return nil, fmt.Errorf("ошибка декодирования сообщения: %s", err.Error())
	}

	return msg, nil
}

func (t *TcpAdapter) Write(bytes []byte) error {
	binaryData, err := encoding.NewEncode(bytes)
	if err != nil {
		return err
	}

	_, err = t.conn.Write(binaryData)
	return err
}

func (t *TcpAdapter) Close() error {
	return t.conn.Close()
}

func (t *TcpAdapter) SetCloseHandler(fn func(code int, text string) error) {
	t.hookClose = fn
}
