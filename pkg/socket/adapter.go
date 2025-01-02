// Copyright (c) 2025 Magomedcoder <info@magomedcoder.ru>
// Distributed under the GPL v3 License, see https://github.com/voo-su/server/blob/main/LICENSE

package socket

type IConn interface {
	Read() ([]byte, error)

	Write([]byte) error

	Close() error

	SetCloseHandler(fn func(code int, text string) error)

	Network() string
}
