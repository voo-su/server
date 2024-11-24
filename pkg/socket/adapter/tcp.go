package adapter

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"voo.su/pkg/socket/adapter/encoding"
)

type TcpAdapter struct {
	Conn      net.Conn
	Reader    *bufio.Reader
	HookClose func(code int, text string) error
}

func NewTcpAdapter(conn net.Conn) (*TcpAdapter, error) {
	return &TcpAdapter{
		Conn:   conn,
		Reader: bufio.NewReader(conn),
	}, nil
}

func (t *TcpAdapter) Network() string {
	return NetworkTcp
}

func (t *TcpAdapter) Read() ([]byte, error) {
	msg, err := encoding.NewDecode(t.Reader)
	if err == io.EOF {
		if t.HookClose != nil {
			if err := t.HookClose(1000, "Клиент закрыт"); err != nil {
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

	_, err = t.Conn.Write(binaryData)
	return err
}

func (t *TcpAdapter) Close() error {
	return t.Conn.Close()
}

func (t *TcpAdapter) SetCloseHandler(fn func(code int, text string) error) {
	t.HookClose = fn
}
