package socket

type IConn interface {
	Read() ([]byte, error)
	Write([]byte) error
	Close() error
	SetCloseHandler(fn func(code int, text string) error)
	Network() string
}
